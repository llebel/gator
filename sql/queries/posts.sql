-- name: CreatePost :one
INSERT INTO posts (title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPostsForUser :many
SELECT
    posts.*
FROM posts
INNER JOIN feeds ON posts.feed_id = feeds.id
INNER JOIN feed_follows ON feeds.id = feed_follows.feed_id
INNER JOIN users ON users.id = feed_follows.user_id
WHERE users.name = $1
ORDER BY posts.published_at DESC
LIMIT $2 OFFSET $3
;
