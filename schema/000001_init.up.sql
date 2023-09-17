CREATE TABLE public.mc_users (
    id SERIAL PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    password varchar(255) NOT NULL
);

CREATE TABLE public.user_permissions (
    id SERIAL PRIMARY KEY,
    method varchar(255) NOT NULL,
    qty integer,
    date_to datetime NOT NULL,
    usert_id integer REFERENCES public.mc_users(id),
    CONSTRAINT unique_user_method UNIQUE (usert_id, method)
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
    created_at timestamp with time zone DEFAULT current_timestamp,
    CONSTRAINT fk_card_stock FOREIGN KEY (sku) REFERENCES public.card (sku)
);