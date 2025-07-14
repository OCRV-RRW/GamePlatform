create extension if not exists "uuid-ossp";

create table games (
	id uuid primary key default uuid_generate_v4(),
	title varchar(100) not null,
	description text not null,
	src varchar(200) not null,
	preview varchar(200) not null,
	created timestamp not null default now()
);
--
-- CREATE TABLE `games` (
--   id int NOT NULL AUTO_INCREMENT,
--   `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
--   `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
--   `src` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
--   `preview` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
--   `created` datetime NOT NULL,
--   PRIMARY KEY (`id`)
-- ) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
--
-- CREATE TABLE IF NOT EXISTS public.games (
--     id uuid NOT NULL,
--
--     username character varying(50) COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
--     given_name character varying(50) COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
--     family_name character varying(50) COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
--     enabled boolean NOT NULL DEFAULT false,
--     CONSTRAINT user_pkey PRIMARY KEY (id),
--     CONSTRAINT user_username_key UNIQUE (username)
-- );
