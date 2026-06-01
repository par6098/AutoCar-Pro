import { authenticate } from "./plugins/auth.js";
import {
  sendNotification,
  createCampaign,
  sendCampaign,
  getNotificationLogs
} from "./service.js";

export async function routes(app) {
  app.post("/notifications/send", { preHandler: authenticate }, sendNotification);
  app.post("/notifications/campaigns", { preHandler: authenticate }, createCampaign);
  app.post("/notifications/campaigns/:id/send", { preHandler: authenticate }, sendCampaign);
  app.get("/notifications/logs", { preHandler: authenticate }, getNotificationLogs);
}