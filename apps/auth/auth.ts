import { betterAuth } from "better-auth";
import { drizzleAdapter } from "better-auth/adapters/drizzle";
import { db } from "./src/db/index";
import { jwt } from "better-auth/plugins";
import * as schema from "./src/db/schema";

export const auth = betterAuth({
  plugins: [jwt()],
  database: drizzleAdapter(db, {
    provider: "sqlite",
    schema: schema,
  }),
  emailAndPassword: {
    enabled: true,
  },
  trustedOrigins: ["http://localhost:5173"],
});
