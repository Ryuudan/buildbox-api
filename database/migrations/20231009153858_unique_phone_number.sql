ALTER TABLE "accounts"
ADD CONSTRAINT unique_phone_number UNIQUE (phone_number);

ALTER TABLE "users"
ADD COLUMN "phone_number" varchar UNIQUE;
