-- +goose Up
CREATE TABLE
  covids (
    id UUID PRIMARY KEY,
    nama STRING NOT NULL,
    kota STRING NOT NULL,
    sembuh INT NOT NULL DEFAULT 0,
    dirawat INT NOT NULL DEFAULT 0,
    meninggal INT NOT NULL DEFAULT 0,
    total INT NOT NULL DEFAULT 0
  );

-- +goose Down
DROP TABLE covids;
