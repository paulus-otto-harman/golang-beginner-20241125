--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Homebrew)
-- Dumped by pg_dump version 16.4 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
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
-- Name: addresses; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.addresses (
    id smallint NOT NULL,
    customer_id smallint NOT NULL,
    fullname character varying(30) NOT NULL,
    email character varying(30) NOT NULL,
    detail character varying(255) NOT NULL,
    is_default boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.addresses OWNER TO paul;

--
-- Name: addresses_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.addresses_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.addresses_id_seq OWNER TO paul;

--
-- Name: addresses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.addresses_id_seq OWNED BY public.addresses.id;


--
-- Name: order_items; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.order_items (
    id smallint NOT NULL,
    order_id smallint NOT NULL,
    product_id smallint NOT NULL,
    quantity smallint NOT NULL,
    unit_price money NOT NULL,
    discount_rate numeric DEFAULT 0 NOT NULL,
    net_price money NOT NULL
);


ALTER TABLE public.order_items OWNER TO paul;

--
-- Name: reviews; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.reviews (
    id smallint NOT NULL,
    order_item_id smallint NOT NULL,
    rating smallint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.reviews OWNER TO paul;

--
-- Name: average_product_reviews; Type: VIEW; Schema: public; Owner: paul
--

CREATE VIEW public.average_product_reviews AS
 SELECT order_items.product_id,
    COALESCE(avg(reviews.rating), (0)::numeric) AS rating
   FROM (public.order_items
     LEFT JOIN public.reviews ON ((order_items.id = reviews.order_item_id)))
  GROUP BY order_items.product_id;


ALTER VIEW public.average_product_reviews OWNER TO paul;

--
-- Name: banners; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.banners (
    id smallint NOT NULL,
    photo character varying(255) NOT NULL,
    title character varying(60),
    subtitle character varying(60),
    path_page character varying(255),
    started_at timestamp with time zone,
    finished_at timestamp with time zone,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.banners OWNER TO paul;

--
-- Name: banners_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.banners_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.banners_id_seq OWNER TO paul;

--
-- Name: banners_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.banners_id_seq OWNED BY public.banners.id;


--
-- Name: cart_items; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.cart_items (
    id smallint NOT NULL,
    shopping_session_id smallint NOT NULL,
    product_id smallint NOT NULL,
    quantity smallint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone
);


ALTER TABLE public.cart_items OWNER TO paul;

--
-- Name: cart_items_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.cart_items_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cart_items_id_seq OWNER TO paul;

--
-- Name: cart_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.cart_items_id_seq OWNED BY public.cart_items.id;


--
-- Name: categories; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.categories (
    id smallint NOT NULL,
    name character varying(25) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.categories OWNER TO paul;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.categories_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_id_seq OWNER TO paul;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: customers; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.customers (
    id smallint NOT NULL,
    user_id smallint NOT NULL,
    name character varying(30) NOT NULL
);


ALTER TABLE public.customers OWNER TO paul;

--
-- Name: sessions; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.sessions (
    token uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id smallint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    expired_at timestamp with time zone
);


ALTER TABLE public.sessions OWNER TO paul;

--
-- Name: customer_tokens; Type: VIEW; Schema: public; Owner: paul
--

CREATE VIEW public.customer_tokens AS
 SELECT sessions.token,
    customers.id AS customer_id
   FROM (public.sessions
     JOIN public.customers ON ((sessions.user_id = customers.user_id)))
  WHERE (sessions.expired_at >= now());


ALTER VIEW public.customer_tokens OWNER TO paul;

--
-- Name: customers_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.customers_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.customers_id_seq OWNER TO paul;

--
-- Name: customers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.customers_id_seq OWNED BY public.customers.id;


--
-- Name: order_items_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.order_items_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.order_items_id_seq OWNER TO paul;

--
-- Name: order_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.order_items_id_seq OWNED BY public.order_items.id;


--
-- Name: orders; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.orders (
    id smallint NOT NULL,
    customer_id smallint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    address_id smallint,
    coupon_code character varying(20),
    shipping character varying(10),
    notes character varying(128),
    payment_method character varying(10)
);


ALTER TABLE public.orders OWNER TO paul;

--
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.orders_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.orders_id_seq OWNER TO paul;

--
-- Name: orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;


--
-- Name: product_photos; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.product_photos (
    id smallint NOT NULL,
    product_id smallint NOT NULL,
    photo_url character varying(255) NOT NULL
);


ALTER TABLE public.product_photos OWNER TO paul;

--
-- Name: product_photos_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.product_photos_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_photos_id_seq OWNER TO paul;

--
-- Name: product_photos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.product_photos_id_seq OWNED BY public.product_photos.id;


--
-- Name: product_variants; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.product_variants (
    id smallint NOT NULL,
    product_id smallint NOT NULL,
    name character varying(15) NOT NULL
);


ALTER TABLE public.product_variants OWNER TO paul;

--
-- Name: product_variants_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.product_variants_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_variants_id_seq OWNER TO paul;

--
-- Name: product_variants_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.product_variants_id_seq OWNED BY public.product_variants.id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.products (
    id smallint NOT NULL,
    name character varying(40) NOT NULL,
    thumbnail character varying(255) NOT NULL,
    price money,
    discount_rate numeric(5,2),
    category_id smallint,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.products OWNER TO paul;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.products_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.products_id_seq OWNER TO paul;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: recommendations; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.recommendations (
    id smallint NOT NULL,
    title character varying(48),
    subtitle character varying(128),
    photo character varying(255),
    product_id smallint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.recommendations OWNER TO paul;

--
-- Name: recommendations_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.recommendations_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.recommendations_id_seq OWNER TO paul;

--
-- Name: recommendations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.recommendations_id_seq OWNED BY public.recommendations.id;


--
-- Name: reviews_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.reviews_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.reviews_id_seq OWNER TO paul;

--
-- Name: reviews_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.reviews_id_seq OWNED BY public.reviews.id;


--
-- Name: shopping_sessions; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.shopping_sessions (
    id smallint NOT NULL,
    customer_id smallint NOT NULL
);


ALTER TABLE public.shopping_sessions OWNER TO paul;

--
-- Name: shopping_session_tokens; Type: VIEW; Schema: public; Owner: paul
--

CREATE VIEW public.shopping_session_tokens AS
 SELECT customer_tokens.token,
    shopping_sessions.id AS shopping_session_id
   FROM (public.customer_tokens
     LEFT JOIN public.shopping_sessions ON ((customer_tokens.customer_id = shopping_sessions.customer_id)));


ALTER VIEW public.shopping_session_tokens OWNER TO paul;

--
-- Name: shopping_sessions_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.shopping_sessions_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.shopping_sessions_id_seq OWNER TO paul;

--
-- Name: shopping_sessions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.shopping_sessions_id_seq OWNED BY public.shopping_sessions.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.users (
    id smallint NOT NULL,
    username character varying(40) NOT NULL,
    password character varying(16) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO paul;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.user_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_id_seq OWNER TO paul;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.user_id_seq OWNED BY public.users.id;


--
-- Name: weeklies; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.weeklies (
    id smallint NOT NULL,
    photo character varying(255),
    title character varying(48),
    subtitle character varying(128),
    product_id smallint NOT NULL,
    started_at timestamp with time zone NOT NULL,
    finished_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.weeklies OWNER TO paul;

--
-- Name: weeklies_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.weeklies_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.weeklies_id_seq OWNER TO paul;

--
-- Name: weeklies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.weeklies_id_seq OWNED BY public.weeklies.id;


--
-- Name: wishlists; Type: TABLE; Schema: public; Owner: paul
--

CREATE TABLE public.wishlists (
    id smallint NOT NULL,
    customer_id smallint NOT NULL,
    product_id smallint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.wishlists OWNER TO paul;

--
-- Name: wishlist_tokens; Type: VIEW; Schema: public; Owner: paul
--

CREATE VIEW public.wishlist_tokens AS
 SELECT customer_tokens.token,
    wishlists.product_id
   FROM (public.customer_tokens
     LEFT JOIN public.wishlists ON ((customer_tokens.customer_id = wishlists.customer_id)));


ALTER VIEW public.wishlist_tokens OWNER TO paul;

--
-- Name: wishlists_id_seq; Type: SEQUENCE; Schema: public; Owner: paul
--

CREATE SEQUENCE public.wishlists_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.wishlists_id_seq OWNER TO paul;

--
-- Name: wishlists_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paul
--

ALTER SEQUENCE public.wishlists_id_seq OWNED BY public.wishlists.id;


--
-- Name: addresses id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.addresses ALTER COLUMN id SET DEFAULT nextval('public.addresses_id_seq'::regclass);


--
-- Name: banners id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.banners ALTER COLUMN id SET DEFAULT nextval('public.banners_id_seq'::regclass);


--
-- Name: cart_items id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.cart_items ALTER COLUMN id SET DEFAULT nextval('public.cart_items_id_seq'::regclass);


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: customers id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.customers ALTER COLUMN id SET DEFAULT nextval('public.customers_id_seq'::regclass);


--
-- Name: order_items id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.order_items ALTER COLUMN id SET DEFAULT nextval('public.order_items_id_seq'::regclass);


--
-- Name: orders id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);


--
-- Name: product_photos id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.product_photos ALTER COLUMN id SET DEFAULT nextval('public.product_photos_id_seq'::regclass);


--
-- Name: product_variants id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.product_variants ALTER COLUMN id SET DEFAULT nextval('public.product_variants_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Name: recommendations id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.recommendations ALTER COLUMN id SET DEFAULT nextval('public.recommendations_id_seq'::regclass);


--
-- Name: reviews id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.reviews ALTER COLUMN id SET DEFAULT nextval('public.reviews_id_seq'::regclass);


--
-- Name: shopping_sessions id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.shopping_sessions ALTER COLUMN id SET DEFAULT nextval('public.shopping_sessions_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- Name: weeklies id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.weeklies ALTER COLUMN id SET DEFAULT nextval('public.weeklies_id_seq'::regclass);


--
-- Name: wishlists id; Type: DEFAULT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.wishlists ALTER COLUMN id SET DEFAULT nextval('public.wishlists_id_seq'::regclass);


--
-- Data for Name: addresses; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.addresses (id, customer_id, fullname, email, detail, is_default, created_at, updated_at, deleted_at) FROM stdin;
3	1	customer 1	customer1b	street 2	t	2024-11-23 05:50:04.423229+07	\N	\N
4	2	customer 2	customer2a	street 2a	t	2024-11-24 19:19:27.313344+07	\N	\N
\.


--
-- Data for Name: banners; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.banners (id, photo, title, subtitle, path_page, started_at, finished_at, created_at, updated_at, deleted_at) FROM stdin;
1	https://placehold.co/600x400?text=Banner+1	Banner Title 1	Banner Subtitle 1	https://excitable-circumference.info	2024-11-01 16:27:17.832431+07	2024-11-30 16:27:17.832431+07	2024-11-19 16:27:17.832431+07	\N	\N
2	https://placehold.co/600x400?text=Banner+2	Banner Title 2	Banner Subtitle 2	https://simple-information.org/	2024-10-19 16:28:02.252159+07	2024-11-23 16:27:17.832431+07	2024-11-19 16:28:02.252159+07	\N	\N
3	https://placehold.co/600x400?text=Banner+3	Banner Title 3	Banner Subtitle 3	https://sociable-casket.com/	2024-11-01 16:29:19.136979+07	2024-12-01 16:27:17.832431+07	2024-11-19 16:29:19.136979+07	\N	\N
4	https://placehold.co/600x400?text=Banner+4	Banner Title 4	Banner Subtitle 4	https://inferior-encouragement.com/	2024-11-05 16:30:12.888904+07	2024-11-21 16:27:17.832431+07	2024-11-19 16:30:12.888904+07	\N	\N
5	https://placehold.co/600x400?text=Banner+5	Banner Title 5	Banner Subtitle 5	https://flowery-abacus.biz/	2024-09-16 16:30:55.012185+07	2024-12-31 16:27:17.832431+07	2024-11-19 16:30:55.012185+07	\N	\N
\.


--
-- Data for Name: cart_items; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.cart_items (id, shopping_session_id, product_id, quantity, created_at, updated_at) FROM stdin;
42	2	3	2	2024-11-24 15:12:04.30275+07	\N
43	1	4	1	2024-11-24 15:46:42.600041+07	\N
\.


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.categories (id, name, created_at, updated_at, deleted_at) FROM stdin;
1	Category A	2024-11-19 15:25:57.76109+07	\N	\N
2	Category B	2024-11-19 15:25:57.76109+07	\N	\N
3	Category C	2024-11-19 15:25:57.76109+07	\N	\N
4	Category D	2024-11-19 15:25:57.76109+07	\N	\N
5	Category E	2024-11-19 15:25:57.76109+07	\N	\N
6	Category F	2024-11-19 15:25:57.76109+07	\N	\N
7	Category G	2024-11-19 15:25:57.76109+07	\N	\N
8	Category H	2024-11-19 15:25:57.76109+07	\N	\N
9	Category I	2024-11-19 15:25:57.76109+07	\N	\N
10	Category J	2024-11-19 15:25:57.76109+07	\N	\N
11	Category K	2024-11-19 15:25:57.76109+07	\N	\N
12	Category L	2024-11-19 15:25:57.76109+07	\N	\N
13	Category M	2024-11-19 15:25:57.76109+07	\N	\N
14	Category N	2024-11-19 15:25:57.76109+07	\N	\N
15	Category O	2024-11-19 15:25:57.76109+07	\N	\N
16	Category P	2024-11-19 15:25:57.76109+07	\N	\N
17	Category Q	2024-11-19 15:25:57.76109+07	\N	\N
18	Category R	2024-11-19 15:25:57.76109+07	\N	\N
19	Category S	2024-11-19 15:25:57.76109+07	\N	\N
20	Category T	2024-11-19 15:25:57.76109+07	\N	\N
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.customers (id, user_id, name) FROM stdin;
1	1	a
2	2	b
3	3	c
4	4	d
\.


--
-- Data for Name: order_items; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.order_items (id, order_id, product_id, quantity, unit_price, discount_rate, net_price) FROM stdin;
1	1	36	2	$30.27	6.22	$28.39
2	1	21	1	$41.96	5.11	$39.82
3	2	64	1	$27.74	7.13	$25.76
4	2	7	3	$17.16	6.26	$16.09
5	3	1	5	$40.50	9.22	$36.77
6	3	18	5	$13.72	4.99	$13.04
7	3	30	1	$27.18	5.14	$25.78
8	4	22	2	$47.22	4.05	$45.31
9	4	58	3	$33.48	0.26	$33.39
10	5	2	1	$12.01	4.27	$11.50
11	5	28	3	$30.18	0.65	$29.98
12	6	62	3	$20.48	0.66	$20.34
13	6	11	1	$23.05	1.82	$22.63
14	7	17	1	$31.84	1.12	$31.48
15	7	49	1	$11.00	1.02	$10.89
16	7	22	5	$25.66	5.86	$24.16
17	8	1	2	$48.41	1.03	$47.91
18	8	75	4	$10.46	3.12	$10.13
19	9	43	2	$12.83	9.61	$11.60
20	9	54	4	$19.83	1.64	$19.50
21	10	76	1	$35.82	2.45	$34.94
22	10	70	4	$20.31	2.33	$19.84
23	10	47	1	$21.94	4.12	$21.04
24	11	68	1	$44.21	4.37	$42.28
25	11	45	2	$32.48	0.54	$32.30
26	12	71	3	$38.18	8.38	$34.98
27	12	38	2	$39.01	1.73	$38.34
28	12	61	2	$10.69	2.94	$10.38
29	13	45	3	$33.44	7.81	$30.83
30	13	17	4	$14.39	9.89	$12.97
31	14	26	3	$21.61	0.49	$21.50
32	14	11	3	$32.93	4.36	$31.49
33	15	34	4	$40.46	7.13	$37.58
34	15	49	3	$39.37	9.3	$35.71
35	15	27	4	$30.26	6.5	$28.29
36	16	20	5	$35.99	9.63	$32.52
37	16	76	2	$40.05	6.51	$37.44
38	17	62	1	$40.86	0.0	$40.86
39	17	19	1	$35.96	7.88	$33.13
40	17	52	1	$16.13	8.08	$14.83
41	18	16	1	$31.17	4.24	$29.85
42	18	29	3	$21.23	9.47	$19.22
43	19	50	4	$30.45	6.39	$28.50
44	19	69	2	$33.15	9.34	$30.05
45	19	81	5	$38.54	5.85	$36.29
46	20	12	4	$27.42	4.35	$26.23
47	20	3	4	$35.45	3.56	$34.19
48	21	84	3	$18.14	4.87	$17.26
49	21	43	4	$20.66	0.59	$20.54
50	21	54	1	$10.42	7.09	$9.68
51	22	12	1	$13.72	1.2	$13.56
52	22	30	4	$30.33	4.3	$29.03
53	23	7	1	$19.65	0.96	$19.46
54	23	89	4	$37.30	4.35	$35.68
55	24	65	4	$29.11	3.77	$28.01
56	24	85	1	$19.22	7.03	$17.87
57	24	10	5	$34.63	8.35	$31.74
58	25	69	3	$26.98	2.17	$26.39
59	25	9	5	$25.25	5.26	$23.92
60	26	29	3	$14.95	6.5	$13.98
61	26	39	1	$37.95	3.17	$36.75
62	27	36	5	$32.32	9.55	$29.23
63	27	72	1	$37.54	8.0	$34.54
64	28	30	2	$25.87	1.78	$25.41
65	28	13	2	$48.44	6.19	$45.44
66	29	6	1	$27.37	4.91	$26.03
67	29	11	1	$44.11	9.35	$39.99
68	29	65	3	$30.60	0.9	$30.32
69	30	22	1	$48.65	6.79	$45.35
70	30	79	5	$14.19	0.63	$14.10
71	31	90	4	$36.25	7.41	$33.56
72	31	64	4	$35.79	1.8	$35.15
73	32	32	5	$41.49	9.67	$37.48
74	32	72	2	$43.97	7.25	$40.78
75	33	1	2	$22.28	5.42	$21.07
76	33	25	2	$33.31	6.92	$31.00
77	34	62	5	$10.16	6.08	$9.54
78	34	85	1	$18.90	4.43	$18.06
79	35	90	5	$33.38	1.0	$33.05
80	35	35	5	$35.05	5.63	$33.08
81	36	84	5	$49.35	9.05	$44.88
82	36	83	5	$10.44	7.33	$9.67
83	37	64	1	$33.51	4.04	$32.16
84	37	87	4	$25.72	9.39	$23.30
85	38	12	2	$34.31	2.48	$33.46
86	38	46	4	$15.69	2.69	$15.27
87	39	35	1	$36.58	2.94	$35.50
88	39	16	4	$31.17	4.77	$29.68
89	40	75	1	$32.81	4.85	$31.22
90	40	11	4	$41.36	2.61	$40.28
91	40	17	4	$22.04	5.06	$20.92
92	41	24	5	$37.18	1.14	$36.76
93	41	27	1	$20.62	8.61	$18.84
94	42	88	1	$40.47	3.37	$39.11
95	42	1	2	$44.46	4.6	$42.41
96	42	17	1	$17.55	7.64	$16.21
97	43	58	5	$13.30	1.32	$13.12
98	43	21	3	$43.74	2.68	$42.57
99	44	68	3	$17.97	4.45	$17.17
100	44	64	3	$41.88	7.86	$38.59
101	45	61	2	$47.23	4.93	$44.90
102	45	84	4	$36.09	6.9	$33.60
103	46	37	3	$19.86	5.47	$18.77
104	46	66	1	$18.41	9.46	$16.67
105	47	24	4	$17.46	1.3	$17.23
106	47	75	4	$21.16	4.27	$20.26
107	48	74	5	$22.19	5.65	$20.94
108	48	12	1	$16.59	9.45	$15.02
109	49	41	2	$49.64	6.6	$46.36
110	49	26	1	$18.20	2.99	$17.66
111	49	5	1	$39.30	7.92	$36.19
112	50	32	4	$10.24	0.95	$10.14
113	50	49	3	$20.41	4.37	$19.52
114	50	89	3	$48.97	9.19	$44.47
115	57	2	3	$39.02	0.48	$38.84
116	57	4	1	$36.55	6.85	$34.05
117	57	1	5	$12.30	10.00	$11.07
118	58	2	3	$39.02	0.48	$38.84
119	58	4	1	$36.55	6.85	$34.05
120	58	1	5	$12.30	10.00	$11.07
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.orders (id, customer_id, created_at, address_id, coupon_code, shipping, notes, payment_method) FROM stdin;
1	2	2024-09-22 14:36:24.262179+07	\N	\N	\N	\N	\N
2	3	2024-09-22 06:16:36.932724+07	\N	\N	\N	\N	\N
3	3	2024-10-27 10:58:31.59277+07	\N	\N	\N	\N	\N
4	4	2024-10-04 01:45:42.205598+07	\N	\N	\N	\N	\N
5	1	2024-10-04 23:14:07.793514+07	\N	\N	\N	\N	\N
6	2	2024-09-23 00:48:19.138694+07	\N	\N	\N	\N	\N
7	2	2024-11-16 19:42:20.751679+07	\N	\N	\N	\N	\N
8	4	2024-10-05 00:51:26.020249+07	\N	\N	\N	\N	\N
9	3	2024-10-22 01:39:29.134912+07	\N	\N	\N	\N	\N
10	2	2024-10-21 02:10:34.557492+07	\N	\N	\N	\N	\N
11	2	2024-10-24 16:17:32.566363+07	\N	\N	\N	\N	\N
12	1	2024-10-29 17:45:10.572222+07	\N	\N	\N	\N	\N
13	2	2024-11-16 06:32:08.615492+07	\N	\N	\N	\N	\N
14	1	2024-10-23 08:28:26.628135+07	\N	\N	\N	\N	\N
15	3	2024-10-31 19:24:15.312573+07	\N	\N	\N	\N	\N
16	1	2024-10-20 10:14:41.040212+07	\N	\N	\N	\N	\N
17	1	2024-11-05 17:09:01.639667+07	\N	\N	\N	\N	\N
18	1	2024-11-05 21:20:16.714129+07	\N	\N	\N	\N	\N
19	1	2024-09-27 12:35:18.790115+07	\N	\N	\N	\N	\N
20	2	2024-11-18 19:47:37.128781+07	\N	\N	\N	\N	\N
21	1	2024-11-06 17:41:41.774433+07	\N	\N	\N	\N	\N
22	1	2024-10-02 04:00:49.110527+07	\N	\N	\N	\N	\N
23	1	2024-10-25 13:19:00.192716+07	\N	\N	\N	\N	\N
24	2	2024-11-06 18:53:55.15969+07	\N	\N	\N	\N	\N
25	1	2024-09-27 22:54:03.014917+07	\N	\N	\N	\N	\N
26	2	2024-10-18 11:44:58.231413+07	\N	\N	\N	\N	\N
27	3	2024-10-20 19:46:51.260663+07	\N	\N	\N	\N	\N
28	2	2024-10-03 09:49:47.093068+07	\N	\N	\N	\N	\N
29	3	2024-10-07 02:16:00.270417+07	\N	\N	\N	\N	\N
30	4	2024-10-27 18:20:38.114955+07	\N	\N	\N	\N	\N
31	1	2024-11-14 12:32:44.221651+07	\N	\N	\N	\N	\N
32	1	2024-11-05 08:23:36.485636+07	\N	\N	\N	\N	\N
33	2	2024-11-10 11:56:14.63202+07	\N	\N	\N	\N	\N
34	2	2024-10-16 19:45:42.028129+07	\N	\N	\N	\N	\N
35	3	2024-10-09 01:50:45.692388+07	\N	\N	\N	\N	\N
36	4	2024-11-15 22:20:33.599225+07	\N	\N	\N	\N	\N
37	1	2024-09-21 17:32:08.485093+07	\N	\N	\N	\N	\N
38	2	2024-10-18 19:05:25.025397+07	\N	\N	\N	\N	\N
39	2	2024-11-15 23:13:10.584899+07	\N	\N	\N	\N	\N
40	4	2024-09-22 06:33:32.877298+07	\N	\N	\N	\N	\N
41	1	2024-11-13 09:40:48.594687+07	\N	\N	\N	\N	\N
42	2	2024-09-22 08:24:36.555532+07	\N	\N	\N	\N	\N
43	3	2024-11-08 07:25:23.980806+07	\N	\N	\N	\N	\N
44	1	2024-10-05 17:40:55.560276+07	\N	\N	\N	\N	\N
45	1	2024-10-15 12:35:19.712487+07	\N	\N	\N	\N	\N
46	4	2024-10-24 12:59:36.430457+07	\N	\N	\N	\N	\N
47	1	2024-10-29 09:07:56.231172+07	\N	\N	\N	\N	\N
48	2	2024-09-27 07:04:33.968429+07	\N	\N	\N	\N	\N
49	2	2024-11-12 19:54:48.926803+07	\N	\N	\N	\N	\N
50	4	2024-10-05 03:47:38.447427+07	\N	\N	\N	\N	\N
51	1	2024-11-24 14:18:52.123922+07	2	abcde	regular		bank
52	1	2024-11-24 14:21:45.410003+07	2	abcde	regular	lorem ipsum	bank
53	1	2024-11-24 15:14:24.510416+07	2	abcde	regular	lorem ipsum	bank
54	1	2024-11-24 15:18:47.704027+07	2	abcde	regular	lorem ipsum	bank
55	1	2024-11-24 15:19:56.892317+07	2	abcde	regular	lorem ipsum	bank
56	1	2024-11-24 15:23:27.842507+07	2	abcde	regular	lorem ipsum	bank
57	1	2024-11-24 15:25:39.249024+07	2	abcde	regular	lorem ipsum	bank
58	1	2024-11-24 15:34:34.586619+07	2	abcde	regular	lorem ipsum	bank
\.


--
-- Data for Name: product_photos; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.product_photos (id, product_id, photo_url) FROM stdin;
1	1	https://placehold.co/600x400/jpg?text=Product+1+A
2	1	https://placehold.co/600x400/jpg?text=Product+1+B
3	1	https://placehold.co/600x400/jpg?text=Product+1+C
4	1	https://placehold.co/600x400/jpg?text=Product+1+D
5	2	https://placehold.co/600x400/jpg?text=Product+2+A
6	2	https://placehold.co/600x400/jpg?text=Product+2+B
7	2	https://placehold.co/600x400/jpg?text=Product+2+C
8	2	https://placehold.co/600x400/jpg?text=Product+2+D
9	2	https://placehold.co/600x400/jpg?text=Product+2+E
10	3	https://placehold.co/600x400/jpg?text=Product+3+A
11	3	https://placehold.co/600x400/jpg?text=Product+3+B
12	3	https://placehold.co/600x400/jpg?text=Product+3+C
13	4	https://placehold.co/600x400/jpg?text=Product+4+A
14	4	https://placehold.co/600x400/jpg?text=Product+4+B
15	4	https://placehold.co/600x400/jpg?text=Product+4+C
16	4	https://placehold.co/600x400/jpg?text=Product+4+D
17	5	https://placehold.co/600x400/jpg?text=Product+5+A
18	5	https://placehold.co/600x400/jpg?text=Product+5+B
19	5	https://placehold.co/600x400/jpg?text=Product+5+C
20	5	https://placehold.co/600x400/jpg?text=Product+5+D
21	5	https://placehold.co/600x400/jpg?text=Product+5+E
22	6	https://placehold.co/600x400/jpg?text=Product+6+A
23	6	https://placehold.co/600x400/jpg?text=Product+6+B
24	6	https://placehold.co/600x400/jpg?text=Product+6+C
25	7	https://placehold.co/600x400/jpg?text=Product+7+A
26	7	https://placehold.co/600x400/jpg?text=Product+7+B
27	7	https://placehold.co/600x400/jpg?text=Product+7+C
28	7	https://placehold.co/600x400/jpg?text=Product+7+D
29	8	https://placehold.co/600x400/jpg?text=Product+8+A
30	8	https://placehold.co/600x400/jpg?text=Product+8+B
31	8	https://placehold.co/600x400/jpg?text=Product+8+C
32	8	https://placehold.co/600x400/jpg?text=Product+8+D
33	8	https://placehold.co/600x400/jpg?text=Product+8+E
34	9	https://placehold.co/600x400/jpg?text=Product+9+A
35	9	https://placehold.co/600x400/jpg?text=Product+9+B
36	9	https://placehold.co/600x400/jpg?text=Product+9+C
37	10	https://placehold.co/600x400/jpg?text=Product+10+A
38	10	https://placehold.co/600x400/jpg?text=Product+10+B
39	10	https://placehold.co/600x400/jpg?text=Product+10+C
40	10	https://placehold.co/600x400/jpg?text=Product+10+D
41	11	https://placehold.co/600x400/jpg?text=Product+11+A
42	11	https://placehold.co/600x400/jpg?text=Product+11+B
43	11	https://placehold.co/600x400/jpg?text=Product+11+C
44	11	https://placehold.co/600x400/jpg?text=Product+11+D
45	11	https://placehold.co/600x400/jpg?text=Product+11+E
46	12	https://placehold.co/600x400/jpg?text=Product+12+A
47	12	https://placehold.co/600x400/jpg?text=Product+12+B
48	12	https://placehold.co/600x400/jpg?text=Product+12+C
49	13	https://placehold.co/600x400/jpg?text=Product+13+A
50	13	https://placehold.co/600x400/jpg?text=Product+13+B
51	13	https://placehold.co/600x400/jpg?text=Product+13+C
52	13	https://placehold.co/600x400/jpg?text=Product+13+D
53	14	https://placehold.co/600x400/jpg?text=Product+14+A
54	14	https://placehold.co/600x400/jpg?text=Product+14+B
55	14	https://placehold.co/600x400/jpg?text=Product+14+C
56	14	https://placehold.co/600x400/jpg?text=Product+14+D
57	14	https://placehold.co/600x400/jpg?text=Product+14+E
58	15	https://placehold.co/600x400/jpg?text=Product+15+A
59	15	https://placehold.co/600x400/jpg?text=Product+15+B
60	15	https://placehold.co/600x400/jpg?text=Product+15+C
61	16	https://placehold.co/600x400/jpg?text=Product+16+A
62	16	https://placehold.co/600x400/jpg?text=Product+16+B
63	16	https://placehold.co/600x400/jpg?text=Product+16+C
64	16	https://placehold.co/600x400/jpg?text=Product+16+D
65	17	https://placehold.co/600x400/jpg?text=Product+17+A
66	17	https://placehold.co/600x400/jpg?text=Product+17+B
67	17	https://placehold.co/600x400/jpg?text=Product+17+C
68	17	https://placehold.co/600x400/jpg?text=Product+17+D
69	17	https://placehold.co/600x400/jpg?text=Product+17+E
70	18	https://placehold.co/600x400/jpg?text=Product+18+A
71	18	https://placehold.co/600x400/jpg?text=Product+18+B
72	18	https://placehold.co/600x400/jpg?text=Product+18+C
73	19	https://placehold.co/600x400/jpg?text=Product+19+A
74	19	https://placehold.co/600x400/jpg?text=Product+19+B
75	19	https://placehold.co/600x400/jpg?text=Product+19+C
76	19	https://placehold.co/600x400/jpg?text=Product+19+D
77	20	https://placehold.co/600x400/jpg?text=Product+20+A
78	20	https://placehold.co/600x400/jpg?text=Product+20+B
79	20	https://placehold.co/600x400/jpg?text=Product+20+C
80	20	https://placehold.co/600x400/jpg?text=Product+20+D
81	20	https://placehold.co/600x400/jpg?text=Product+20+E
82	21	https://placehold.co/600x400/jpg?text=Product+21+A
83	21	https://placehold.co/600x400/jpg?text=Product+21+B
84	21	https://placehold.co/600x400/jpg?text=Product+21+C
85	22	https://placehold.co/600x400/jpg?text=Product+22+A
86	22	https://placehold.co/600x400/jpg?text=Product+22+B
87	22	https://placehold.co/600x400/jpg?text=Product+22+C
88	22	https://placehold.co/600x400/jpg?text=Product+22+D
89	23	https://placehold.co/600x400/jpg?text=Product+23+A
90	23	https://placehold.co/600x400/jpg?text=Product+23+B
91	23	https://placehold.co/600x400/jpg?text=Product+23+C
92	23	https://placehold.co/600x400/jpg?text=Product+23+D
93	23	https://placehold.co/600x400/jpg?text=Product+23+E
94	24	https://placehold.co/600x400/jpg?text=Product+24+A
95	24	https://placehold.co/600x400/jpg?text=Product+24+B
96	24	https://placehold.co/600x400/jpg?text=Product+24+C
97	25	https://placehold.co/600x400/jpg?text=Product+25+A
98	25	https://placehold.co/600x400/jpg?text=Product+25+B
99	25	https://placehold.co/600x400/jpg?text=Product+25+C
100	25	https://placehold.co/600x400/jpg?text=Product+25+D
101	26	https://placehold.co/600x400/jpg?text=Product+26+A
102	26	https://placehold.co/600x400/jpg?text=Product+26+B
103	26	https://placehold.co/600x400/jpg?text=Product+26+C
104	26	https://placehold.co/600x400/jpg?text=Product+26+D
105	26	https://placehold.co/600x400/jpg?text=Product+26+E
106	27	https://placehold.co/600x400/jpg?text=Product+27+A
107	27	https://placehold.co/600x400/jpg?text=Product+27+B
108	27	https://placehold.co/600x400/jpg?text=Product+27+C
109	28	https://placehold.co/600x400/jpg?text=Product+28+A
110	28	https://placehold.co/600x400/jpg?text=Product+28+B
111	28	https://placehold.co/600x400/jpg?text=Product+28+C
112	28	https://placehold.co/600x400/jpg?text=Product+28+D
113	29	https://placehold.co/600x400/jpg?text=Product+29+A
114	29	https://placehold.co/600x400/jpg?text=Product+29+B
115	29	https://placehold.co/600x400/jpg?text=Product+29+C
116	29	https://placehold.co/600x400/jpg?text=Product+29+D
117	29	https://placehold.co/600x400/jpg?text=Product+29+E
118	30	https://placehold.co/600x400/jpg?text=Product+30+A
119	30	https://placehold.co/600x400/jpg?text=Product+30+B
120	30	https://placehold.co/600x400/jpg?text=Product+30+C
121	31	https://placehold.co/600x400/jpg?text=Product+31+A
122	31	https://placehold.co/600x400/jpg?text=Product+31+B
123	31	https://placehold.co/600x400/jpg?text=Product+31+C
124	31	https://placehold.co/600x400/jpg?text=Product+31+D
125	32	https://placehold.co/600x400/jpg?text=Product+32+A
126	32	https://placehold.co/600x400/jpg?text=Product+32+B
127	32	https://placehold.co/600x400/jpg?text=Product+32+C
128	32	https://placehold.co/600x400/jpg?text=Product+32+D
129	32	https://placehold.co/600x400/jpg?text=Product+32+E
130	33	https://placehold.co/600x400/jpg?text=Product+33+A
131	33	https://placehold.co/600x400/jpg?text=Product+33+B
132	33	https://placehold.co/600x400/jpg?text=Product+33+C
133	34	https://placehold.co/600x400/jpg?text=Product+34+A
134	34	https://placehold.co/600x400/jpg?text=Product+34+B
135	34	https://placehold.co/600x400/jpg?text=Product+34+C
136	34	https://placehold.co/600x400/jpg?text=Product+34+D
137	35	https://placehold.co/600x400/jpg?text=Product+35+A
138	35	https://placehold.co/600x400/jpg?text=Product+35+B
139	35	https://placehold.co/600x400/jpg?text=Product+35+C
140	35	https://placehold.co/600x400/jpg?text=Product+35+D
141	35	https://placehold.co/600x400/jpg?text=Product+35+E
142	36	https://placehold.co/600x400/jpg?text=Product+36+A
143	36	https://placehold.co/600x400/jpg?text=Product+36+B
144	36	https://placehold.co/600x400/jpg?text=Product+36+C
145	37	https://placehold.co/600x400/jpg?text=Product+37+A
146	37	https://placehold.co/600x400/jpg?text=Product+37+B
147	37	https://placehold.co/600x400/jpg?text=Product+37+C
148	37	https://placehold.co/600x400/jpg?text=Product+37+D
149	38	https://placehold.co/600x400/jpg?text=Product+38+A
150	38	https://placehold.co/600x400/jpg?text=Product+38+B
151	38	https://placehold.co/600x400/jpg?text=Product+38+C
152	38	https://placehold.co/600x400/jpg?text=Product+38+D
153	38	https://placehold.co/600x400/jpg?text=Product+38+E
154	39	https://placehold.co/600x400/jpg?text=Product+39+A
155	39	https://placehold.co/600x400/jpg?text=Product+39+B
156	39	https://placehold.co/600x400/jpg?text=Product+39+C
157	40	https://placehold.co/600x400/jpg?text=Product+40+A
158	40	https://placehold.co/600x400/jpg?text=Product+40+B
159	40	https://placehold.co/600x400/jpg?text=Product+40+C
160	40	https://placehold.co/600x400/jpg?text=Product+40+D
161	41	https://placehold.co/600x400/jpg?text=Product+41+A
162	41	https://placehold.co/600x400/jpg?text=Product+41+B
163	41	https://placehold.co/600x400/jpg?text=Product+41+C
164	41	https://placehold.co/600x400/jpg?text=Product+41+D
165	41	https://placehold.co/600x400/jpg?text=Product+41+E
166	42	https://placehold.co/600x400/jpg?text=Product+42+A
167	42	https://placehold.co/600x400/jpg?text=Product+42+B
168	42	https://placehold.co/600x400/jpg?text=Product+42+C
169	43	https://placehold.co/600x400/jpg?text=Product+43+A
170	43	https://placehold.co/600x400/jpg?text=Product+43+B
171	43	https://placehold.co/600x400/jpg?text=Product+43+C
172	43	https://placehold.co/600x400/jpg?text=Product+43+D
173	44	https://placehold.co/600x400/jpg?text=Product+44+A
174	44	https://placehold.co/600x400/jpg?text=Product+44+B
175	44	https://placehold.co/600x400/jpg?text=Product+44+C
176	44	https://placehold.co/600x400/jpg?text=Product+44+D
177	44	https://placehold.co/600x400/jpg?text=Product+44+E
178	45	https://placehold.co/600x400/jpg?text=Product+45+A
179	45	https://placehold.co/600x400/jpg?text=Product+45+B
180	45	https://placehold.co/600x400/jpg?text=Product+45+C
181	46	https://placehold.co/600x400/jpg?text=Product+46+A
182	46	https://placehold.co/600x400/jpg?text=Product+46+B
183	46	https://placehold.co/600x400/jpg?text=Product+46+C
184	46	https://placehold.co/600x400/jpg?text=Product+46+D
185	47	https://placehold.co/600x400/jpg?text=Product+47+A
186	47	https://placehold.co/600x400/jpg?text=Product+47+B
187	47	https://placehold.co/600x400/jpg?text=Product+47+C
188	47	https://placehold.co/600x400/jpg?text=Product+47+D
189	47	https://placehold.co/600x400/jpg?text=Product+47+E
190	48	https://placehold.co/600x400/jpg?text=Product+48+A
191	48	https://placehold.co/600x400/jpg?text=Product+48+B
192	48	https://placehold.co/600x400/jpg?text=Product+48+C
193	49	https://placehold.co/600x400/jpg?text=Product+49+A
194	49	https://placehold.co/600x400/jpg?text=Product+49+B
195	49	https://placehold.co/600x400/jpg?text=Product+49+C
196	49	https://placehold.co/600x400/jpg?text=Product+49+D
197	50	https://placehold.co/600x400/jpg?text=Product+50+A
198	50	https://placehold.co/600x400/jpg?text=Product+50+B
199	50	https://placehold.co/600x400/jpg?text=Product+50+C
200	50	https://placehold.co/600x400/jpg?text=Product+50+D
201	50	https://placehold.co/600x400/jpg?text=Product+50+E
202	51	https://placehold.co/600x400/jpg?text=Product+51+A
203	51	https://placehold.co/600x400/jpg?text=Product+51+B
204	51	https://placehold.co/600x400/jpg?text=Product+51+C
205	52	https://placehold.co/600x400/jpg?text=Product+52+A
206	52	https://placehold.co/600x400/jpg?text=Product+52+B
207	52	https://placehold.co/600x400/jpg?text=Product+52+C
208	52	https://placehold.co/600x400/jpg?text=Product+52+D
209	53	https://placehold.co/600x400/jpg?text=Product+53+A
210	53	https://placehold.co/600x400/jpg?text=Product+53+B
211	53	https://placehold.co/600x400/jpg?text=Product+53+C
212	53	https://placehold.co/600x400/jpg?text=Product+53+D
213	53	https://placehold.co/600x400/jpg?text=Product+53+E
214	54	https://placehold.co/600x400/jpg?text=Product+54+A
215	54	https://placehold.co/600x400/jpg?text=Product+54+B
216	54	https://placehold.co/600x400/jpg?text=Product+54+C
217	55	https://placehold.co/600x400/jpg?text=Product+55+A
218	55	https://placehold.co/600x400/jpg?text=Product+55+B
219	55	https://placehold.co/600x400/jpg?text=Product+55+C
220	55	https://placehold.co/600x400/jpg?text=Product+55+D
221	56	https://placehold.co/600x400/jpg?text=Product+56+A
222	56	https://placehold.co/600x400/jpg?text=Product+56+B
223	56	https://placehold.co/600x400/jpg?text=Product+56+C
224	56	https://placehold.co/600x400/jpg?text=Product+56+D
225	56	https://placehold.co/600x400/jpg?text=Product+56+E
226	57	https://placehold.co/600x400/jpg?text=Product+57+A
227	57	https://placehold.co/600x400/jpg?text=Product+57+B
228	57	https://placehold.co/600x400/jpg?text=Product+57+C
229	58	https://placehold.co/600x400/jpg?text=Product+58+A
230	58	https://placehold.co/600x400/jpg?text=Product+58+B
231	58	https://placehold.co/600x400/jpg?text=Product+58+C
232	58	https://placehold.co/600x400/jpg?text=Product+58+D
233	59	https://placehold.co/600x400/jpg?text=Product+59+A
234	59	https://placehold.co/600x400/jpg?text=Product+59+B
235	59	https://placehold.co/600x400/jpg?text=Product+59+C
236	59	https://placehold.co/600x400/jpg?text=Product+59+D
237	59	https://placehold.co/600x400/jpg?text=Product+59+E
238	60	https://placehold.co/600x400/jpg?text=Product+60+A
239	60	https://placehold.co/600x400/jpg?text=Product+60+B
240	60	https://placehold.co/600x400/jpg?text=Product+60+C
241	61	https://placehold.co/600x400/jpg?text=Product+61+A
242	61	https://placehold.co/600x400/jpg?text=Product+61+B
243	61	https://placehold.co/600x400/jpg?text=Product+61+C
244	61	https://placehold.co/600x400/jpg?text=Product+61+D
245	62	https://placehold.co/600x400/jpg?text=Product+62+A
246	62	https://placehold.co/600x400/jpg?text=Product+62+B
247	62	https://placehold.co/600x400/jpg?text=Product+62+C
248	62	https://placehold.co/600x400/jpg?text=Product+62+D
249	62	https://placehold.co/600x400/jpg?text=Product+62+E
250	63	https://placehold.co/600x400/jpg?text=Product+63+A
251	63	https://placehold.co/600x400/jpg?text=Product+63+B
252	63	https://placehold.co/600x400/jpg?text=Product+63+C
253	64	https://placehold.co/600x400/jpg?text=Product+64+A
254	64	https://placehold.co/600x400/jpg?text=Product+64+B
255	64	https://placehold.co/600x400/jpg?text=Product+64+C
256	64	https://placehold.co/600x400/jpg?text=Product+64+D
257	65	https://placehold.co/600x400/jpg?text=Product+65+A
258	65	https://placehold.co/600x400/jpg?text=Product+65+B
259	65	https://placehold.co/600x400/jpg?text=Product+65+C
260	65	https://placehold.co/600x400/jpg?text=Product+65+D
261	65	https://placehold.co/600x400/jpg?text=Product+65+E
262	66	https://placehold.co/600x400/jpg?text=Product+66+A
263	66	https://placehold.co/600x400/jpg?text=Product+66+B
264	66	https://placehold.co/600x400/jpg?text=Product+66+C
265	67	https://placehold.co/600x400/jpg?text=Product+67+A
266	67	https://placehold.co/600x400/jpg?text=Product+67+B
267	67	https://placehold.co/600x400/jpg?text=Product+67+C
268	67	https://placehold.co/600x400/jpg?text=Product+67+D
269	68	https://placehold.co/600x400/jpg?text=Product+68+A
270	68	https://placehold.co/600x400/jpg?text=Product+68+B
271	68	https://placehold.co/600x400/jpg?text=Product+68+C
272	68	https://placehold.co/600x400/jpg?text=Product+68+D
273	68	https://placehold.co/600x400/jpg?text=Product+68+E
274	69	https://placehold.co/600x400/jpg?text=Product+69+A
275	69	https://placehold.co/600x400/jpg?text=Product+69+B
276	69	https://placehold.co/600x400/jpg?text=Product+69+C
277	70	https://placehold.co/600x400/jpg?text=Product+70+A
278	70	https://placehold.co/600x400/jpg?text=Product+70+B
279	70	https://placehold.co/600x400/jpg?text=Product+70+C
280	70	https://placehold.co/600x400/jpg?text=Product+70+D
281	71	https://placehold.co/600x400/jpg?text=Product+71+A
282	71	https://placehold.co/600x400/jpg?text=Product+71+B
283	71	https://placehold.co/600x400/jpg?text=Product+71+C
284	71	https://placehold.co/600x400/jpg?text=Product+71+D
285	71	https://placehold.co/600x400/jpg?text=Product+71+E
286	72	https://placehold.co/600x400/jpg?text=Product+72+A
287	72	https://placehold.co/600x400/jpg?text=Product+72+B
288	72	https://placehold.co/600x400/jpg?text=Product+72+C
289	73	https://placehold.co/600x400/jpg?text=Product+73+A
290	73	https://placehold.co/600x400/jpg?text=Product+73+B
291	73	https://placehold.co/600x400/jpg?text=Product+73+C
292	73	https://placehold.co/600x400/jpg?text=Product+73+D
293	74	https://placehold.co/600x400/jpg?text=Product+74+A
294	74	https://placehold.co/600x400/jpg?text=Product+74+B
295	74	https://placehold.co/600x400/jpg?text=Product+74+C
296	74	https://placehold.co/600x400/jpg?text=Product+74+D
297	74	https://placehold.co/600x400/jpg?text=Product+74+E
298	75	https://placehold.co/600x400/jpg?text=Product+75+A
299	75	https://placehold.co/600x400/jpg?text=Product+75+B
300	75	https://placehold.co/600x400/jpg?text=Product+75+C
301	76	https://placehold.co/600x400/jpg?text=Product+76+A
302	76	https://placehold.co/600x400/jpg?text=Product+76+B
303	76	https://placehold.co/600x400/jpg?text=Product+76+C
304	76	https://placehold.co/600x400/jpg?text=Product+76+D
305	77	https://placehold.co/600x400/jpg?text=Product+77+A
306	77	https://placehold.co/600x400/jpg?text=Product+77+B
307	77	https://placehold.co/600x400/jpg?text=Product+77+C
308	77	https://placehold.co/600x400/jpg?text=Product+77+D
309	77	https://placehold.co/600x400/jpg?text=Product+77+E
310	78	https://placehold.co/600x400/jpg?text=Product+78+A
311	78	https://placehold.co/600x400/jpg?text=Product+78+B
312	78	https://placehold.co/600x400/jpg?text=Product+78+C
313	79	https://placehold.co/600x400/jpg?text=Product+79+A
314	79	https://placehold.co/600x400/jpg?text=Product+79+B
315	79	https://placehold.co/600x400/jpg?text=Product+79+C
316	79	https://placehold.co/600x400/jpg?text=Product+79+D
317	80	https://placehold.co/600x400/jpg?text=Product+80+A
318	80	https://placehold.co/600x400/jpg?text=Product+80+B
319	80	https://placehold.co/600x400/jpg?text=Product+80+C
320	80	https://placehold.co/600x400/jpg?text=Product+80+D
321	80	https://placehold.co/600x400/jpg?text=Product+80+E
322	81	https://placehold.co/600x400/jpg?text=Product+81+A
323	81	https://placehold.co/600x400/jpg?text=Product+81+B
324	81	https://placehold.co/600x400/jpg?text=Product+81+C
325	82	https://placehold.co/600x400/jpg?text=Product+82+A
326	82	https://placehold.co/600x400/jpg?text=Product+82+B
327	82	https://placehold.co/600x400/jpg?text=Product+82+C
328	82	https://placehold.co/600x400/jpg?text=Product+82+D
329	83	https://placehold.co/600x400/jpg?text=Product+83+A
330	83	https://placehold.co/600x400/jpg?text=Product+83+B
331	83	https://placehold.co/600x400/jpg?text=Product+83+C
332	83	https://placehold.co/600x400/jpg?text=Product+83+D
333	83	https://placehold.co/600x400/jpg?text=Product+83+E
334	84	https://placehold.co/600x400/jpg?text=Product+84+A
335	84	https://placehold.co/600x400/jpg?text=Product+84+B
336	84	https://placehold.co/600x400/jpg?text=Product+84+C
337	85	https://placehold.co/600x400/jpg?text=Product+85+A
338	85	https://placehold.co/600x400/jpg?text=Product+85+B
339	85	https://placehold.co/600x400/jpg?text=Product+85+C
340	85	https://placehold.co/600x400/jpg?text=Product+85+D
341	86	https://placehold.co/600x400/jpg?text=Product+86+A
342	86	https://placehold.co/600x400/jpg?text=Product+86+B
343	86	https://placehold.co/600x400/jpg?text=Product+86+C
344	86	https://placehold.co/600x400/jpg?text=Product+86+D
345	86	https://placehold.co/600x400/jpg?text=Product+86+E
346	87	https://placehold.co/600x400/jpg?text=Product+87+A
347	87	https://placehold.co/600x400/jpg?text=Product+87+B
348	87	https://placehold.co/600x400/jpg?text=Product+87+C
349	88	https://placehold.co/600x400/jpg?text=Product+88+A
350	88	https://placehold.co/600x400/jpg?text=Product+88+B
351	88	https://placehold.co/600x400/jpg?text=Product+88+C
352	88	https://placehold.co/600x400/jpg?text=Product+88+D
353	89	https://placehold.co/600x400/jpg?text=Product+89+A
354	89	https://placehold.co/600x400/jpg?text=Product+89+B
355	89	https://placehold.co/600x400/jpg?text=Product+89+C
356	89	https://placehold.co/600x400/jpg?text=Product+89+D
357	89	https://placehold.co/600x400/jpg?text=Product+89+E
358	90	https://placehold.co/600x400/jpg?text=Product+90+A
359	90	https://placehold.co/600x400/jpg?text=Product+90+B
360	90	https://placehold.co/600x400/jpg?text=Product+90+C
\.


--
-- Data for Name: product_variants; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.product_variants (id, product_id, name) FROM stdin;
1	1	Variant 1A
2	1	Variant 1B
3	1	Variant 1C
4	2	Variant 2A
5	2	Variant 2B
6	2	Variant 2C
7	2	Variant 2D
8	3	Variant 3A
9	3	Variant 3B
10	3	Variant 3C
11	3	Variant 3D
12	3	Variant 3E
13	4	Variant 4A
14	4	Variant 4B
15	5	Variant 5A
16	5	Variant 5B
17	5	Variant 5C
18	6	Variant 6A
19	6	Variant 6B
20	6	Variant 6C
21	6	Variant 6D
22	7	Variant 7A
23	7	Variant 7B
24	7	Variant 7C
25	7	Variant 7D
26	7	Variant 7E
27	8	Variant 8A
28	8	Variant 8B
29	9	Variant 9A
30	9	Variant 9B
31	9	Variant 9C
32	10	Variant 10A
33	10	Variant 10B
34	10	Variant 10C
35	10	Variant 10D
36	11	Variant 11A
37	11	Variant 11B
38	11	Variant 11C
39	11	Variant 11D
40	11	Variant 11E
41	12	Variant 12A
42	12	Variant 12B
43	13	Variant 13A
44	13	Variant 13B
45	13	Variant 13C
46	14	Variant 14A
47	14	Variant 14B
48	14	Variant 14C
49	14	Variant 14D
50	15	Variant 15A
51	15	Variant 15B
52	15	Variant 15C
53	15	Variant 15D
54	15	Variant 15E
55	16	Variant 16A
56	16	Variant 16B
57	17	Variant 17A
58	17	Variant 17B
59	17	Variant 17C
60	18	Variant 18A
61	18	Variant 18B
62	18	Variant 18C
63	18	Variant 18D
64	19	Variant 19A
65	19	Variant 19B
66	19	Variant 19C
67	19	Variant 19D
68	19	Variant 19E
69	20	Variant 20A
70	20	Variant 20B
71	21	Variant 21A
72	21	Variant 21B
73	21	Variant 21C
74	22	Variant 22A
75	22	Variant 22B
76	22	Variant 22C
77	22	Variant 22D
78	23	Variant 23A
79	23	Variant 23B
80	23	Variant 23C
81	23	Variant 23D
82	23	Variant 23E
83	24	Variant 24A
84	24	Variant 24B
85	25	Variant 25A
86	25	Variant 25B
87	25	Variant 25C
88	26	Variant 26A
89	26	Variant 26B
90	26	Variant 26C
91	26	Variant 26D
92	27	Variant 27A
93	27	Variant 27B
94	27	Variant 27C
95	27	Variant 27D
96	27	Variant 27E
97	28	Variant 28A
98	28	Variant 28B
99	29	Variant 29A
100	29	Variant 29B
101	29	Variant 29C
102	30	Variant 30A
103	30	Variant 30B
104	30	Variant 30C
105	30	Variant 30D
106	31	Variant 31A
107	31	Variant 31B
108	31	Variant 31C
109	31	Variant 31D
110	31	Variant 31E
111	32	Variant 32A
112	32	Variant 32B
113	33	Variant 33A
114	33	Variant 33B
115	33	Variant 33C
116	34	Variant 34A
117	34	Variant 34B
118	34	Variant 34C
119	34	Variant 34D
120	35	Variant 35A
121	35	Variant 35B
122	35	Variant 35C
123	35	Variant 35D
124	35	Variant 35E
125	36	Variant 36A
126	36	Variant 36B
127	37	Variant 37A
128	37	Variant 37B
129	37	Variant 37C
130	38	Variant 38A
131	38	Variant 38B
132	38	Variant 38C
133	38	Variant 38D
134	39	Variant 39A
135	39	Variant 39B
136	39	Variant 39C
137	39	Variant 39D
138	39	Variant 39E
139	40	Variant 40A
140	40	Variant 40B
141	41	Variant 41A
142	41	Variant 41B
143	41	Variant 41C
144	42	Variant 42A
145	42	Variant 42B
146	42	Variant 42C
147	42	Variant 42D
148	43	Variant 43A
149	43	Variant 43B
150	43	Variant 43C
151	43	Variant 43D
152	43	Variant 43E
153	44	Variant 44A
154	44	Variant 44B
155	45	Variant 45A
156	45	Variant 45B
157	45	Variant 45C
158	46	Variant 46A
159	46	Variant 46B
160	46	Variant 46C
161	46	Variant 46D
162	47	Variant 47A
163	47	Variant 47B
164	47	Variant 47C
165	47	Variant 47D
166	47	Variant 47E
167	48	Variant 48A
168	48	Variant 48B
169	49	Variant 49A
170	49	Variant 49B
171	49	Variant 49C
172	50	Variant 50A
173	50	Variant 50B
174	50	Variant 50C
175	50	Variant 50D
176	51	Variant 51A
177	51	Variant 51B
178	51	Variant 51C
179	51	Variant 51D
180	51	Variant 51E
181	52	Variant 52A
182	52	Variant 52B
183	53	Variant 53A
184	53	Variant 53B
185	53	Variant 53C
186	54	Variant 54A
187	54	Variant 54B
188	54	Variant 54C
189	54	Variant 54D
190	55	Variant 55A
191	55	Variant 55B
192	55	Variant 55C
193	55	Variant 55D
194	55	Variant 55E
195	56	Variant 56A
196	56	Variant 56B
197	57	Variant 57A
198	57	Variant 57B
199	57	Variant 57C
200	58	Variant 58A
201	58	Variant 58B
202	58	Variant 58C
203	58	Variant 58D
204	59	Variant 59A
205	59	Variant 59B
206	59	Variant 59C
207	59	Variant 59D
208	59	Variant 59E
209	60	Variant 60A
210	60	Variant 60B
211	61	Variant 61A
212	61	Variant 61B
213	61	Variant 61C
214	62	Variant 62A
215	62	Variant 62B
216	62	Variant 62C
217	62	Variant 62D
218	63	Variant 63A
219	63	Variant 63B
220	63	Variant 63C
221	63	Variant 63D
222	63	Variant 63E
223	64	Variant 64A
224	64	Variant 64B
225	65	Variant 65A
226	65	Variant 65B
227	65	Variant 65C
228	66	Variant 66A
229	66	Variant 66B
230	66	Variant 66C
231	66	Variant 66D
232	67	Variant 67A
233	67	Variant 67B
234	67	Variant 67C
235	67	Variant 67D
236	67	Variant 67E
237	68	Variant 68A
238	68	Variant 68B
239	69	Variant 69A
240	69	Variant 69B
241	69	Variant 69C
242	70	Variant 70A
243	70	Variant 70B
244	70	Variant 70C
245	70	Variant 70D
246	71	Variant 71A
247	71	Variant 71B
248	71	Variant 71C
249	71	Variant 71D
250	71	Variant 71E
251	72	Variant 72A
252	72	Variant 72B
253	73	Variant 73A
254	73	Variant 73B
255	73	Variant 73C
256	74	Variant 74A
257	74	Variant 74B
258	74	Variant 74C
259	74	Variant 74D
260	75	Variant 75A
261	75	Variant 75B
262	75	Variant 75C
263	75	Variant 75D
264	75	Variant 75E
265	76	Variant 76A
266	76	Variant 76B
267	77	Variant 77A
268	77	Variant 77B
269	77	Variant 77C
270	78	Variant 78A
271	78	Variant 78B
272	78	Variant 78C
273	78	Variant 78D
274	79	Variant 79A
275	79	Variant 79B
276	79	Variant 79C
277	79	Variant 79D
278	79	Variant 79E
279	80	Variant 80A
280	80	Variant 80B
281	81	Variant 81A
282	81	Variant 81B
283	81	Variant 81C
284	82	Variant 82A
285	82	Variant 82B
286	82	Variant 82C
287	82	Variant 82D
288	83	Variant 83A
289	83	Variant 83B
290	83	Variant 83C
291	83	Variant 83D
292	83	Variant 83E
293	84	Variant 84A
294	84	Variant 84B
295	85	Variant 85A
296	85	Variant 85B
297	85	Variant 85C
298	86	Variant 86A
299	86	Variant 86B
300	86	Variant 86C
301	86	Variant 86D
302	87	Variant 87A
303	87	Variant 87B
304	87	Variant 87C
305	87	Variant 87D
306	87	Variant 87E
307	88	Variant 88A
308	88	Variant 88B
309	89	Variant 89A
310	89	Variant 89B
311	89	Variant 89C
312	90	Variant 90A
313	90	Variant 90B
314	90	Variant 90C
315	90	Variant 90D
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.products (id, name, thumbnail, price, discount_rate, category_id, created_at) FROM stdin;
2	Product 2	https://placehold.co/600x400?text=Product+2	$39.02	0.48	8	2024-11-08 06:33:03.882359+07
3	Product 3	https://placehold.co/600x400?text=Product+3	$34.04	3.75	3	2024-10-19 05:50:12.934505+07
4	Product 4	https://placehold.co/600x400?text=Product+4	$36.55	6.85	7	2024-10-19 17:08:23.864501+07
5	Product 5	https://placehold.co/600x400?text=Product+5	$23.90	4.98	10	2024-11-16 13:21:06.889226+07
6	Product 6	https://placehold.co/600x400?text=Product+6	$21.77	1.34	15	2024-11-07 00:27:59.318074+07
7	Product 7	https://placehold.co/600x400?text=Product+7	$3.46	5.63	12	2024-11-01 09:13:57.724775+07
8	Product 8	https://placehold.co/600x400?text=Product+8	$49.04	7.69	2	2024-10-17 02:49:01.175354+07
9	Product 9	https://placehold.co/600x400?text=Product+9	$30.83	8.05	18	2024-11-11 22:45:31.299136+07
10	Product 10	https://placehold.co/600x400?text=Product+10	$44.99	1.45	6	2024-11-15 17:51:57.565745+07
11	Product 11	https://placehold.co/600x400?text=Product+11	$37.76	3.73	4	2024-11-16 07:43:22.300449+07
12	Product 12	https://placehold.co/600x400?text=Product+12	$22.79	8.72	9	2024-10-24 20:42:32.869289+07
13	Product 13	https://placehold.co/600x400?text=Product+13	$12.44	7.56	13	2024-11-11 06:34:03.462995+07
14	Product 14	https://placehold.co/600x400?text=Product+14	$46.12	4.25	17	2024-10-20 06:49:21.386183+07
15	Product 15	https://placehold.co/600x400?text=Product+15	$6.07	8.73	11	2024-10-31 08:07:51.439055+07
16	Product 16	https://placehold.co/600x400?text=Product+16	$1.03	2.12	5	2024-11-21 14:14:37.31765+07
17	Product 17	https://placehold.co/600x400?text=Product+17	$34.17	3.76	8	2024-10-30 01:15:39.936231+07
18	Product 18	https://placehold.co/600x400?text=Product+18	$34.37	2.28	3	2024-11-07 14:31:15.777626+07
19	Product 19	https://placehold.co/600x400?text=Product+19	$25.15	5.45	7	2024-10-17 05:12:59.92021+07
20	Product 20	https://placehold.co/600x400?text=Product+20	$41.79	2.20	10	2024-11-16 00:35:52.003414+07
21	Product 21	https://placehold.co/600x400?text=Product+21	$1.37	5.64	15	2024-11-03 14:56:13.925907+07
22	Product 22	https://placehold.co/600x400?text=Product+22	$31.76	0.22	12	2024-10-10 06:17:51.400683+07
23	Product 23	https://placehold.co/600x400?text=Product+23	$22.47	4.87	2	2024-11-12 02:47:35.614526+07
24	Product 24	https://placehold.co/600x400?text=Product+24	$47.68	8.16	18	2024-11-05 06:12:53.124913+07
25	Product 25	https://placehold.co/600x400?text=Product+25	$40.06	1.81	6	2024-10-22 14:47:19.818114+07
26	Product 26	https://placehold.co/600x400?text=Product+26	$35.73	0.49	4	2024-11-07 07:36:24.996611+07
27	Product 27	https://placehold.co/600x400?text=Product+27	$24.03	7.99	9	2024-11-20 07:14:02.899859+07
28	Product 28	https://placehold.co/600x400?text=Product+28	$20.60	9.46	13	2024-11-06 10:57:17.003934+07
29	Product 29	https://placehold.co/600x400?text=Product+29	$3.17	2.00	17	2024-10-21 12:21:47.242512+07
30	Product 30	https://placehold.co/600x400?text=Product+30	$30.86	3.57	11	2024-10-31 16:22:25.514569+07
31	Product 31	https://placehold.co/600x400?text=Product+31	$36.10	0.26	5	2024-10-28 12:20:08.260136+07
32	Product 32	https://placehold.co/600x400?text=Product+32	$28.55	7.13	8	2024-11-17 11:22:04.164832+07
33	Product 33	https://placehold.co/600x400?text=Product+33	$33.33	0.38	3	2024-10-18 10:02:51.460079+07
34	Product 34	https://placehold.co/600x400?text=Product+34	$11.52	5.61	7	2024-10-27 09:36:33.917578+07
35	Product 35	https://placehold.co/600x400?text=Product+35	$47.45	7.87	10	2024-10-13 06:50:54.944044+07
36	Product 36	https://placehold.co/600x400?text=Product+36	$13.72	1.86	15	2024-11-21 00:57:35.113536+07
37	Product 37	https://placehold.co/600x400?text=Product+37	$28.31	7.27	12	2024-10-25 20:01:23.922391+07
38	Product 38	https://placehold.co/600x400?text=Product+38	$24.12	5.97	2	2024-11-02 15:50:02.64661+07
39	Product 39	https://placehold.co/600x400?text=Product+39	$28.91	1.25	18	2024-11-09 19:36:58.202922+07
40	Product 40	https://placehold.co/600x400?text=Product+40	$30.66	1.56	6	2024-10-31 13:54:04.554226+07
41	Product 41	https://placehold.co/600x400?text=Product+41	$32.99	2.11	4	2024-10-15 06:52:02.636511+07
42	Product 42	https://placehold.co/600x400?text=Product+42	$26.26	6.28	9	2024-11-05 00:11:54.219941+07
43	Product 43	https://placehold.co/600x400?text=Product+43	$15.56	7.42	13	2024-11-19 02:59:26.771171+07
44	Product 44	https://placehold.co/600x400?text=Product+44	$34.26	4.06	17	2024-11-04 18:01:32.417444+07
45	Product 45	https://placehold.co/600x400?text=Product+45	$35.09	9.11	11	2024-10-30 19:45:23.675491+07
46	Product 46	https://placehold.co/600x400?text=Product+46	$14.52	9.74	5	2024-11-01 20:35:25.86401+07
47	Product 47	https://placehold.co/600x400?text=Product+47	$12.82	0.36	8	2024-11-23 16:32:48.80475+07
48	Product 48	https://placehold.co/600x400?text=Product+48	$4.17	8.18	3	2024-10-12 07:42:57.030081+07
49	Product 49	https://placehold.co/600x400?text=Product+49	$42.63	8.51	7	2024-11-12 15:21:53.478335+07
50	Product 50	https://placehold.co/600x400?text=Product+50	$47.50	5.03	10	2024-10-26 01:52:03.239925+07
51	Product 51	https://placehold.co/600x400?text=Product+51	$7.86	6.98	15	2024-11-16 04:15:10.641749+07
52	Product 52	https://placehold.co/600x400?text=Product+52	$1.96	0.74	12	2024-11-01 16:04:37.574431+07
1	Product 1	https://placehold.co/600x400?text=Product+1	$12.30	10.00	5	2024-11-15 11:16:26.09646+07
53	Product 53	https://placehold.co/600x400?text=Product+53	$29.24	3.92	2	2024-11-17 02:15:43.454226+07
54	Product 54	https://placehold.co/600x400?text=Product+54	$13.06	4.83	18	2024-10-27 02:17:01.511209+07
55	Product 55	https://placehold.co/600x400?text=Product+55	$12.41	8.02	6	2024-11-04 19:42:10.250881+07
56	Product 56	https://placehold.co/600x400?text=Product+56	$21.37	7.27	4	2024-10-09 23:27:19.798286+07
57	Product 57	https://placehold.co/600x400?text=Product+57	$40.60	3.75	9	2024-11-07 13:40:25.950107+07
58	Product 58	https://placehold.co/600x400?text=Product+58	$20.91	4.43	13	2024-10-28 13:23:37.43704+07
59	Product 59	https://placehold.co/600x400?text=Product+59	$16.46	9.51	17	2024-11-10 13:47:19.689348+07
60	Product 60	https://placehold.co/600x400?text=Product+60	$15.11	0.55	11	2024-10-16 17:37:33.929585+07
61	Product 61	https://placehold.co/600x400?text=Product+61	$31.41	8.69	14	2024-10-13 07:09:16.746483+07
62	Product 62	https://placehold.co/600x400?text=Product+62	$17.89	6.73	16	2024-10-19 15:17:24.915688+07
63	Product 63	https://placehold.co/600x400?text=Product+63	$48.19	4.66	19	2024-10-22 02:25:05.581424+07
64	Product 64	https://placehold.co/600x400?text=Product+64	$8.49	7.13	20	2024-11-22 14:03:04.496336+07
65	Product 65	https://placehold.co/600x400?text=Product+65	$26.42	3.70	1	2024-10-17 08:13:21.866575+07
66	Product 66	https://placehold.co/600x400?text=Product+66	$2.73	2.45	14	2024-11-16 06:18:40.358849+07
67	Product 67	https://placehold.co/600x400?text=Product+67	$42.44	0.34	19	2024-10-19 06:33:51.354149+07
68	Product 68	https://placehold.co/600x400?text=Product+68	$25.01	3.83	16	2024-10-16 18:40:10.15807+07
69	Product 69	https://placehold.co/600x400?text=Product+69	$6.82	6.54	14	2024-11-21 18:08:39.593086+07
70	Product 70	https://placehold.co/600x400?text=Product+70	$46.80	4.64	20	2024-11-06 18:26:04.919659+07
71	Product 71	https://placehold.co/600x400?text=Product+71	$47.08	2.14	1	2024-10-16 10:57:02.888926+07
72	Product 72	https://placehold.co/600x400?text=Product+72	$4.93	1.19	19	2024-11-16 21:41:10.409462+07
73	Product 73	https://placehold.co/600x400?text=Product+73	$38.52	6.43	16	2024-11-02 11:16:21.5123+07
74	Product 74	https://placehold.co/600x400?text=Product+74	$11.68	3.52	14	2024-10-23 02:15:56.055791+07
75	Product 75	https://placehold.co/600x400?text=Product+75	$15.38	4.25	1	2024-11-21 17:34:01.000999+07
76	Product 76	https://placehold.co/600x400?text=Product+76	$44.64	0.58	20	2024-11-18 18:36:11.779516+07
77	Product 77	https://placehold.co/600x400?text=Product+77	$9.75	0.35	19	2024-10-20 00:55:57.443372+07
78	Product 78	https://placehold.co/600x400?text=Product+78	$3.48	2.46	16	2024-11-02 21:26:26.768745+07
79	Product 79	https://placehold.co/600x400?text=Product+79	$8.98	0.43	14	2024-11-18 20:40:16.794342+07
80	Product 80	https://placehold.co/600x400?text=Product+80	$6.03	2.95	1	2024-11-16 09:46:33.787058+07
81	Product 81	https://placehold.co/600x400?text=Product+81	$17.01	1.92	20	2024-10-21 02:55:50.678865+07
82	Product 82	https://placehold.co/600x400?text=Product+82	$48.12	0.88	19	2024-11-15 12:43:06.487137+07
83	Product 83	https://placehold.co/600x400?text=Product+83	$5.52	0.01	14	2024-11-15 14:21:37.360189+07
84	Product 84	https://placehold.co/600x400?text=Product+84	$26.18	6.37	16	2024-11-03 09:03:37.45838+07
85	Product 85	https://placehold.co/600x400?text=Product+85	$35.65	3.05	1	2024-10-26 13:50:10.152037+07
86	Product 86	https://placehold.co/600x400?text=Product+86	$27.84	9.77	20	2024-10-27 15:26:39.128562+07
87	Product 87	https://placehold.co/600x400?text=Product+87	$32.56	9.53	19	2024-10-11 13:26:49.066947+07
88	Product 88	https://placehold.co/600x400?text=Product+88	$47.19	9.58	14	2024-10-19 18:50:29.495674+07
89	Product 89	https://placehold.co/600x400?text=Product+89	$38.00	6.65	16	2024-10-17 19:31:44.894186+07
90	Product 90	https://placehold.co/600x400?text=Product+90	$33.74	9.54	1	2024-11-16 11:57:12.021684+07
\.


--
-- Data for Name: recommendations; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.recommendations (id, title, subtitle, photo, product_id, created_at) FROM stdin;
1	Title Recommendation 1	Subtitle Recommendation 1	https://placehold.co/600x400?text=Recommendation+1	1	2024-11-21 01:27:52.464452+07
\.


--
-- Data for Name: reviews; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.reviews (id, order_item_id, rating, created_at) FROM stdin;
\.


--
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.sessions (token, user_id, created_at, expired_at) FROM stdin;
0d0078aa-7b5e-429b-ab56-e53f0ff3060e	1	2024-11-23 09:08:35.859445+07	2024-11-23 09:10:35.859445+07
e94a6165-5eac-4bef-bebb-03b7676a2893	1	2024-11-23 09:22:07.761547+07	2024-11-23 09:24:07.761547+07
0d2a0814-f6f7-40be-b379-6f075acde2f7	1	2024-11-23 09:35:36.540172+07	2024-11-23 09:36:36.540172+07
62a5e9af-4fc4-4109-b99c-83cb50726fd0	1	2024-11-23 09:44:54.542623+07	6107-12-17 11:51:54.542623+07
1868dc61-154f-4588-9772-c282e4384ae4	1	2024-11-23 09:50:27.886266+07	2024-11-23 09:51:27.886266+07
f30ac5f9-7d18-4278-b773-d2471f81c3cc	1	2024-11-23 09:53:07.872221+07	6107-12-17 12:00:07.872221+07
55900225-10ca-4282-a2ca-a81638ae286d	1	2024-11-24 07:55:06.316124+07	2024-11-24 07:56:06.316124+07
83a81910-87f6-454d-a4d0-efc4d6d08dd0	1	2024-11-24 07:55:39.034521+07	6107-12-18 10:02:39.034521+07
\.


--
-- Data for Name: shopping_sessions; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.shopping_sessions (id, customer_id) FROM stdin;
1	1
2	2
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.users (id, username, password, created_at, updated_at, deleted_at) FROM stdin;
1	a	a	2024-11-19 03:02:56.4429+07	\N	\N
2	b	b	2024-11-19 03:03:56.208187+07	\N	\N
3	c	c	2024-11-19 03:04:32.709533+07	\N	\N
4	d	d	2024-11-19 03:06:57.647632+07	\N	\N
5	e	e	2024-11-19 23:49:47.777936+07	\N	\N
\.


--
-- Data for Name: weeklies; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.weeklies (id, photo, title, subtitle, product_id, started_at, finished_at, created_at) FROM stdin;
1	\N	\N	\N	1	2024-11-18 00:00:00+07	2024-11-24 00:00:00+07	2024-11-20 23:44:51.989212+07
2	\N	\N	\N	2	2024-11-11 00:00:00+07	2024-11-17 00:00:00+07	2024-11-20 23:44:51.989212+07
3	\N	\N	\N	3	2024-11-25 00:00:00+07	2024-12-01 00:00:00+07	2024-11-20 23:44:51.989212+07
\.


--
-- Data for Name: wishlists; Type: TABLE DATA; Schema: public; Owner: paul
--

COPY public.wishlists (id, customer_id, product_id, created_at) FROM stdin;
1	1	1	2024-11-23 14:53:50.404052+07
3	1	4	2024-11-24 08:23:43.828323+07
\.


--
-- Name: addresses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.addresses_id_seq', 4, true);


--
-- Name: banners_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.banners_id_seq', 5, true);


--
-- Name: cart_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.cart_items_id_seq', 43, true);


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.categories_id_seq', 20, true);


--
-- Name: customers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.customers_id_seq', 4, true);


--
-- Name: order_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.order_items_id_seq', 120, true);


--
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.orders_id_seq', 58, true);


--
-- Name: product_photos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.product_photos_id_seq', 360, true);


--
-- Name: product_variants_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.product_variants_id_seq', 315, true);


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.products_id_seq', 90, true);


--
-- Name: recommendations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.recommendations_id_seq', 1, true);


--
-- Name: reviews_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.reviews_id_seq', 1, false);


--
-- Name: shopping_sessions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.shopping_sessions_id_seq', 11, true);


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.user_id_seq', 5, true);


--
-- Name: weeklies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.weeklies_id_seq', 3, true);


--
-- Name: wishlists_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paul
--

SELECT pg_catalog.setval('public.wishlists_id_seq', 3, true);


--
-- Name: addresses addresses_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.addresses
    ADD CONSTRAINT addresses_pkey PRIMARY KEY (id);


--
-- Name: banners banners_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.banners
    ADD CONSTRAINT banners_pkey PRIMARY KEY (id);


--
-- Name: cart_items cart_items_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.cart_items
    ADD CONSTRAINT cart_items_pkey PRIMARY KEY (id);


--
-- Name: cart_items cart_items_shopping_session_id_product_id_key; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.cart_items
    ADD CONSTRAINT cart_items_shopping_session_id_product_id_key UNIQUE (shopping_session_id, product_id);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: customers customers_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);


--
-- Name: order_items order_items_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_pkey PRIMARY KEY (id);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: product_photos product_photos_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.product_photos
    ADD CONSTRAINT product_photos_pkey PRIMARY KEY (id);


--
-- Name: product_variants product_variants_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.product_variants
    ADD CONSTRAINT product_variants_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: recommendations recommendations_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.recommendations
    ADD CONSTRAINT recommendations_pkey PRIMARY KEY (id);


--
-- Name: reviews reviews_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_pkey PRIMARY KEY (id);


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (token);


--
-- Name: shopping_sessions shopping_sessions_customer_id_key; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.shopping_sessions
    ADD CONSTRAINT shopping_sessions_customer_id_key UNIQUE (customer_id);


--
-- Name: shopping_sessions shopping_sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.shopping_sessions
    ADD CONSTRAINT shopping_sessions_pkey PRIMARY KEY (id);


--
-- Name: users user_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: weeklies weeklies_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.weeklies
    ADD CONSTRAINT weeklies_pkey PRIMARY KEY (id);


--
-- Name: wishlists wishlists_pkey; Type: CONSTRAINT; Schema: public; Owner: paul
--

ALTER TABLE ONLY public.wishlists
    ADD CONSTRAINT wishlists_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

