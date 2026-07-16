CREATE TABLE "wallets" {
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "id_user" BIGINT REFERENCES "users"("id") NOT NULL,
    "balance" NUMERIC(18,2) NOT NULL DEFAULT 0,
    "currency" VARCHAR(3) NOT NULL DEFAULT 'IDR',
    "status" VARCHAR(20) NOT NULL DEFAULT 'active'
        CHECK ("status_account" IN ('active','suspended','blocked')),
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW()
};