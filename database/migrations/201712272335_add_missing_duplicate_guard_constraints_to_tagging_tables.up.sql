
/* Ensures a single user can't tag a book more than once with the same tag */
ALTER TABLE book_tags
  ADD CONSTRAINT "Book_Tags_unique_cols_book_id_user_id_tag"
  UNIQUE(book_id, user_id,"tag");


/* Ensures there isn't more than one Count entry for a given tag in a book */
ALTER TABLE book_tags_count
  ADD CONSTRAINT "Book_Tags_Count_unique_cols_book_id_tag"
  UNIQUE(book_id, "tag");