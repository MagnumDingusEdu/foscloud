CREATE TABLE "nodes" (
  "id" bigserial PRIMARY KEY,
  "parent_id" bigint,
  "name" varchar NOT NULL,
  "is_dir" boolean NOT NULL DEFAULT false,
  "filesize" bigint,
  "depth" int,
  "lineage" varchar,
  "owner" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "last_login" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "links" (
  "id" bigserial PRIMARY KEY,
  "node" bigint NOT NULL,
  "link" varchar NOT NULL,
  "clicks" int DEFAULT 0 NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "nodes" ADD FOREIGN KEY ("parent_id") REFERENCES "nodes" ("id") ON DELETE CASCADE;

ALTER TABLE "nodes" ADD FOREIGN KEY ("owner") REFERENCES "accounts" ("id") ON DELETE CASCADE;

ALTER TABLE "links" ADD FOREIGN KEY ("node") REFERENCES "nodes" ("id") ON DELETE CASCADE;

CREATE INDEX ON "nodes" ("name");

CREATE INDEX ON "accounts" ("username");

CREATE INDEX ON "accounts" ("name");

CREATE INDEX ON "accounts" ("email");

COMMENT ON COLUMN "nodes"."parent_id" IS 'parent_id is null for root node';

COMMENT ON COLUMN "nodes"."depth" IS 'depth starting from parent node (0)';

COMMENT ON COLUMN "nodes"."lineage" IS 'used for breadcrumbs';

COMMENT ON COLUMN "accounts"."password" IS 'hashed password';

