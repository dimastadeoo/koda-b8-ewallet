CREATE TABLE "transactions" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "id_wallet" BIGINT REFERENCES "wallets"("id") NOT NULL,
    "transaction_type" VARCHAR(30) NOT NULL,
    -- TOPUP
    -- TRANSFER
    -- PAYMENT
    -- WITHDRAW
    "amount" NUMERIC(20,2) NOT NULL,
    "status" VARCHAR(20) DEFAULT 'PENDING',
    -- PENDING
    -- SUCCESS
    -- FAILED
    -- CANCELLED
    "description" TEXT,
    "created_at" TIMESTAMP DEFAULT NOW(),
    "completed_at" TIMESTAMP
);