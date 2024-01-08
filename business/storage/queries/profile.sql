-- name: GetProfiles :many
SELECT user.id, user.username, ud.school, up.first_name, up.last_name, up.phone, up.address, up.city FROM user
LEFT JOIN user_data AS ud ON user.id = ud.user_id
LEFT JOIN user_profile AS up ON user.id = up.user_id;

-- name: GetProfileByUsername :one
SELECT user.id, user.username, ud.school, up.first_name, up.last_name, up.phone, up.address, up.city FROM user
LEFT JOIN user_data AS ud ON user.id = ud.user_id
LEFT JOIN user_profile AS up ON user.id = up.user_id
WHERE user.username = sqlc.arg(username)
LIMIT 1;