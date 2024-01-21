CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "username" varchar,
  "email" varchar,
  "posts" posts,
  "subreddit" subreddit,
  "created_at" timestamp
);

CREATE TABLE "subreddit" (
  "id" integer PRIMARY KEY,
  "owner" users,
  "owner_id" integer,
  "created_at" timestamp,
  "posts" posts
);

CREATE TABLE "posts" (
  "id" integer PRIMARY KEY,
  "title" varchar,
  "body" text,
  "user_id" integer,
  "status" varchar,
  "created_at" timestamp,
  "subreddit" subreddit,
  "subreddit_id" integer
);

CREATE TABLE "likes" (
  "id" integer PRIMARY KEY,
  "post_id" integer,
  "post" posts,
  "user_id" integer,
  "user" users
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "subreddit" ("owner_id");

CREATE INDEX ON "posts" ("user_id");

CREATE INDEX ON "likes" ("post_id");

CREATE INDEX ON "likes" ("user_id");

CREATE INDEX ON "likes" ("user_id", "post_id");

COMMENT ON COLUMN "posts"."body" IS 'Content of the post';

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "subreddit" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("subreddit_id") REFERENCES "subreddit" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("id") REFERENCES "likes" ("user_id");

ALTER TABLE "users" ADD FOREIGN KEY ("id") REFERENCES "likes" ("post_id");
