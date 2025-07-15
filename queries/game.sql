-- name: GetGameByID :one
select g.id, title, description, src, icon, created from game g
where g.id = $1;

-- name: GetGames :many
select id, title, description, src, icon, created from game
order by created desc;

-- name: GetGamePreview :many
select p.id, p.image, p.video from game_preview gp
join preview p on p.id = gp.preview_id
where gp.game_id = $1;
