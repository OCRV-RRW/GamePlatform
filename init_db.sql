--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4
-- Dumped by pg_dump version 17.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: game; Type: TABLE; Schema: public; Owner: test
--

CREATE TABLE public.game (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    title character varying(100) NOT NULL,
    description text NOT NULL,
    src character varying(200) NOT NULL,
    icon character varying(200) NOT NULL,
    created timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.game OWNER TO test;

--
-- Name: game_preview; Type: TABLE; Schema: public; Owner: test
--

CREATE TABLE public.game_preview (
    game_id uuid NOT NULL,
    preview_id uuid NOT NULL
);


ALTER TABLE public.game_preview OWNER TO test;

--
-- Name: preview; Type: TABLE; Schema: public; Owner: test
--

CREATE TABLE public.preview (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    image character varying(200) NOT NULL,
    video character varying(200)
);


ALTER TABLE public.preview OWNER TO test;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: test
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO test;

--
-- Data for Name: game; Type: TABLE DATA; Schema: public; Owner: test
--

INSERT INTO public.game (id, title, description, src, icon, created) VALUES ('27421cdd-a91d-487a-9887-e881f4d97bae', 'Сортировочная станция', 'Рассортируй поступающие на станцию вагоны по цветам депо', 'https://s3.ocrv-game.ru/games/Warehouse/index.html', 'https://s3.ocrv-game.ru/platform/games/WagonsTracks/preview.png', '2025-07-15 11:16:22.554032');
INSERT INTO public.game (id, title, description, src, icon, created) VALUES ('697cae11-1219-41a6-bd6c-78d42ed996e4', 'Карманный склад', 'Управляйте грузопотоками на собственной железнодорожной станции и получайте вознаграждение за успехи', 'https://s3.ocrv-game.ru/games/Warehouse/index.html', 'https://s3.ocrv-game.ru/games/Warehouse/StreamingAssets/pocket-preview.png', '2025-01-20 07:28:18');
INSERT INTO public.game (id, title, description, src, icon, created) VALUES ('d35a178b-e8a3-4005-b024-8fa7175912db', 'Логическое депо', 'Размещайте, удаляйте и изменяйте маршрут соединений таким образом, чтобы вагоны безопасно подключались к локомотиву. Но будьте осторожны и не заставьте их столкнуться друг с другом!', 'https://s3.ocrv-game.ru/games/games_gzip/rail-bound/index.html', 'https://s3.ocrv-game.ru/games/games_gzip/rail-bound/StreamingAssets/preveiw-logic.png', '2025-01-21 14:35:18');
INSERT INTO public.game (id, title, description, src, icon, created) VALUES ('1d0d9b5c-d1d4-402f-8b2e-b7e29b1e6944', 'Найди фигуру', 'Найди едиственнную уникальную фигуру из множества похожих', 'https://s3.ocrv-game.ru/games/figure_search_no_platfrom/index.html', 'https://s3.ocrv-game.ru/platform/games/figure_search/preview.png', '2025-01-22 12:00:00');
INSERT INTO public.game (id, title, description, src, icon, created) VALUES ('7fd25082-396b-4e8e-82c9-dec04f36fa4b', 'Мнемотренажер', 'Запомни фигуру и повтори', 'https://s3.ocrv-game.ru/games/RememberMe/index.html', 'https://s3.ocrv-game.ru/platform/games/RememberMe/preview.png', '2025-01-24 21:10:30');


--
-- Data for Name: game_preview; Type: TABLE DATA; Schema: public; Owner: test
--

INSERT INTO public.game_preview (game_id, preview_id) VALUES ('27421cdd-a91d-487a-9887-e881f4d97bae', '7724dc47-913e-4864-96a5-8e2d10033077');
INSERT INTO public.game_preview (game_id, preview_id) VALUES ('27421cdd-a91d-487a-9887-e881f4d97bae', '205b5297-d6f7-4a71-b53d-e2949e11bb59');
INSERT INTO public.game_preview (game_id, preview_id) VALUES ('7fd25082-396b-4e8e-82c9-dec04f36fa4b', '610f957e-f1cb-4134-adb9-b7f4697e073c');


--
-- Data for Name: preview; Type: TABLE DATA; Schema: public; Owner: test
--

INSERT INTO public.preview (id, image, video) VALUES ('7724dc47-913e-4864-96a5-8e2d10033077', 'https://s3.ocrv-game.ru/games/Warehouse/StreamingAssets/pocket-preview.png', NULL);
INSERT INTO public.preview (id, image, video) VALUES ('205b5297-d6f7-4a71-b53d-e2949e11bb59', 'https://s3.ocrv-game.ru/platform/games/WagonsTracks/preview.png', 'https://s3.ocrv-game.ru/platform/games/WagonsTracks/video.mp4');
INSERT INTO public.preview (id, image, video) VALUES ('610f957e-f1cb-4134-adb9-b7f4697e073c', 'https://s3.ocrv-game.ru/platform/games/RememberMe/preview.png', 'https://s3.ocrv-game.ru/platform/games/RememberMe/video.mp4');


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: test
--

INSERT INTO public.schema_migrations (version, dirty) VALUES (1, false);


--
-- Name: game game_pkey; Type: CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.game
    ADD CONSTRAINT game_pkey PRIMARY KEY (id);


--
-- Name: game_preview game_preview_pkey; Type: CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.game_preview
    ADD CONSTRAINT game_preview_pkey PRIMARY KEY (game_id, preview_id);


--
-- Name: preview preview_pkey; Type: CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.preview
    ADD CONSTRAINT preview_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: game_preview game_preview_game_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.game_preview
    ADD CONSTRAINT game_preview_game_id_fkey FOREIGN KEY (game_id) REFERENCES public.game(id) ON DELETE CASCADE;


--
-- Name: game_preview game_preview_preview_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: test
--

ALTER TABLE ONLY public.game_preview
    ADD CONSTRAINT game_preview_preview_id_fkey FOREIGN KEY (preview_id) REFERENCES public.preview(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

