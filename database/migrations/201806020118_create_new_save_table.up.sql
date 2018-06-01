/*just recreate the table since we haven't used it yet at this point*/
CREATE TABLE save_states (
  id          SERIAL PRIMARY KEY,
  state       JSON, /* in the shape of { history: [{health:10}, {health: 5},...], cursor: 1 } */

  created_by  INT REFERENCES users(id) NULL, /* nullable since guests can also create story saves */
  book_id     INT REFERENCES books(id),

  created_at  TIMESTAMP WITHOUT TIME ZONE,
  updated_at  TIMESTAMP WITHOUT TIME ZONE
)