CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "username" varchar,
  "role" varchar,
  "created_at" timestamps
);

CREATE TABLE "boards" (
  "id" integer PRIMARY KEY,
  "title" varchar,
  "user_id" integer,
  "status" varchar,
  "created_at" timestamp
);

CREATE TABLE "columns" (
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

CREATE TABLE "subtasks" (
  "id" integer PRIMARY KEY,
  "task_id" integer,
  "created_at" timestamp
);

ALTER TABLE "boards" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "columns" ADD FOREIGN KEY ("board_id") REFERENCES "boards" ("id");

ALTER TABLE "task" ADD FOREIGN KEY ("column_id") REFERENCES "columns" ("id");

ALTER TABLE "subtasks" ADD FOREIGN KEY ("task_id") REFERENCES "task" ("id");
