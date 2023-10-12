CREATE TABLE "teams" (
    "id" serial PRIMARY KEY,
    "account_id" bigint NOT NULL,
    "created_by" bigint NOT NULL,
    "name" varchar NOT NULL,
    "description" character varying,
    "status" character varying NOT NULL,
    "updated_at" timestamp with time zone DEFAULT current_timestamp,
    "created_at" timestamp with time zone DEFAULT current_timestamp,
    "deleted" boolean DEFAULT false,
    "uuid" uuid NOT NULL,
    CONSTRAINT "teams_account_id_fkey" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id"),
    CONSTRAINT "teams_created_by_fkey" FOREIGN KEY ("created_by") REFERENCES "users" ("id")
);
