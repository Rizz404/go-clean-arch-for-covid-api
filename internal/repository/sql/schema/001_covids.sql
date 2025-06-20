-- +goose Up
CREATE TABLE
  covids (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nama TEXT NOT NULL,
    kota TEXT NOT NULL,
    sembuh INTEGER NOT NULL DEFAULT 0,
    dirawat INTEGER NOT NULL DEFAULT 0,
    meninggal INTEGER NOT NULL DEFAULT 0,
    total INTEGER NOT NULL DEFAULT 0
  );

-- +goose Down
DROP TABLE covids;
