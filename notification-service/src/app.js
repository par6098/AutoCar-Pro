import Fastify from "fastify";
import { routes } from "./routes.js";
import { startConsumers } from "./consumer.js";

export function buildApp() {
  const app = Fastify({ logger: true });
  app.register(routes);
  return app;
}

if (process.env.LOCAL_RUN === "true") {
  const app = buildApp();

  startConsumers();

  app.listen({
    port: Number(process.env.PORT || 8086),
    host: "0.0.0.0"
  });
}