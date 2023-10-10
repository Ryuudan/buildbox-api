CREATE TABLE "service_providers" (
    "id" serial PRIMARY KEY,
    "account_id" bigint NOT NULL,
    "created_by" bigint NOT NULL,
    "name" varchar NOT NULL,
    "email" varchar  NOT NULL UNIQUE,
    "description" character varying,
    "status" character varying NOT NULL,
    "phone_number" varchar UNIQUE,
    "updated_at" timestamp with time zone DEFAULT current_timestamp,
    "created_at" timestamp with time zone DEFAULT current_timestamp,
    "deleted" boolean DEFAULT false,
    "uuid" uuid NOT NULL,
    CONSTRAINT "service_providers_account_id_fkey" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id"),
    CONSTRAINT "service_providers_created_by_fkey" FOREIGN KEY ("created_by") REFERENCES "users" ("id")
);
