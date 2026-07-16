CREATE TABLE "profiles" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "id_user" BIGINT UNIQUE REFERENCES "users"("id") NOT NULL,
    "nik" VARCHAR(16) UNIQUE,
    "name" VARCHAR(150) NOT NULL,
    "address" TEXT,
    "gender" VARCHAR(1) CHECK ("gender" IN ('M', 'F')) NOT NULL,
    "place_birth" VARCHAR(100),
    "date_birth" DATE,
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW()
);

CREATE UNIQUE INDEX unique_nik_idx ON "profiles"("nik") WHERE "nik" IS NOT NULL;