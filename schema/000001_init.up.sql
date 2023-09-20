CREATE TABLE public.mc_users (
    id SERIAL PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    password varchar(255) NOT NULL
);
CREATE TABLE public.users_subscriptions (
    id SERIAL PRIMARY KEY,
    user_id integer REFERENCES public.mc_users (id) ON DELETE CASCADE,
    price real,
    end_date date,
    quantity int
);
CREATE TABLE public.card (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL DEFAULT '',
    user_id integer REFERENCES public.mc_users (id) ON DELETE CASCADE,
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
-- GRANT CONNECT ON DATABASE * TO *;
-- GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO *;
-- GRANT ALL PRIVILEGES ON TABLE public.mc_users TO *;
-- GRANT ALL PRIVILEGES ON TABLE public.users_subscriptions TO *;
-- GRANT ALL PRIVILEGES ON TABLE public.card TO *;
-- GRANT ALL PRIVILEGES ON TABLE public.stock TO *;