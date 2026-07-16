CREATE TABLE "users_logs"(
    "id_user" BIGINT REFERENCES "users"("id") NOT NULL,
    "id_sesion" BIGINT REFERENCES "sessions"("id") NOT NULL, 
    "activity_detail" TEXT NOT NULL,
    "ip_address" VARCHAR(80),
    "created_at" TIMESTAMP DEFAULT NOW()
);