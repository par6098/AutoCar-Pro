import jwt from "jsonwebtoken";
import { config } from "../config.js";

export async function authenticate(request, reply) {
  const authHeader = request.headers.authorization;

  if (!authHeader || !authHeader.startsWith("Bearer ")) {
    return reply.code(401).send({
      error: "missing or invalid authorization header"
    });
  }

  const token = authHeader.replace("Bearer ", "");

  try {
    const decoded = jwt.verify(token, config.jwtSecret);
    request.user = decoded;
  } catch (err) {
    return reply.code(401).send({
      error: "invalid token"
    });
  }
}