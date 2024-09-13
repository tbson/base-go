-- Create "variables" table
CREATE TABLE "public"."variables" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "key" character varying(255) NOT NULL,
  "value" character varying(255) NOT NULL DEFAULT '',
  "description" character varying(255) NOT NULL DEFAULT '',
  "type" smallint NOT NULL DEFAULT 1,
  PRIMARY KEY ("id")
);
-- Create index "idx_variables_deleted_at" to table: "variables"
CREATE INDEX "idx_variables_deleted_at" ON "public"."variables" ("deleted_at");
-- Create index "uni_variables_key" to table: "variables"
CREATE UNIQUE INDEX "uni_variables_key" ON "public"."variables" ("key");
