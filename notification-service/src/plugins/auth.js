import jwt from "jsonwebtoken";
import { config } from "../config.js";

export async function authenticate(request, reply) {
  const header = request.headers.authorization;

  if (!header || !header.startsWith("Bearer ")) {
    return reply.code(401).send({ error: "missing or invalid authorization header" });
  }

  try {
    request.user = jwt.verify(header.replace("Bearer ", ""), config.jwtSecret);
  } catch {
    return reply.code(401).send({ error: "invalid token" });
  }
}