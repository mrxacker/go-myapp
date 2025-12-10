CREATE TABLE users (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   username TEXT NOT NULL,
   email TEXT NOT NULL,
   created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
   updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX ON users (email);