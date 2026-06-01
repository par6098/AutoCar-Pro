import pg from "pg";
import twilio from "twilio";
import { Resend } from "resend";
import { config } from "./config.js";

const { Pool } = pg;

const pool = new Pool({
  connectionString: config.databaseUrl,
  connectionTimeoutMillis: 5000,
  idleTimeoutMillis: 10000,
  max: 5
});

const twilioClient =
  config.twilioSid && config.twilioToken
    ? twilio(config.twilioSid, config.twilioToken)
    : null;

const resend =
  config.resendApiKey
    ? new Resend(config.resendApiKey)
    : null;

export async function sendNotification(request, reply) {
  const result = await sendNotificationInternal(request.body);
  return reply.code(201).send(result);
}

export async function sendNotificationInternal(payload) {
  const {
    channel,
    recipient,
    subject,
    message,
    template_code = "MANUAL",
    metadata = {}
  } = payload;

  let provider = "MOCK";
  let providerResponse = {};

  if (channel === "SMS") {
    provider = config.twilioSid ? "TWILIO" : "MOCK";

    if (twilioClient) {
      providerResponse = await twilioClient.messages.create({
        from: config.twilioFrom,
        to: recipient,
        body: message
      });
    }
  }

  if (channel === "EMAIL") {
    provider = config.resendApiKey ? "RESEND" : "MOCK";

    if (resend) {
      providerResponse = await resend.emails.send({
        from: config.emailFrom,
        to: recipient,
        subject: subject || "AutoCare Notification",
        html: `<p>${message}</p>`
      });
    }
  }

  if (channel === "PUSH") {
    provider = "FCM_MOCK";
    providerResponse = { status: "push mock sent" };
  }

  const dbResult = await pool.query(
    `INSERT INTO notification_logs
     (channel, recipient, subject, message, template_code, provider, provider_response, status, metadata, created_at)
     VALUES ($1,$2,$3,$4,$5,$6,$7,'SENT',$8,NOW())
     RETURNING *`,
    [
      channel,
      recipient,
      subject || null,
      message,
      template_code,
      provider,
      providerResponse,
      metadata
    ]
  );

  return dbResult.rows[0];
}

export async function createCampaign(request, reply) {
  const { name, channel, audience_query, message, subject } = request.body;

  const result = await pool.query(
    `INSERT INTO notification_campaigns
     (name, channel, audience_query, subject, message, status, created_at)
     VALUES ($1,$2,$3,$4,$5,'DRAFT',NOW())
     RETURNING *`,
    [name, channel, audience_query, subject || null, message]
  );

  return reply.code(201).send(result.rows[0]);
}

export async function sendCampaign(request, reply) {
  const { id } = request.params;

  await pool.query(
    `UPDATE notification_campaigns
     SET status='SENT'
     WHERE id=$1`,
    [id]
  );

  return { message: "campaign marked as sent" };
}

export async function getNotificationLogs(request, reply) {
  const result = await pool.query(
    `SELECT *
     FROM notification_logs
     ORDER BY created_at DESC
     LIMIT 100`
  );

  return result.rows;
}