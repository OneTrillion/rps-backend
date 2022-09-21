CREATE TABLE "player" (
  "id" bigint PRIMARY KEY,
  "username" varchar NOT NULL,
  "score" int NOT NULL,
  "health" int NOT NULL DEFAULT 100,
  "ult_meter" int NOT NULL DEFAULT 0
);

CREATE TABLE "opponent" (
  "id" bigint PRIMARY KEY,
  "name" varchar NOT NULL,
  "health" int NOT NULL DEFAULT 100
);

CREATE TABLE "scoreboard" (
  "id" bigint PRIMARY KEY,
  "player_id" bigint NOT NULL,
  "username" varchar NOT NULL,
  "score" int NOT NULL
);

ALTER TABLE "scoreboard" ADD FOREIGN KEY ("player_id") REFERENCES "player" ("id");
