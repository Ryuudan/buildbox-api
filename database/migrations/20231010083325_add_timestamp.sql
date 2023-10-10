-- Add "updated_at" and "created_at" columns to the ProjectServiceProviders table
ALTER TABLE "project_service_providers"
ADD COLUMN "updated_at" timestamp with time zone DEFAULT current_timestamp,
ADD COLUMN "created_at" timestamp with time zone DEFAULT current_timestamp;
