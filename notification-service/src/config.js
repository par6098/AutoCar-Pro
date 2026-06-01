import dotenv from "dotenv";
dotenv.config();

export const config = {
  port: Number(process.env.PORT || 8086),
  databaseUrl: process.env.DATABASE_URL,

  jwtSecret: process.env.JWT_SECRET || "autocare_secret_key",

  redisHost: process.env.REDIS_ADDR || "localhost",
  redisPort: Number(process.env.REDIS_PORT || 6379),
  bookingEventQueue: process.env.BOOKING_EVENT_QUEUE || "booking-events",
  notificationEventQueue: process.env.NOTIFICATION_EVENT_QUEUE || "notification-events",

  twilioSid: process.env.TWILIO_ACCOUNT_SID || "",
  twilioToken: process.env.TWILIO_AUTH_TOKEN || "",
  twilioFrom: process.env.TWILIO_FROM_NUMBER || "",

  msg91AuthKey: process.env.MSG91_AUTH_KEY || "",
  msg91SenderId: process.env.MSG91_SENDER_ID || "",

  resendApiKey: process.env.RESEND_API_KEY || "",
  emailFrom: process.env.EMAIL_FROM || "no-reply@autocare.local",

  fcmServiceAccountJson: process.env.FCM_SERVICE_ACCOUNT_JSON || ""
};