import awsLambdaFastify from "@fastify/aws-lambda";
import { buildApp } from "./app.js";

const app = buildApp();

const proxy = awsLambdaFastify(app);

export const handler = async (event, context) => {
  context.callbackWaitsForEmptyEventLoop = false;

  console.log("PaymentFunction invoked:", {
    path: event.path,
    method: event.httpMethod
  });

  return proxy(event, context);
};