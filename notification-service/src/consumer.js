import Redis from "ioredis";
import { config } from "./config.js";
import { sendNotificationInternal } from "./service.js";

export function startConsumers() {
  const redis = new Redis({
    host: config.redisHost,
    port: config.redisPort
  });

  redis.subscribe(config.bookingEventQueue, err => {
    if (err) {
      console.error("Redis subscribe failed", err);
      return;
    }

    console.log("Subscribed to", config.bookingEventQueue);
  });

  redis.on("message", async (channel, message) => {
    try {
      const event = JSON.parse(message);

      if (event.event_type === "BOOKING_CREATED") {
        await sendNotificationInternal({
          channel: "EMAIL",
          recipient: "customer@test.com",
          subject: "Booking Created",
          message: `Your booking ${event.booking?.id || ""} has been created.`,
          template_code: "BOOKING_CREATED",
          metadata: event
        });
      }

      if (event.event_type === "BOOKING_CANCELLED") {
        await sendNotificationInternal({
          channel: "SMS",
          recipient: "+919999999999",
          message: `Booking ${event.booking?.id || ""} has been cancelled.`,
          template_code: "BOOKING_CANCELLED",
          metadata: event
        });
      }
    } catch (err) {
      console.error("Notification consumer failed", err);
    }
  });
}