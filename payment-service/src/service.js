import crypto from "crypto";
import QRCode from "qrcode";
import PDFDocument from "pdfkit";
import Stripe from "stripe";
import Razorpay from "razorpay";
import pg from "pg";
import { config } from "./config.js";

const { Pool } = pg;

const pool = new Pool({
  connectionString: config.databaseUrl,
  connectionTimeoutMillis: 5000,
  idleTimeoutMillis: 10000,
  max: 5
});

const stripe = config.stripeSecretKey
  ? new Stripe(config.stripeSecretKey)
  : null;

const razorpay = config.razorpayKeyId
  ? new Razorpay({
      key_id: config.razorpayKeyId,
      key_secret: config.razorpayKeySecret
    })
  : null;

export async function createPayment(request, reply) {
  const {
    booking_id,
    customer_id,
    amount,
    provider,
    currency = "INR"
  } = request.body;

  const gstAmount = Number(((amount * config.gstPercent) / 100).toFixed(2));
  const totalAmount = Number((amount + gstAmount).toFixed(2));

  let providerOrderId = null;

  if (provider === "RAZORPAY" && razorpay) {
    const order = await razorpay.orders.create({
      amount: Math.round(totalAmount * 100),
      currency,
      receipt: booking_id
    });

    providerOrderId = order.id;
  }

  if (provider === "STRIPE" && stripe) {
    const intent = await stripe.paymentIntents.create({
      amount: Math.round(totalAmount * 100),
      currency: currency.toLowerCase(),
      metadata: {
        booking_id,
        customer_id
      }
    });

    providerOrderId = intent.id;
  }

  if (provider === "PHONEPE") {
    providerOrderId = `phonepe_${Date.now()}`;
  }

  const result = await pool.query(
    `INSERT INTO payments
     (booking_id, customer_id, provider, provider_order_id, amount, gst_amount, total_amount, currency, status, created_at, updated_at)
     VALUES ($1,$2,$3,$4,$5,$6,$7,$8,'CREATED',NOW(),NOW())
     RETURNING *`,
    [
      booking_id,
      customer_id,
      provider,
      providerOrderId,
      amount,
      gstAmount,
      totalAmount,
      currency
    ]
  );

  return reply.code(201).send(result.rows[0]);
}

export async function getPayment(request, reply) {
  const { id } = request.params;

  const result = await pool.query(
    `SELECT * FROM payments WHERE id=$1`,
    [id]
  );

  if (result.rowCount === 0) {
    return reply.code(404).send({ error: "payment not found" });
  }

  return result.rows[0];
}

export async function generateUpiQr(request, reply) {
  const {
    payee_vpa,
    payee_name,
    amount,
    transaction_note = "AutoCare Payment"
  } = request.body;

  const upiUrl =
    `upi://pay?pa=${encodeURIComponent(payee_vpa)}` +
    `&pn=${encodeURIComponent(payee_name)}` +
    `&am=${encodeURIComponent(amount)}` +
    `&cu=INR` +
    `&tn=${encodeURIComponent(transaction_note)}`;

  const qrBase64 = await QRCode.toDataURL(upiUrl);

  return {
    upi_url: upiUrl,
    qr_base64: qrBase64
  };
}

export async function generateInvoice(request, reply) {
  const { id } = request.params;

  const result = await pool.query(
    `SELECT * FROM payments WHERE id=$1`,
    [id]
  );

  if (result.rowCount === 0) {
    return reply.code(404).send({ error: "payment not found" });
  }

  const payment = result.rows[0];

  const doc = new PDFDocument();
  const chunks = [];

  doc.on("data", chunk => chunks.push(chunk));

  const pdfPromise = new Promise(resolve => {
    doc.on("end", () => resolve(Buffer.concat(chunks)));
  });

  doc.fontSize(20).text("AutoCare Invoice", { align: "center" });
  doc.moveDown();
  doc.fontSize(12).text(`Invoice ID: ${payment.id}`);
  doc.text(`Booking ID: ${payment.booking_id}`);
  doc.text(`Customer ID: ${payment.customer_id}`);
  doc.text(`Provider: ${payment.provider}`);
  doc.text(`Amount: ${payment.amount}`);
  doc.text(`GST: ${payment.gst_amount}`);
  doc.text(`Total: ${payment.total_amount}`);
  doc.text(`Status: ${payment.status}`);

  doc.end();

  const pdf = await pdfPromise;

  reply
    .header("Content-Type", "application/pdf")
    .header("Content-Disposition", `inline; filename=invoice-${payment.id}.pdf`)
    .send(pdf);
}

export async function calculateGst(request, reply) {
  const { amount } = request.body;

  const gstAmount = Number(((amount * config.gstPercent) / 100).toFixed(2));
  const totalAmount = Number((amount + gstAmount).toFixed(2));

  return {
    amount,
    gst_percent: config.gstPercent,
    gst_amount: gstAmount,
    total_amount: totalAmount
  };
}

export async function handleWebhook(request, reply) {
  const { provider } = request.params;

  if (provider === "razorpay") {
    return handleRazorpayWebhook(request, reply);
  }

  if (provider === "stripe") {
    return handleStripeWebhook(request, reply);
  }

  if (provider === "phonepe") {
    return handlePhonePeWebhook(request, reply);
  }

  return reply.code(400).send({ error: "unsupported provider" });
}

async function handleRazorpayWebhook(request, reply) {
  const signature = request.headers["x-razorpay-signature"];
  const body = JSON.stringify(request.body);

  const expected = crypto
    .createHmac("sha256", config.razorpayKeySecret)
    .update(body)
    .digest("hex");

  if (signature !== expected) {
    return reply.code(401).send({ error: "invalid razorpay signature" });
  }

  const event = request.body;

  await pool.query(
    `INSERT INTO payment_webhooks (provider, event_type, payload, created_at)
     VALUES ('RAZORPAY', $1, $2, NOW())`,
    [event.event || "unknown", event]
  );

  return { status: "ok" };
}

async function handleStripeWebhook(request, reply) {
  if (!stripe || !config.stripeWebhookSecret) {
    return reply.code(400).send({ error: "stripe not configured" });
  }

  await pool.query(
    `INSERT INTO payment_webhooks (provider, event_type, payload, created_at)
     VALUES ('STRIPE', $1, $2, NOW())`,
    [request.body.type || "unknown", request.body]
  );

  return { status: "ok" };
}

async function handlePhonePeWebhook(request, reply) {
  await pool.query(
    `INSERT INTO payment_webhooks (provider, event_type, payload, created_at)
     VALUES ('PHONEPE', $1, $2, NOW())`,
    ["payment_update", request.body]
  );

  return { status: "ok" };
}