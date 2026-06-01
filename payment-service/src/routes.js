import { authenticate } from "./plugins/auth.js";
import {
  createPayment,
  generateUpiQr,
  generateInvoice,
  handleWebhook,
  getPayment,
  calculateGst
} from "./service.js";

export async function routes(app) {
  app.post("/payments/create", { preHandler: authenticate }, createPayment);
  app.get("/payments/:id", { preHandler: authenticate }, getPayment);
  app.post("/payments/upi/qr", { preHandler: authenticate }, generateUpiQr);
  app.get("/payments/:id/invoice", { preHandler: authenticate }, generateInvoice);
  app.post("/payments/gst/calculate", { preHandler: authenticate }, calculateGst);

  // Webhooks should not use JWT; they use provider signature verification.
  app.post("/payments/webhook/:provider", handleWebhook);
}