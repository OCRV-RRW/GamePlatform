create extension if not exists "uuid-ossp";

create table game (
	id uuid primary key default uuid_generate_v4(),
	title varchar(100) not null,
	description text not null,
	src varchar(200) not null,
	icon varchar(200) not null,
	created timestamp not null default now()
);

create table preview (
	id uuid primary key default uuid_generate_v4(),
	image varchar(200) not null,
	video varchar(200)
);

create table game_preview (
	game_id uuid references game on delete cascade,
	preview_id uuid references preview on delete cascade,
	primary key (game_id, preview_id)
);
