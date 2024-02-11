CREATE TABLE "songs" (
  "song_id" bigserial PRIMARY KEY,
  "song_name" varchar NOT NULL,
  "artist_id" bigint NOT NULL,
  "artist_name" varchar NOT NULL, 
  "thumbnail_s3_path" varchar NOT NULL,
  "mp3_s3_path" varchar NOT NULL
);

CREATE TABLE "artists" (
  "artist_id" bigserial PRIMARY KEY,
  "artist_name" varchar NOT NULL
);

CREATE TABLE "users" (
  "user_id" bigserial PRIMARY KEY,
  "user_name" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "data_of_birth" timestamptz,
  "last_login" timestamptz,
  "created_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "songs" ("song_name");

CREATE INDEX ON "artists" ("artist_name");

ALTER TABLE "songs" ADD FOREIGN KEY ("artist_id") REFERENCES "artists" ("artist_id");

