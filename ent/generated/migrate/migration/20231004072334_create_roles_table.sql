-- Create the "roles" table
CREATE TABLE "roles" (
    "id" serial PRIMARY KEY,
    "account_id" integer,
    "name" varchar NOT NULL,
    "description" varchar,
    "updated_at"  timestamp with time zone DEFAULT current_timestamp,
    "created_at"  timestamp with time zone DEFAULT current_timestamp,
    CONSTRAINT "fk_account_id" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id")
);
