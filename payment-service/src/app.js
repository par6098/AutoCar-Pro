import Fastify from "fastify";
import { routes } from "./routes.js";

export function buildApp() {
  const app = Fastify({
    logger: true
  });

  app.register(routes);

  return app;
}

if (process.env.LOCAL_RUN === "true") {
  const app = buildApp();

  app.listen({
    port: Number(process.env.PORT || 8085),
    host: "0.0.0.0"
  });
}