CREATE TABLE "transfers" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "trx_id" BIGINT REFERENCES "transactions"("id") NOT NULL,
    "sender_wallet_id" BIGINT REFERENCES "wallets"("id") NOT NULL,
    "receiver_wallet_id" BIGINT REFERENCES "wallets"("id") NOT NULL,
    "notes" TEXT,
    "created_at" TIMESTAMP DEFAULT NOW()
);