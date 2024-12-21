CREATE TABLE IF NOT EXISTS "user" (
  id uuid DEFAULT gen_random_uuid(),
  username VARCHAR NOT NULL UNIQUE, 
  password VARCHAR NOT NULL,
  email VARCHAR NOT NULL UNIQUE CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'), 
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  PRIMARY KEY (id)
);

CREATE INDEX idx_user_name ON "user" ("username");
