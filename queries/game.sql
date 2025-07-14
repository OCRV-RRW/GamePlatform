-- name: GetGameByID :one
select id, title, description, src, preview, created from games
where id = $1;

-- name: GetGames :many
select id, title, description, src, preview, created from games
order by created desc;
