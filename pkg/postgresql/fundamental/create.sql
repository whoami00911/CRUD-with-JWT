-- Table: public.AbuseEntity

-- DROP TABLE IF EXISTS public."AbuseEntity";

CREATE TABLE IF NOT EXISTS public."AbuseEntity"
(
    "ipAddress" character varying(15) COLLATE pg_catalog."default",
    "isPublic" boolean,
    "ipVersion" integer,
    "isWhitelisted" boolean,
    "abuseConfidenceScore" integer,
    "countryCode" character varying(5) COLLATE pg_catalog."default",
    "countryName" character varying(100) COLLATE pg_catalog."default",
    "usageType" character varying(100) COLLATE pg_catalog."default",
    isp character varying(200) COLLATE pg_catalog."default",
    "isTor" boolean,
    "isFromDB" boolean,
    CONSTRAINT "unique_ipAddress" UNIQUE ("ipAddress")
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."AbuseEntity"
    OWNER to postgres;