CREATE TABLE "sessions"(
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "token" VARCHAR(100) UNIQUE,
    "status" VARCHAR(30) NOT NULL CHECK ("status" IN ('login', 'logout', 'register','forgot-pin')),
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW()
);