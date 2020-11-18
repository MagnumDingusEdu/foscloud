CREATE TABLE "authtoken"
(
    "id"         bigserial PRIMARY KEY,
    "token"      varchar     NOT NULL,
    "account"    bigint      NOT NULL UNIQUE,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "last_used"  timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "authtoken"
    ADD FOREIGN KEY ("account") REFERENCES "accounts" ("id") ON DELETE CASCADE;
