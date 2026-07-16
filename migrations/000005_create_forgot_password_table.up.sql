CREATE TABLE "forgot_pin"(
    "id_sessions" BIGINT REFERENCES "sessions"("id") NOT NULL,
    "id_user" BIGINT REFERENCES "users"("id") NOT NULL, 
    "previos_pin" VARCHAR(100) NOT NULL,
    "update_pin" VARCHAR(100) NOT NULL,
    "ip_address" VARCHAR(80),
    "created_at" TIMESTAMP DEFAULT NOW()
);