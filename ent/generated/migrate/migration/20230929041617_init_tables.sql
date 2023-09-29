-- Create the "accounts" table
CREATE TABLE "accounts" (
    "id" serial PRIMARY KEY,
    "name" varchar NOT NULL,
    "email" varchar UNIQUE,
    "phone_number" varchar,
    "updated_at" timestamp with time zone,
    "created_at" timestamp with time zone,
    "uuid" uuid NOT NULL
);

-- Create the "plans" table
CREATE TABLE "plans" (
    "id" serial PRIMARY KEY,
    "name" varchar NOT NULL UNIQUE,
    "description" text NOT NULL,
    "price" double precision DEFAULT 0,
    "updated_at" timestamp DEFAULT current_timestamp,
    "created_at" timestamp DEFAULT current_timestamp,
    "uuid" uuid NOT NULL
);

-- Create the "users" table
CREATE TABLE "users" (
    "id" serial PRIMARY KEY,
    "account_id" bigint NOT NULL,
    "first_name" character varying NOT NULL,
    "last_name" character varying NOT NULL,
    "middle_name" character varying,
    "birthday" timestamp with time zone,
    "email" character varying NOT NULL UNIQUE,
    "password" character varying NOT NULL,
    "updated_at" timestamp with time zone,
    "created_at" timestamp with time zone,
    "uuid" uuid NOT NULL,
    CONSTRAINT "users_account_id_fkey" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id")
);

-- Create the "subscriptions" table
CREATE TABLE "subscriptions" (
    "id" serial PRIMARY KEY,
    "account_id" bigint NOT NULL,
    "plan_id" bigint NOT NULL,
    "start_date" timestamp with time zone DEFAULT current_timestamp,
    "end_date" timestamp with time zone,
    "status" character varying DEFAULT 'active' CHECK (status IN ('active', 'canceled', 'expired')),
    "billing_cycle" character varying DEFAULT 'monthly' CHECK (billing_cycle IN ('monthly', 'yearly')),
    "discount" double precision,
    "updated_at" timestamp with time zone,
    "created_at" timestamp with time zone DEFAULT current_timestamp,
    "uuid" uuid NOT NULL,
    CONSTRAINT "subscriptions_account_id_fkey" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id"),
    CONSTRAINT "subscriptions_plan_id_fkey" FOREIGN KEY ("plan_id") REFERENCES "plans" ("id")
);

-- Create the "projects" table
CREATE TABLE "projects" (
    "id" serial PRIMARY KEY,
    "account_id" bigint NOT NULL,
    "created_by" bigint NOT NULL,
    "client_id" bigint,
    "manager_id" bigint,
    "name" character varying(100) NOT NULL,
    "status" character varying,
    "location" character varying,
    "budget" double precision,
    "description" character varying(300),
    "notes" character varying(200),
    "start_date" timestamp with time zone DEFAULT current_timestamp,
    "end_date" timestamp with time zone,
    "uuid" uuid NOT NULL,
    "deleted" boolean DEFAULT false,
    "updated_at" timestamp with time zone DEFAULT current_timestamp,
    "created_at" timestamp with time zone DEFAULT current_timestamp,
    CONSTRAINT "projects_account_id_fkey" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id")
);
