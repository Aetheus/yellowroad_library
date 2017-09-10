-- +goose Up
-- +goose StatementBegin

CREATE TABLE "Books" (
    id integer NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    "deletedAt" timestamp with time zone,
    "createdAt" timestamp with time zone NOT NULL,
    "updatedAt" timestamp with time zone NOT NULL,
    first_chapter_id integer,
    creator_id integer,
    permissions text
);
CREATE SEQUENCE "Books_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE "Books_id_seq" OWNED BY "Books".id;



CREATE TABLE "Chapter_Paths" (
    id integer NOT NULL,
    from_chapter_id integer,
    to_chapter_id integer,
    "createdAt" timestamp with time zone NOT NULL,
    "updatedAt" timestamp with time zone NOT NULL
);

CREATE SEQUENCE "Chapter_Paths_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE "Chapter_Paths_id_seq" OWNED BY "Chapter_Paths".id;

CREATE TABLE "Chapters" (
    id integer NOT NULL,
    title text NOT NULL,
    body text NOT NULL,
    book_id integer NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "updatedAt" timestamp with time zone NOT NULL,
    "deletedAt" timestamp with time zone,
    creator_id integer
);

CREATE SEQUENCE "Chapters_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE "Chapters_id_seq" OWNED BY "Chapters".id;

CREATE TABLE "Users" (
    id integer NOT NULL,
    username text,
    password text,
    email text,
    "createdAt" timestamp with time zone NOT NULL,
    "updatedAt" timestamp with time zone NOT NULL,
    "deletedAt" timestamp with time zone
);

CREATE SEQUENCE "Users_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE "Users_id_seq" OWNED BY "Users".id;

ALTER TABLE ONLY "Books" ALTER COLUMN id SET DEFAULT nextval('"Books_id_seq"'::regclass);

ALTER TABLE ONLY "Chapter_Paths" ALTER COLUMN id SET DEFAULT nextval('"Chapter_Paths_id_seq"'::regclass);

ALTER TABLE ONLY "Chapters" ALTER COLUMN id SET DEFAULT nextval('"Chapters_id_seq"'::regclass);

ALTER TABLE ONLY "Users" ALTER COLUMN id SET DEFAULT nextval('"Users_id_seq"'::regclass);

ALTER TABLE ONLY "Books"
    ADD CONSTRAINT "Books_pkey" PRIMARY KEY (id);

ALTER TABLE ONLY "Chapter_Paths"
    ADD CONSTRAINT "Chapter_Paths_pkey" PRIMARY KEY (id);

ALTER TABLE ONLY "Chapters"
    ADD CONSTRAINT "Chapters_pkey" PRIMARY KEY (id);

ALTER TABLE ONLY "Users"
    ADD CONSTRAINT "Users_pkey" PRIMARY KEY (id);

CREATE UNIQUE INDEX "Unique_From_Chapter_To_Chapter" ON "Chapter_Paths" USING btree (from_chapter_id, to_chapter_id);

CREATE UNIQUE INDEX user_unique_email ON "Users" USING btree (email);

CREATE UNIQUE INDEX user_unique_username ON "Users" USING btree (username);

ALTER TABLE ONLY "Books"
    ADD CONSTRAINT "Books_creator_id_fkey" FOREIGN KEY (creator_id) REFERENCES "Users"(id) ON UPDATE CASCADE ON DELETE SET NULL;

ALTER TABLE ONLY "Books"
    ADD CONSTRAINT "Books_first_chapter_id_fkey" FOREIGN KEY (first_chapter_id) REFERENCES "Chapters"(id);

ALTER TABLE ONLY "Chapter_Paths"
    ADD CONSTRAINT "Chapter_Paths_from_chapter_id_fkey" FOREIGN KEY (from_chapter_id) REFERENCES "Chapters"(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY "Chapter_Paths"
    ADD CONSTRAINT "Chapter_Paths_to_chapter_id_fkey" FOREIGN KEY (to_chapter_id) REFERENCES "Chapters"(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY "Chapters"
    ADD CONSTRAINT "Chapters_book_id_fkey" FOREIGN KEY (book_id) REFERENCES "Books"(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY "Chapters"
    ADD CONSTRAINT "Chapters_creator_id_fkey" FOREIGN KEY (creator_id) REFERENCES "Users"(id) ON UPDATE CASCADE ON DELETE SET NULL;

-- +goose StatementEnd


-- +goose Down