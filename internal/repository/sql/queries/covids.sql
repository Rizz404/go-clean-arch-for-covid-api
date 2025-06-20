-- name: CreateCovid :one
INSERT INTO
  covids (id, nama, kota, sembuh, dirawat, meninggal, total)
VALUES
  ($1, $2, $3, $4, $5, $6, $7)
RETURNING
  *;

-- name: GetCovid :one
SELECT
  *
FROM
  covids
WHERE
  id = $1
LIMIT
  1;

-- name: GetCovids :many
SELECT
  *
FROM
  covids
ORDER BY
  id;

-- name: UpdateCovid :one
UPDATE covids
SET
  nama = $1,
  kota = $2,
  sembuh = $3,
  dirawat = $4,
  meninggal = $5,
  total = $6
WHERE
  id = $7
RETURNING
  *;

-- name: DeleteCovid :exec
DELETE FROM covids
WHERE
  id = $1;
