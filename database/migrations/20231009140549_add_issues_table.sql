CREATE TABLE "issues" (
    "id" serial PRIMARY KEY,
    "account_id" bigint NOT NULL,
    "created_by" bigint NOT NULL,
    "project_id" bigint NOT NULL,
    "title" character varying NOT NULL,
    "description" character varying,
    "updated_at" timestamp with time zone DEFAULT current_timestamp,
    "created_at" timestamp with time zone DEFAULT current_timestamp,
    "deleted" boolean DEFAULT false,
    "uuid" uuid NOT NULL,
    CONSTRAINT "issues_account_id_fkey" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id"),
    CONSTRAINT "issues_project_id_fkey" FOREIGN KEY ("project_id") REFERENCES "projects" ("id"),
    CONSTRAINT "issues_created_by_fkey" FOREIGN KEY ("created_by") REFERENCES "users" ("id")
);
