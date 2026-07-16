CREATE TABLE "wallet_logs" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "wallet_id" BIGINT NOT NULL
        REFERENCES wallets(id),
    "action" VARCHAR(50) NOT NULL,
    "description" TEXT,
    "created_at" TIMESTAMP DEFAULT NOW()
);