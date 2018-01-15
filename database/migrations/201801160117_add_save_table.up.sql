CREATE TABLE story_saves (
  token TEXT PRIMARY KEY,
  save  JSON,

  created_at timestamp without time zone,
  created_by INT REFERENCES users(id) NULL, /* nullable since guests can also create story saves */

  book_id INT REFERENCES books(id),
  chapter_id INT REFERENCES chapters(id)
)