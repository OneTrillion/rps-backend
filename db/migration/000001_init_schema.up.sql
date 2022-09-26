CREATE TABLE "player" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL DEFAULT(''),
  "score" int NOT NULL DEFAULT 0,
  "health" int NOT NULL DEFAULT 100,
  "ult_meter" int NOT NULL DEFAULT 0
);

CREATE TABLE "opponent" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "health" int NOT NULL DEFAULT 100
);

CREATE TABLE "scoreboard" (
  "id" bigserial PRIMARY KEY,
  "player_id" bigint NOT NULL,
  "username" varchar NOT NULL,
  "score" int NOT NULL
);

CREATE TABLE "choice" (
  "id" bigserial PRIMARY KEY,
  "player_id" bigint NOT NULL,
  "rps_choice" int NOT NULL
);

ALTER TABLE "scoreboard" ADD FOREIGN KEY ("player_id") REFERENCES "player" ("id");
ALTER TABLE "choice" ADD FOREIGN KEY ("player_id") REFERENCES "player" ("id");
