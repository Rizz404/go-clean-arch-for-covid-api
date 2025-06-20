-- name: CreateCovid :one
INSERT INTO
  covids (nama, kota, sembuh, dirawat, meninggal, total)
VALUES
  (?, ?, ?, ?, ?, ?) RETURNING *;

-- name: GetCovid :one
SELECT
  *
FROM
  covids
WHERE
  id = ?
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
  nama = ?,
  kota = ?,
  sembuh = ?,
  dirawat = ?,
  meninggal = ?,
  total = ?
WHERE
  id = ? RETURNING *;

-- name: DeleteCovid :exec
DELETE FROM covids
WHERE
  id = ?;
