ALTER TABLE book_tags DROP CONSTRAINT "Book_Tags_user_id_fkey";
ALTER TABLE book_tags
  ADD CONSTRAINT "Book_Tags_user_id_fkey" FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE;