-- name: GetGameByID :one
select g.id, title, description, src, icon, created from platform.game g
where g.id = $1;

-- name: GetAllGames :many
select id, title, description, src, icon, created from platform.game g
order by created desc;

-- name: CreateGame :one
insert into platform.game(title, description, src, icon)
values ($1, $2, $3, $4)
returning *;

-- name: GetGamePreview :many
select p.id, p.image, p.video from platform.game_preview gp
join platform.preview p on p.id = gp.preview_id
where gp.game_id = $1;

-- name: UpdateGame :exec
update platform.game set
	title = $2,
	description = $3,
	src = $4
where id = $1;

-- name: UpdateGameIcon :exec
update platform.game set
	icon = $2
where id = $1;

-- name: DeleteGame :exec
delete from platform.game
where id = $1;

-- name: GetPreviewByID :one
select id, image, video from platform.preview p
where id = $1;

-- name: GetGamePreviewByID :many
select image, video from platform.game g
join platform.game_preview gp on gp.id = g.id
join platform.preview p on p.id = gp.id
where g.id = $1
order by p.video;

-- name: CreatePreview :one
with preview_insert as (
insert into platform.preview(image, video)
values ($2, $3)
returning *
)
insert into platform.game_preview (game_id, preview_id)
select $1, pi.id from preview_insert pi
returning
	(SELECT id FROM preview_insert) AS id,
    (SELECT image FROM preview_insert) AS image,
    (SELECT video FROM preview_insert) AS video;

-- name: DeletePreview :exec
delete from platform.preview
where id = $1;
