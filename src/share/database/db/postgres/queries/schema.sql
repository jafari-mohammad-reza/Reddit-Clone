
DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'reaction_type') THEN
            CREATE TYPE reaction_type AS ENUM ('like', 'dislike');
        END IF;
    END
$$;


CREATE TABLE  IF NOT EXISTS  "user" (
                                        "id" SERIAL PRIMARY KEY,
                                        "username" VARCHAR UNIQUE,
                                        "email" VARCHAR UNIQUE,
                                        "password" VARCHAR,
                                        "bio" TEXT,
                                        "profile_picture_url" VARCHAR DEFAULT 'https://i.redd.it/snoovatar/avatars/cba29d3e-fcc2-4f64-b9ca-6d89beac6557.png',
                                        "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_user_username ON "user" ("username");
CREATE INDEX idx_user_email ON "user" ("email");
CREATE TABLE IF NOT EXISTS category (
                                        id SERIAL PRIMARY KEY,
                                        name VARCHAR(255) NOT NULL UNIQUE ,
                                        parent_category_id INTEGER NULL,
                                        FOREIGN KEY (parent_category_id) REFERENCES category(id)
);

CREATE INDEX idx_category_name ON category(name);


CREATE TABLE  IF NOT EXISTS  "subreddit" (
                                             "id" SERIAL PRIMARY KEY,
                                             "name" VARCHAR UNIQUE,
                                             "description" TEXT,
                                             "owner_id" INTEGER,
                                             "category" VARCHAR,
                                             "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                             FOREIGN KEY ("owner_id") REFERENCES "user" ("id")
);
CREATE INDEX idx_subreddit_owner ON "subreddit" ("owner_id");
CREATE INDEX idx_subreddit_name ON "subreddit" ("name");
CREATE INDEX idx_subreddit_category ON "subreddit" ("category");


CREATE TABLE  IF NOT EXISTS  "post" (
                                        "id" SERIAL PRIMARY KEY,
                                        "title" VARCHAR,
                                        "body" TEXT,
                                        "user_id" INTEGER,
                                        "subreddit_id" INTEGER,
                                        "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        "view_count" INTEGER DEFAULT 0,
                                        FOREIGN KEY ("user_id") REFERENCES "user" ("id"),
                                        FOREIGN KEY ("subreddit_id") REFERENCES "subreddit" ("id")
);
CREATE INDEX idx_post_user ON "post" ("user_id");
CREATE INDEX idx_post_subreddit ON "post" ("subreddit_id");
CREATE INDEX idx_post_view_count ON "post" ("view_count");


CREATE TABLE  IF NOT EXISTS  "post_reaction" (
                                                 "id" SERIAL PRIMARY KEY,
                                                 "post_id" INTEGER,
                                                 "user_id" INTEGER,
                                                 "reaction" reaction_type,
                                                 "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                 FOREIGN KEY ("post_id") REFERENCES "post" ("id"),
                                                 FOREIGN KEY ("user_id") REFERENCES "user" ("id")
);
CREATE INDEX idx_post_reaction_post ON "post_reaction" ("post_id");
CREATE INDEX idx_post_reaction_user ON "post_reaction" ("user_id");
CREATE INDEX idx_post_reaction_type ON "post_reaction" ("reaction");


CREATE TABLE  IF NOT EXISTS  "comment" (
                                           "id" SERIAL PRIMARY KEY,
                                           "post_id" INTEGER,
                                           "parent_comment_id" INTEGER,
                                           "user_id" INTEGER,
                                           "content" TEXT,
                                           "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                           "view_count" INTEGER DEFAULT 0,
                                           FOREIGN KEY ("post_id") REFERENCES "post" ("id"),
                                           FOREIGN KEY ("parent_comment_id") REFERENCES "comment" ("id"),
                                           FOREIGN KEY ("user_id") REFERENCES "user" ("id")
);
CREATE INDEX idx_comment_post ON "comment" ("post_id");
CREATE INDEX idx_comment_parent ON "comment" ("parent_comment_id");
CREATE INDEX idx_comment_user ON "comment" ("user_id");
CREATE INDEX idx_comment_view_count ON "comment" ("view_count");


CREATE TABLE  IF NOT EXISTS  "comment_reaction" (
                                                    "id" SERIAL PRIMARY KEY,
                                                    "comment_id" INTEGER,
                                                    "user_id" INTEGER,
                                                    "reaction" reaction_type,
                                                    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                    FOREIGN KEY ("comment_id") REFERENCES "comment" ("id"),
                                                    FOREIGN KEY ("user_id") REFERENCES "user" ("id")
);
CREATE INDEX idx_comment_reaction_comment ON "comment_reaction" ("comment_id");
CREATE INDEX idx_comment_reaction_user ON "comment_reaction" ("user_id");
CREATE INDEX idx_comment_reaction_type ON "comment_reaction" ("reaction");


CREATE TABLE  IF NOT EXISTS  "user_follower" (
                                                 "id" SERIAL PRIMARY KEY,
                                                 "user_id" INTEGER,
                                                 "follower_id" INTEGER,
                                                 "followed_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                 FOREIGN KEY ("user_id") REFERENCES "user" ("id"),
                                                 FOREIGN KEY ("follower_id") REFERENCES "user" ("id")
);
CREATE INDEX idx_user_follower_user ON "user_follower" ("user_id");
CREATE INDEX idx_user_follower_follower ON "user_follower" ("follower_id");
