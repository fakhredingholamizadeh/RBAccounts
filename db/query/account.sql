-- name: CreateAccount :one
INSERT INTO  account(
  username,
  hashed_password,
  full_name,
  email,
  password_changed_at
  ) VALUES (
  $1, $2, $3, $4, $5
)RETURNING *;

-- name: GetAccount :one
SELECT * FROM account
WHERE id=$1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM account
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccountFullName :one
UPDATE account 
SET full_name=$2
WHERE id = $1
RETURNING *;

-- name: UpdateAccountEmail :one
UPDATE account 
SET email=$2
WHERE id = $1
RETURNING *;

-- name: UpdateAccountHashedPassword :one
UPDATE account 
SET hashed_password=$2
WHERE id = $1
RETURNING *;


-- name: DeleteAccountEx :exec
DELETE FROM account
WHERE id = $1;

-- name: DeleteAccountOne :one
DELETE FROM account
WHERE id = $1
RETURNING *;
