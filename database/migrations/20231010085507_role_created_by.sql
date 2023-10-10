-- Add "updated_at" and "created_at" columns to the ProjectServiceProviders table
ALTER TABLE "roles"
ADD COLUMN  "created_by" bigint NOT NULL;
