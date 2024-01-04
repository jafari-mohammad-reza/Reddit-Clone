CREATE TYPE "comminutyTypeEnum" AS ENUM (
  'Private',
  'Public',
  'Restricted'
);

CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "username" varchar,
  "email" varchar,
  "posts" posts,
  "subreddit" subreddit,
  "createdSubreddits" subreddit,
  "joinedSubreddits" subreddit,
  "created_at" timestamp
);

CREATE TABLE "comminuty" (
  "id" integer,
  "topic" varchar,
  "comminutyType" comminutyTypeEnum,
  "description" varchar,
  "name" varchar,
  "karmaPoint" integer,
  "subreddit_id" integer,
  "subreddit" subreddit,
  "post" posts
);

CREATE TABLE "subreddit" (
  "id" integer PRIMARY KEY,
  "moderators" users,
  "name" varchar,
  "joined_users" users,
  "created_at" timestamp,
  "posts" posts,
  "comminuty" comminuty
);

CREATE TABLE "posts" (
  "id" integer PRIMARY KEY,
  "title" varchar,
  "body" text,
  "user_id" integer,
  "status" varchar,
  "created_at" timestamp,
  "subreddit" subreddit,
  "subreddit_id" integer,
  "comminuty_id" integer
);

CREATE TABLE "likes" (
  "id" integer PRIMARY KEY,
  "post_id" integer,
  "post" posts,
  "user_id" integer,
  "user" users
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "posts" ("user_id");

CREATE INDEX ON "likes" ("post_id");

CREATE INDEX ON "likes" ("user_id");

CREATE INDEX ON "likes" ("user_id", "post_id");

COMMENT ON COLUMN "posts"."body" IS 'Content of the post';

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "subreddit" ADD FOREIGN KEY ("moderators") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("subreddit_id") REFERENCES "subreddit" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("id") REFERENCES "likes" ("user_id");

ALTER TABLE "users" ADD FOREIGN KEY ("id") REFERENCES "likes" ("post_id");

CREATE TABLE "subreddit_users" (
  "subreddit_id" integer,
  "users_id" integer,
  PRIMARY KEY ("subreddit_id", "users_id")
);

ALTER TABLE "subreddit_users" ADD FOREIGN KEY ("subreddit_id") REFERENCES "subreddit" ("id");

ALTER TABLE "subreddit_users" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("id");


ALTER TABLE "comminuty" ADD FOREIGN KEY ("subreddit_id") REFERENCES "subreddit" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("comminuty_id") REFERENCES "comminuty" ("id");

ALTER TABLE "subreddit" ADD FOREIGN KEY ("joined_users") REFERENCES "users" ("subreddit");

ALTER TABLE "users" ADD FOREIGN KEY ("createdSubreddits") REFERENCES "subreddit" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("joinedSubreddits") REFERENCES "subreddit" ("id");

ALTER TABLE "comminuty" ADD FOREIGN KEY ("post") REFERENCES "posts" ("id");
