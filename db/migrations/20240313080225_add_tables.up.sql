CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "hand" varchar,
  "score" int,
  "room" varchar
);

CREATE TABLE "rooms" (
  "id" bigserial PRIMARY KEY,
  "players" varchar,
  "code" int,
  "deck" varchar,
  "scores" varchar
)