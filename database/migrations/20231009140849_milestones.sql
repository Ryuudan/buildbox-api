-- Rename the "tasks_account_id_fkey" constraint to "milestones_account_id_fkey"
ALTER TABLE "milestones"
RENAME CONSTRAINT "tasks_account_id_fkey" TO "milestones_account_id_fkey";

-- Rename the "tasks_project_id_fkey" constraint to "milestones_project_id_fkey"
ALTER TABLE "milestones"
RENAME CONSTRAINT "tasks_project_id_fkey" TO "milestones_project_id_fkey";

-- Rename the "tasks_created_by_fkey" constraint to "milestones_created_by_fkey"
ALTER TABLE "milestones"
RENAME CONSTRAINT "tasks_created_by_fkey" TO "milestones_created_by_fkey";
