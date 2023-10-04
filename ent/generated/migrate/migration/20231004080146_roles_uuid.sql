-- Alter the "roles" table to add a new "uuid" column
ALTER TABLE "roles"
ADD COLUMN "uuid" uuid NOT NULL;
