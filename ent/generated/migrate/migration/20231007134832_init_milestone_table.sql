CREATE TABLE "milestones" (
    "id" serial PRIMARY KEY,
    "account_id" bigint NOT NULL,
    "created_by" bigint NOT NULL,
    "project_id" bigint NOT NULL,
    "title" character varying NOT NULL,
    "description" character varying,
    "end_date" timestamp with time zone,
    "updated_at" timestamp with time zone DEFAULT current_timestamp,
    "created_at" timestamp with time zone DEFAULT current_timestamp,
    "deleted" boolean DEFAULT false,
    "uuid" uuid NOT NULL,
    CONSTRAINT "tasks_account_id_fkey" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id"),
    CONSTRAINT "tasks_project_id_fkey" FOREIGN KEY ("project_id") REFERENCES "projects" ("id"),
    CONSTRAINT "tasks_created_by_fkey" FOREIGN KEY ("created_by") REFERENCES "users" ("id")
);
