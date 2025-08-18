create extension if not exists "uuid-ossp";

create schema platform;

create table platform.user(
	id                uuid primary key default uuid_generate_v4(),
	name              varchar(200) unique not null,
	email             varchar(200) not null unique,
	password          varchar(200) not null,
	is_admin          boolean not null default false ,
	verification_code varchar(200) not null,
	verified          boolean not null,
	birthday          date,
	gender            varchar(100),
	created_at           timestamp not null default now()
);

create table platform.game (
	id          uuid primary key default uuid_generate_v4(),
	title       varchar(100) not null,
	description text not null,
	src         varchar(200) not null,
	icon        varchar(200) not null,
	created     timestamp not null default now()
);

create table platform.preview (
	id    uuid primary key default uuid_generate_v4(),
	image varchar(200) not null,
	video varchar(200)
);

create table platform.game_preview (
	game_id    uuid references platform.game on delete cascade,
	preview_id uuid references platform.preview on delete cascade,
	primary key (game_id, preview_id)
);
