BEGIN;
CREATE TABLE IF NOT EXISTS "posts" (
    "pk" serial NOT NULL,
    "id" character varying(21) NOT NULL,
    "author_id" character varying(21) NOT NULL,
    "title" character varying(255) NOT NULL,
    "content" text NOT NULL,
    "created_at" timestamp without time zone NOT NULL,
    "updated_at" timestamp without time zone NOT NULL,
    CONSTRAINT "pk_posts" PRIMARY KEY ("pk"),
    CONSTRAINT "uq_posts_id" UNIQUE ("id")
);
COMMIT;

