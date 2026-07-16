CREATE TABLE "users" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "pin" VARCHAR(100) NOT NULL,
    "email" VARCHAR(100) UNIQUE NOT NULL,
    "hp_number" VARCHAR(20) UNIQUE NOT NULL,
    "status_account" VARCHAR(20) DEFAULT 'active' NOT NULL
        CHECK ("status_account" IN ('active','suspended','blocked')),
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW()
);