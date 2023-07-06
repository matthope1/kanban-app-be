CREATE TABLE "board" (
  "id" integer PRIMARY KEY,
  "title" varchar,
  "user_email" varchar NOT NULL,
  "status" varchar,
  "created_at" timestamp
);

CREATE TABLE "column" (
  "id" integer PRIMARY KEY,
  "title" varchar,
  "board_id" integer,
  "created_at" timestamp
);

CREATE TABLE "task" (
  "id" integer PRIMARY KEY,
  "status" varchar,
  "desc" varchar,
  "column_id" integer,
  "created_at" timestamp
);

CREATE TABLE "subtask" (
  "id" integer PRIMARY KEY,
  "task_id" integer,
  "created_at" timestamp
);



CREATE INDEX ON "board" ("user_email");

CREATE INDEX ON "column" ("board_id");

ALTER TABLE "column" ADD FOREIGN KEY ("board_id") REFERENCES "board" ("id");

ALTER TABLE "task" ADD FOREIGN KEY ("column_id") REFERENCES "column" ("id");

ALTER TABLE "subtask" ADD FOREIGN KEY ("task_id") REFERENCES "task" ("id");
