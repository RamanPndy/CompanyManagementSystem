CREATE TABLE IF NOT EXISTS public.company (
    id character varying NOT NULL,
    name character varying NOT NULL,
    description character varying,
    employees bigint,
    registered boolean,
    type character varying,
    "createdat" timestamp without time zone NOT NULL,
    "updatedat" timestamp without time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS public.users (
    id SERIAL PRIMARY KEY,
    username character varying NOT NULL,
    password character varying NOT NULL,
    description character varying,
    isactive boolean,
    "createdat" timestamp without time zone NOT NULL,
    "updatedat" timestamp without time zone NOT NULL
);