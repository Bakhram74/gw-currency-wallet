
CREATE TABLE "wallet" (
  "user_id" uuid PRIMARY KEY,
  "usd" float DEFAULT 0,
  "rub" float DEFAULT 0,
  "eur" float DEFAULT 0
);

ALTER TABLE "wallet" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;
