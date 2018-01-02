ALTER TABLE book_tags RENAME TO btags_votes;
ALTER TABLE btags_votes
  ADD COLUMN direction SMALLINT DEFAULT 0;

ALTER TABLE book_tags_count RENAME TO btags_vote_count;