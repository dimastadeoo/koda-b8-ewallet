CREATE TABLE "sessions"(
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "id_user" BIGINT REFERENCES "users"("id") NOT NULL, 
    "token" VARCHAR(100) UNIQUE,
    "status" VARCHAR(30) NOT NULL CHECK ("status" IN ('login', 'logout', 'register','forgot-password')),
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW()
);