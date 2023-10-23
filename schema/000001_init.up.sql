CREATE TABLE public.mc_users (
    id SERIAL PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    end_date date,
    password varchar(255) NOT NULL
);
CREATE TABLE public.card (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL DEFAULT '',
    image varchar(255) NOT NULL DEFAULT '',
    user_id integer REFERENCES public.mc_users (id) ON DELETE CASCADE,
    active boolean DEFAULT true,
    sku integer,
    CONSTRAINT unique_user_sku UNIQUE (user_id, sku)
);
CREATE TABLE public.stock (
    id SERIAL PRIMARY KEY,
    sku integer,
    wh integer,
    qty integer,
    created_at timestamp with time zone DEFAULT current_timestamp
);
CREATE TABLE public.commission (
    id SERIAL PRIMARY KEY,
    category varchar(255),
    subject varchar(255) NOT NULL DEFAULT '',
    commission integer NOT NULL DEFAULT 0,
    fbs integer NOT NULL DEFAULT 0,
    fbo integer NOT NULL DEFAULT 0
);
-- CREATE USER IF NOT EXISTS mc_service WITH ENCRYPTED PASSWORD
-- GRANT CONNECT ON DATABASE * TO *;
-- GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO *;
-- GRANT ALL PRIVILEGES ON TABLE public.mc_users TO *;
-- GRANT ALL PRIVILEGES ON TABLE public.card TO *;
-- GRANT ALL PRIVILEGES ON TABLE public.stock TO *;