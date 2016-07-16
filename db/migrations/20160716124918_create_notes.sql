
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE notes (
  id int NOT NULL,
  title text,
  PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE notes;
