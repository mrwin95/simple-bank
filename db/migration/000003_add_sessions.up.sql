CREATE TABLE
    "sessions" (
        "id" uuid PRIMARY KEY,
        "username" varchar not null,
        "refresh_token" varchar NOT NULL,
        "user_agent" varchar NOT NULL,
        "client_ip" varchar NOT NULL,
        "is_blocked" boolean NOT NULL default false,
        "expires_at" timestamp NOT NULL,
        "created_at" timestamp NOT NULL DEFAULT (now ())
    );

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");