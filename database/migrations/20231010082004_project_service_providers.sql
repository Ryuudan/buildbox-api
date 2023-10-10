-- Create the ProjectServiceProviders (intermediary) table
CREATE TABLE "project_service_providers" (
  "id" serial PRIMARY KEY,
  "created_by" bigint NOT NULL,
  "project_id" bigint NOT NULL,
  "service_provider_id" bigint NOT NULL,
  "uuid" uuid NOT NULL,
  CONSTRAINT "psp_project_id" FOREIGN KEY ("project_id") REFERENCES "projects" ("id"),
  CONSTRAINT "psp_service_provider_id" FOREIGN KEY ("service_provider_id") REFERENCES "service_providers" ("id")
);

-- Create a unique constraint to ensure each service provider is associated with a project only once
-- This prevents duplicates in the many-to-many relationship
CREATE UNIQUE INDEX unique_project_service_providers
ON project_service_providers (project_id, service_provider_id);
