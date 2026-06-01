import dotenv from "dotenv";
dotenv.config();

export const config = {
  port: process.env.PORT || 8085,
  databaseUrl: process.env.DATABASE_URL,
  jwtSecret: process.env.JWT_SECRET || "autocare_secret_key",

  razorpayKeyId: process.env.RAZORPAY_KEY_ID || "",
  razorpayKeySecret: process.env.RAZORPAY_KEY_SECRET || "",

  stripeSecretKey: process.env.STRIPE_SECRET_KEY || "",
  stripeWebhookSecret: process.env.STRIPE_WEBHOOK_SECRET || "",

  phonepeMerchantId: process.env.PHONEPE_MERCHANT_ID || "",
  phonepeSaltKey: process.env.PHONEPE_SALT_KEY || "",
  phonepeSaltIndex: process.env.PHONEPE_SALT_INDEX || "1",

  gstPercent: Number(process.env.GST_PERCENT || 18)
};