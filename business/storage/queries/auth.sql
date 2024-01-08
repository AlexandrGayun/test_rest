-- name: GetApiKeyID :one
SELECT id FROM auth
WHERE api_key = ?;