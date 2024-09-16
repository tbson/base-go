-- Create "variables" table
CREATE TABLE "public"."variables" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "key" text NOT NULL,
  "value" text NOT NULL DEFAULT '',
  "description" text NOT NULL DEFAULT '',
  "data_type" text NOT NULL DEFAULT 'STRING',
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_variables_key" UNIQUE ("key"),
  CONSTRAINT "chk_variables_data_type" CHECK (data_type = ANY (ARRAY['STRING'::text, 'INTEGER'::text, 'FLOAT'::text, 'BOOLEAN'::text, 'DATE'::text, 'DATETIME'::text]))
);
-- Create index "idx_variables_deleted_at" to table: "variables"
CREATE INDEX "idx_variables_deleted_at" ON "public"."variables" ("deleted_at");
