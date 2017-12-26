CREATE TABLE book_tags (
  id bigserial,
  tag text,
  book_id integer,
  user_id integer
);
ALTER TABLE book_tags
  ADD CONSTRAINT "Book_Tags_book_id_fkey" FOREIGN KEY (book_id) REFERENCES books(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE book_tags
  ADD CONSTRAINT "Book_Tags_user_id_fkey" FOREIGN KEY (book_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE;


/* For caching book tags */
CREATE TABLE book_tags_count (
  id serial,
  tag text,
  book_id integer,
  count integer
);

ALTER TABLE book_tags_count
  ADD CONSTRAINT "Book_Tags_Count_book_id_fkey" FOREIGN KEY (book_id) REFERENCES books(id) ON UPDATE CASCADE ON DELETE CASCADE;
