
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE notes (
  id serial NOT NULL,
  title text,
  created_at timestamp,
  updated_at timestamp,
  PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE notes;
