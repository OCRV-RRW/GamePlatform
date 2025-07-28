-- name: GetUserByID :one
select * from platform.user u
where u.id = $1;

-- name: GetAllUsers :many
select * from platform.user u;

-- name: GetUserByEmail :one
select * from platform.user u
where u.email = $1;

-- name: GetUserByVerificationCode :one
select * from platform.user u
where u.verification_code = $1;

-- name: CreateUser :one
insert into platform.user(name, email, password, is_admin, verification_code, verified, birthday, gender)
values ($1, $2, $3, $4, $5, $6, $7, $8)
returning *;

-- name: UpdateUserById :exec
update platform.user set
	name = $2,
	email = $3,
	password  = $4,
	is_admin = $5,
	verified = $6,
	birthday = $7,
	gender = $8
where id = $1;

-- name: UpdateUserVerification :exec
update platform.user set
	verified = $2,
	verification_code = $3
where id  = $1;

-- name: UpdateUserPassword :exec
update platform.user u set
	password = $2
where id = $1;

-- name: DeleteUserById :exec
delete from platform.user u
where u.id = $1;
