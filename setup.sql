CREATE TABLE IF NOT EXISTS public.company (
    id character varying(32) NOT NULL,
    name character varying NOT NULL,
    description character varying NOT NULL,
    employees bigint NOT NULL,
    registered boolean NOT NULL,
    type character varying NOT NULL,
    "createdAt" timestamp without time zone NOT NULL,
    "updatedAt" timestamp without time zone NOT NULL
);