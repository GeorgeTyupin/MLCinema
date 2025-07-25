--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4
-- Dumped by pg_dump version 17.4

-- Started on 2025-07-25 04:30:15

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 218 (class 1259 OID 41751)
-- Name: actors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.actors (
    id bigint NOT NULL,
    name text NOT NULL,
    birth_year bigint,
    nationality text
);


ALTER TABLE public.actors OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 41750)
-- Name: actors_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.actors_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.actors_id_seq OWNER TO postgres;

--
-- TOC entry 4888 (class 0 OID 0)
-- Dependencies: 217
-- Name: actors_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.actors_id_seq OWNED BY public.actors.id;


--
-- TOC entry 223 (class 1259 OID 41784)
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id bigint NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 41783)
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_id_seq OWNER TO postgres;

--
-- TOC entry 4889 (class 0 OID 0)
-- Dependencies: 222
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- TOC entry 221 (class 1259 OID 41768)
-- Name: film_actors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.film_actors (
    film_id bigint NOT NULL,
    actor_id bigint NOT NULL
);


ALTER TABLE public.film_actors OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 41792)
-- Name: film_categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.film_categories (
    category_id bigint NOT NULL,
    film_id bigint NOT NULL
);


ALTER TABLE public.film_categories OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 41760)
-- Name: films; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.films (
    id bigint NOT NULL,
    title text NOT NULL,
    year bigint,
    country text,
    image_path text,
    description text
);


ALTER TABLE public.films OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 41759)
-- Name: films_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.films_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.films_id_seq OWNER TO postgres;

--
-- TOC entry 4890 (class 0 OID 0)
-- Dependencies: 219
-- Name: films_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.films_id_seq OWNED BY public.films.id;


--
-- TOC entry 4713 (class 2604 OID 41754)
-- Name: actors id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.actors ALTER COLUMN id SET DEFAULT nextval('public.actors_id_seq'::regclass);


--
-- TOC entry 4715 (class 2604 OID 41787)
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- TOC entry 4714 (class 2604 OID 41763)
-- Name: films id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.films ALTER COLUMN id SET DEFAULT nextval('public.films_id_seq'::regclass);


--
-- TOC entry 4876 (class 0 OID 41751)
-- Dependencies: 218
-- Data for Name: actors; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.actors (id, name, birth_year, nationality) FROM stdin;
1	Мэттью МакКонахи	0	
2	Энн Хэтэуэй	0	
3	Леонардо Ди Каприо	0	
4	Том Харди	0	
5	Киану Ривз	0	
6	Кэрри-Энн Мосс	0	
7	Брэд Питт	0	
8	Эдвард Нортон	0	
9	Том Хэнкс	0	
10	Тим Роббинс	0	
11	Морган Фримен	0	
12	Кристиан Бейл	0	
13	Хит Леджер	0	
14	Рассел Кроу	0	
15	Майкл Кларк Дункан	0	
16	Франсуа Клюзе	0	
17	Омар Си	0	
18	Жан Рено	0	
19	Натали Портман	0	
20	Кейт Уинслет	0	
21	Джонатан Тейлор Томас	0	
22	Элайджа Вуд	0	
23	Иэн Маккеллен	0	
24	Джон Траволта	0	
25	Сэмюэл Л. Джексон	0	
26	Хью Джекман	0	
27	Джесси Айзенберг	0	
28	Эндрю Гарфилд	0	
29	Хоакин Феникс	0	
30	Райан Гослинг	0	
31	Харрисон Форд	0	
\.


--
-- TOC entry 4881 (class 0 OID 41784)
-- Dependencies: 223
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (id, name) FROM stdin;
1	Фантастика
2	Триллер
3	Драма
4	Биография
5	Боевик
6	История
7	Фэнтези
8	Комедия
9	Мелодрама
10	Мультфильм
11	Криминал
\.


--
-- TOC entry 4879 (class 0 OID 41768)
-- Dependencies: 221
-- Data for Name: film_actors; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.film_actors (film_id, actor_id) FROM stdin;
1	1
1	2
2	3
2	4
3	5
3	6
4	7
4	8
5	9
6	10
6	11
7	12
7	13
8	14
9	9
9	15
10	16
10	17
11	3
12	18
12	19
13	3
13	20
14	21
15	22
15	23
16	24
16	25
17	12
17	26
18	27
18	28
19	29
20	30
20	31
\.


--
-- TOC entry 4882 (class 0 OID 41792)
-- Dependencies: 224
-- Data for Name: film_categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.film_categories (category_id, film_id) FROM stdin;
1	1
1	2
2	2
1	3
3	4
3	5
4	5
3	6
5	7
6	8
5	8
3	9
7	9
3	10
8	10
2	11
5	12
3	12
9	13
10	14
7	15
11	16
3	17
2	17
3	18
4	18
2	19
1	20
\.


--
-- TOC entry 4878 (class 0 OID 41760)
-- Dependencies: 220
-- Data for Name: films; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.films (id, title, year, country, image_path, description) FROM stdin;
1	Интерстеллар	2014	США	/static/img/interstellar.jpg	Будущее Земли под угрозой. Астронавты ищут новую планету.
2	Начало	2010	США	/static/img/inception.jpg	Проникновение в сны и внедрение идей.
3	Матрица	1999	США	/static/img/matrix.jpg	Реальность — это симуляция.
4	Бойцовский клуб	1999	США	/static/img/fight_club.jpg	Мужчина создаёт подпольный бойцовский клуб.
5	Форрест Гамп	1994	США	/static/img/forrest_gump.jpg	История жизни Форреста.
6	Побег из Шоушенка	1994	США	/static/img/shawshank.jpg	Заключённый планирует побег.
7	Темный рыцарь	2008	США	/static/img/dark_knight.jpg	Бэтмен против Джокера.
8	Гладиатор	2000	США	/static/img/gladiator.jpg	Генерал становится гладиатором.
9	Зеленая миля	1999	США	/static/img/green_mile.jpg	Сверхъестественные способности заключённого.
10	1+1	2011	Франция	/static/img/intouchables.jpg	Инвалид и сиделка становятся друзьями.
11	Остров проклятых	2010	США	/static/img/shutter_island.jpg	Маршалы США расследуют исчезновение пациентки.
12	Леон	1994	Франция	/static/img/leon.jpg	Киллер берёт под опеку девочку.
13	Титаник	1997	США	/static/img/titanic.jpg	История любви на лайнере.
14	Король Лев	1994	США	/static/img/lion_king.jpg	Молодой лев Симба ищет своё место.
15	Властелин колец: Братство кольца	2001	Новая Зеландия	/static/img/lotr_fellowship.jpg	Фродо отправляется уничтожить Кольцо.
16	Криминальное чтиво	1994	США	/static/img/pulp_fiction.jpg	Истории мафии в Лос-Анджелесе.
17	Престиж	2006	США	/static/img/prestige.jpg	Противостояние двух фокусников.
18	Социальная сеть	2010	США	/static/img/social_network.jpg	Создание Facebook и конфликты вокруг него.
19	Джокер	2019	США	/static/img/joker.jpg	История становления Джокера.
20	Бегущий по лезвию 2049	2017	США	/static/img/blade_runner_2049.jpg	Офицер Кей раскрывает тайну, угрожающую обществу.
\.


--
-- TOC entry 4891 (class 0 OID 0)
-- Dependencies: 217
-- Name: actors_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.actors_id_seq', 31, true);


--
-- TOC entry 4892 (class 0 OID 0)
-- Dependencies: 222
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_id_seq', 11, true);


--
-- TOC entry 4893 (class 0 OID 0)
-- Dependencies: 219
-- Name: films_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.films_id_seq', 20, true);


--
-- TOC entry 4717 (class 2606 OID 41758)
-- Name: actors actors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.actors
    ADD CONSTRAINT actors_pkey PRIMARY KEY (id);


--
-- TOC entry 4723 (class 2606 OID 41791)
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- TOC entry 4721 (class 2606 OID 41772)
-- Name: film_actors film_actors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.film_actors
    ADD CONSTRAINT film_actors_pkey PRIMARY KEY (film_id, actor_id);


--
-- TOC entry 4725 (class 2606 OID 41796)
-- Name: film_categories film_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.film_categories
    ADD CONSTRAINT film_categories_pkey PRIMARY KEY (category_id, film_id);


--
-- TOC entry 4719 (class 2606 OID 41767)
-- Name: films films_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.films
    ADD CONSTRAINT films_pkey PRIMARY KEY (id);


--
-- TOC entry 4726 (class 2606 OID 41778)
-- Name: film_actors fk_film_actors_actor; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.film_actors
    ADD CONSTRAINT fk_film_actors_actor FOREIGN KEY (actor_id) REFERENCES public.actors(id);


--
-- TOC entry 4727 (class 2606 OID 41773)
-- Name: film_actors fk_film_actors_film; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.film_actors
    ADD CONSTRAINT fk_film_actors_film FOREIGN KEY (film_id) REFERENCES public.films(id);


--
-- TOC entry 4728 (class 2606 OID 41797)
-- Name: film_categories fk_film_categories_category; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.film_categories
    ADD CONSTRAINT fk_film_categories_category FOREIGN KEY (category_id) REFERENCES public.categories(id);


--
-- TOC entry 4729 (class 2606 OID 41802)
-- Name: film_categories fk_film_categories_film; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.film_categories
    ADD CONSTRAINT fk_film_categories_film FOREIGN KEY (film_id) REFERENCES public.films(id);


-- Completed on 2025-07-25 04:30:15

--
-- PostgreSQL database dump complete
--

