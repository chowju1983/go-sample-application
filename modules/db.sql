DO $$
BEGIN

IF NOT EXISTS(
        SELECT schema_name
          FROM information_schema.schemata
          WHERE schema_name = 'FraudRiskSchema'
      )

THEN

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
-- SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;


SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
-- SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 3360 (class 0 OID 0)
-- Dependencies: 3359
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- TOC entry 6 (class 2615 OID 24576)
-- Name: FraudRiskSchema; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA "FraudRiskSchema";


ALTER SCHEMA "FraudRiskSchema" OWNER TO postgres;

SET default_tablespace = '';


--
-- TOC entry 216 (class 1259 OID 24603)
-- Name: Account; Type: TABLE; Schema: FraudRiskSchema; Owner: postgres
--

CREATE TABLE "FraudRiskSchema"."Account" (
    account_id integer NOT NULL,
    account_no integer,
    card_type_id integer,
    card_no integer,
    pin integer,
    aav integer,
    cust_id integer
);


ALTER TABLE "FraudRiskSchema"."Account" OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 24602)
-- Name: Account_account_id_seq; Type: SEQUENCE; Schema: FraudRiskSchema; Owner: postgres
--

CREATE SEQUENCE "FraudRiskSchema"."Account_account_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE "FraudRiskSchema"."Account_account_id_seq" OWNER TO postgres;

--
-- TOC entry 3361 (class 0 OID 0)
-- Dependencies: 215
-- Name: Account_account_id_seq; Type: SEQUENCE OWNED BY; Schema: FraudRiskSchema; Owner: postgres
--

ALTER SEQUENCE "FraudRiskSchema"."Account_account_id_seq" OWNED BY "FraudRiskSchema"."Account".account_id;


--
-- TOC entry 214 (class 1259 OID 24587)
-- Name: CustomerProfile; Type: TABLE; Schema: FraudRiskSchema; Owner: postgres
--

CREATE TABLE "FraudRiskSchema"."CustomerProfile" (
    cust_id integer NOT NULL,
    name text,
    address text,
    age smallint,
    phone integer,
    email text
);


ALTER TABLE "FraudRiskSchema"."CustomerProfile" OWNER TO postgres;

--
-- TOC entry 213 (class 1259 OID 24586)
-- Name: CustomerProfile_cust_id_seq; Type: SEQUENCE; Schema: FraudRiskSchema; Owner: postgres
--

CREATE SEQUENCE "FraudRiskSchema"."CustomerProfile_cust_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE "FraudRiskSchema"."CustomerProfile_cust_id_seq" OWNER TO postgres;

--
-- TOC entry 3362 (class 0 OID 0)
-- Dependencies: 213
-- Name: CustomerProfile_cust_id_seq; Type: SEQUENCE OWNED BY; Schema: FraudRiskSchema; Owner: postgres
--

ALTER SEQUENCE "FraudRiskSchema"."CustomerProfile_cust_id_seq" OWNED BY "FraudRiskSchema"."CustomerProfile".cust_id;


--
-- TOC entry 212 (class 1259 OID 24578)
-- Name: ProductData; Type: TABLE; Schema: FraudRiskSchema; Owner: postgres
--

CREATE TABLE "FraudRiskSchema"."ProductData" (
    card_type_id integer NOT NULL,
    card_type text,
    limits integer,
    currency text
);


ALTER TABLE "FraudRiskSchema"."ProductData" OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 24577)
-- Name: ProductData_card_type_id_seq; Type: SEQUENCE; Schema: FraudRiskSchema; Owner: postgres
--

CREATE SEQUENCE "FraudRiskSchema"."ProductData_card_type_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE "FraudRiskSchema"."ProductData_card_type_id_seq" OWNER TO postgres;

--
-- TOC entry 3363 (class 0 OID 0)
-- Dependencies: 211
-- Name: ProductData_card_type_id_seq; Type: SEQUENCE OWNED BY; Schema: FraudRiskSchema; Owner: postgres
--

ALTER SEQUENCE "FraudRiskSchema"."ProductData_card_type_id_seq" OWNED BY "FraudRiskSchema"."ProductData".card_type_id;


--
-- TOC entry 220 (class 1259 OID 24634)
-- Name: Request; Type: TABLE; Schema: FraudRiskSchema; Owner: postgres
--

CREATE TABLE "FraudRiskSchema"."Request" (
    request_id integer NOT NULL,
    browserinfo text,
    guid text,
    origin_of_transaction text,
    transaction_type text,
    session_id integer,
    channel text,
    ipaddress text,
    device_id text,
    device_model text,
    device_version text,
    journey_type text,
    agent_id text
);


ALTER TABLE "FraudRiskSchema"."Request" OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 24633)
-- Name: Request_request_id_seq; Type: SEQUENCE; Schema: FraudRiskSchema; Owner: postgres
--

CREATE SEQUENCE "FraudRiskSchema"."Request_request_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE "FraudRiskSchema"."Request_request_id_seq" OWNER TO postgres;

--
-- TOC entry 3364 (class 0 OID 0)
-- Dependencies: 219
-- Name: Request_request_id_seq; Type: SEQUENCE OWNED BY; Schema: FraudRiskSchema; Owner: postgres
--

ALTER SEQUENCE "FraudRiskSchema"."Request_request_id_seq" OWNED BY "FraudRiskSchema"."Request".request_id;


--
-- TOC entry 218 (class 1259 OID 24620)
-- Name: Session; Type: TABLE; Schema: FraudRiskSchema; Owner: postgres
--

CREATE TABLE "FraudRiskSchema"."Session" (
    session_id integer NOT NULL,
    account_id integer,
    token text
);


ALTER TABLE "FraudRiskSchema"."Session" OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 24619)
-- Name: Session_session_id_seq; Type: SEQUENCE; Schema: FraudRiskSchema; Owner: postgres
--

CREATE SEQUENCE "FraudRiskSchema"."Session_session_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE "FraudRiskSchema"."Session_session_id_seq" OWNER TO postgres;

--
-- TOC entry 3365 (class 0 OID 0)
-- Dependencies: 217
-- Name: Session_session_id_seq; Type: SEQUENCE OWNED BY; Schema: FraudRiskSchema; Owner: postgres
--

ALTER SEQUENCE "FraudRiskSchema"."Session_session_id_seq" OWNED BY "FraudRiskSchema"."Session".session_id;


--
-- TOC entry 3188 (class 2604 OID 24606)
-- Name: Account account_id; Type: DEFAULT; Schema: FraudRiskSchema; Owner: postgres
--

ALTER TABLE ONLY "FraudRiskSchema"."Account" ALTER COLUMN account_id SET DEFAULT nextval('"FraudRiskSchema"."Account_account_id_seq"'::regclass);


--
-- TOC entry 3187 (class 2604 OID 24590)
-- Name: CustomerProfile cust_id; Type: DEFAULT; Schema: FraudRiskSchema; Owner: postgres
--

ALTER TABLE ONLY "FraudRiskSchema"."CustomerProfile" ALTER COLUMN cust_id SET DEFAULT nextval('"FraudRiskSchema"."CustomerProfile_cust_id_seq"'::regclass);


--
-- TOC entry 3186 (class 2604 OID 24581)
-- Name: ProductData card_type_id; Type: DEFAULT; Schema: FraudRiskSchema; Owner: postgres
--

ALTER TABLE ONLY "FraudRiskSchema"."ProductData" ALTER COLUMN card_type_id SET DEFAULT nextval('"FraudRiskSchema"."ProductData_card_type_id_seq"'::regclass);


--
-- TOC entry 3190 (class 2604 OID 24637)
-- Name: Request request_id; Type: DEFAULT; Schema: FraudRiskSchema; Owner: postgres
--

ALTER TABLE ONLY "FraudRiskSchema"."Request" ALTER COLUMN request_id SET DEFAULT nextval('"FraudRiskSchema"."Request_request_id_seq"'::regclass);


--
-- TOC entry 3189 (class 2604 OID 24623)
-- Name: Session session_id; Type: DEFAULT; Schema: FraudRiskSchema; Owner: postgres
--

ALTER TABLE ONLY "FraudRiskSchema"."Session" ALTER COLUMN session_id SET DEFAULT nextval('"FraudRiskSchema"."Session_session_id_seq"'::regclass);


--
-- TOC entry 3349 (class 0 OID 24603)
-- Dependencies: 216
-- Data for Name: Account; Type: TABLE DATA; Schema: FraudRiskSchema; Owner: postgres
--

INSERT INTO "FraudRiskSchema"."Account" VALUES (1, 334455, 1, 654778999, 7456, 234, 1);
INSERT INTO "FraudRiskSchema"."Account" VALUES (2, 334455, 2, 554778999, 7764, 245, 1);
INSERT INTO "FraudRiskSchema"."Account" VALUES (3, 334456, 1, 454778999, 3456, 789, 2);
INSERT INTO "FraudRiskSchema"."Account" VALUES (4, 334456, 2, 354778999, 3678, 378, 2);
INSERT INTO "FraudRiskSchema"."Account" VALUES (5, 334457, 1, 254778999, 7856, 390, 3);
INSERT INTO "FraudRiskSchema"."Account" VALUES (6, 334457, 2, 154778999, 2345, 789, 3);


--
-- TOC entry 3347 (class 0 OID 24587)
-- Dependencies: 214
-- Data for Name: CustomerProfile; Type: TABLE DATA; Schema: FraudRiskSchema; Owner: postgres
--

INSERT INTO "FraudRiskSchema"."CustomerProfile" VALUES (1, 'Parth', 'Mumbai', 44, 876543223, 'parth@gmail.com');
INSERT INTO "FraudRiskSchema"."CustomerProfile" VALUES (2, 'Kunal', 'Pune', 23, 786543556, 'kunal@yahoo.com');
INSERT INTO "FraudRiskSchema"."CustomerProfile" VALUES (3, 'Raj', 'Noida', 44, 765439887, 'raj@gmail.com');


--
-- TOC entry 3345 (class 0 OID 24578)
-- Dependencies: 212
-- Data for Name: ProductData; Type: TABLE DATA; Schema: FraudRiskSchema; Owner: postgres
--

INSERT INTO "FraudRiskSchema"."ProductData" VALUES (1, 'AMEX debit card', 200000, 'INR');
INSERT INTO "FraudRiskSchema"."ProductData" VALUES (2, 'AMEX credit card', 300000, 'INR');


--
-- TOC entry 3353 (class 0 OID 24634)
-- Dependencies: 220
-- Data for Name: Request; Type: TABLE DATA; Schema: FraudRiskSchema; Owner: postgres
--

INSERT INTO "FraudRiskSchema"."Request" VALUES (1, 'Chrome', 'AAA2345', 'USA', 'FUNDTRANS', 1, 'mobile', '172.2.4.66', 'MB23456', 'SamsungGal', '2', '', '3');


--
-- TOC entry 3351 (class 0 OID 24620)
-- Dependencies: 218
-- Data for Name: Session; Type: TABLE DATA; Schema: FraudRiskSchema; Owner: postgres
--

INSERT INTO "FraudRiskSchema"."Session" VALUES (1, 1, 'JJ567896433');
INSERT INTO "FraudRiskSchema"."Session" VALUES (2, 2, 'JJ567896434');
INSERT INTO "FraudRiskSchema"."Session" VALUES (3, 3, 'JJ567896435');
INSERT INTO "FraudRiskSchema"."Session" VALUES (4, 4, 'JJ567896436');
INSERT INTO "FraudRiskSchema"."Session" VALUES (5, 5, 'JJ567896437');
INSERT INTO "FraudRiskSchema"."Session" VALUES (6, 6, 'JJ567896438');


--
-- TOC entry 3366 (class 0 OID 0)
-- Dependencies: 215
-- Name: Account_account_id_seq; Type: SEQUENCE SET; Schema: FraudRiskSchema; Owner: postgres
--

PERFORM pg_catalog.setval('"FraudRiskSchema"."Account_account_id_seq"', 1, true);


--
-- TOC entry 3367 (class 0 OID 0)
-- Dependencies: 213
-- Name: CustomerProfile_cust_id_seq; Type: SEQUENCE SET; Schema: FraudRiskSchema; Owner: postgres
--

PERFORM pg_catalog.setval('"FraudRiskSchema"."CustomerProfile_cust_id_seq"', 1, true);


--
-- TOC entry 3368 (class 0 OID 0)
-- Dependencies: 211
-- Name: ProductData_card_type_id_seq; Type: SEQUENCE SET; Schema: FraudRiskSchema; Owner: postgres
--

PERFORM pg_catalog.setval('"FraudRiskSchema"."ProductData_card_type_id_seq"', 1, true);


--
-- TOC entry 3369 (class 0 OID 0)
-- Dependencies: 219
-- Name: Request_request_id_seq; Type: SEQUENCE SET; Schema: FraudRiskSchema; Owner: postgres
--

PERFORM pg_catalog.setval('"FraudRiskSchema"."Request_request_id_seq"', 1, true);


--
-- TOC entry 3370 (class 0 OID 0)
-- Dependencies: 217
-- Name: Session_session_id_seq; Type: SEQUENCE SET; Schema: FraudRiskSchema; Owner: postgres
--

PERFORM pg_catalog.setval('"FraudRiskSchema"."Session_session_id_seq"', 1, true);


--
-- TOC entry 3196 (class 2606 OID 24608)
-- Name: Account Account_pkey; Type: CONSTRAINT; Schema: FraudRiskSchema; Owner: postgres
--

ALTER TABLE ONLY "FraudRiskSchema"."Account"
    ADD CONSTRAINT "Account_pkey" PRIMARY KEY (account_id);


--
-- TOC entry 3194 (class 2606 OID 24594)
-- Name: CustomerProfile CustomerProfile_pkey; Type: CONSTRAINT; Schema: FraudRiskSchema; Owner: postgres
--

ALTER TABLE ONLY "FraudRiskSchema"."CustomerProfile"
    ADD CONSTRAINT "CustomerProfile_pkey" PRIMARY KEY (cust_id);


--
-- TOC entry 3192 (class 2606 OID 24585)
-- Name: ProductData ProductData_pkey; Type: CONSTRAINT; Schema: FraudRiskSchema; Owner: postgres
--

ALTER TABLE ONLY "FraudRiskSchema"."ProductData"
    ADD CONSTRAINT "ProductData_pkey" PRIMARY KEY (card_type_id);


--
-- TOC entry 3200 (class 2606 OID 24641)
-- Name: Request Request_pkey; Type: CONSTRAINT; Schema: FraudRiskSchema; Owner: postgres
--

ALTER TABLE ONLY "FraudRiskSchema"."Request"
    ADD CONSTRAINT "Request_pkey" PRIMARY KEY (request_id);


--
-- TOC entry 3198 (class 2606 OID 24627)
-- Name: Session Session_pkey; Type: CONSTRAINT; Schema: FraudRiskSchema; Owner: postgres
--

ALTER TABLE ONLY "FraudRiskSchema"."Session"
    ADD CONSTRAINT "Session_pkey" PRIMARY KEY (session_id);


--
-- TOC entry 3203 (class 2606 OID 24628)
-- Name: Session accountfk; Type: FK CONSTRAINT; Schema: FraudRiskSchema; Owner: postgres
--

ALTER TABLE ONLY "FraudRiskSchema"."Session"
    ADD CONSTRAINT accountfk FOREIGN KEY (account_id) REFERENCES "FraudRiskSchema"."Account"(account_id);


--
-- TOC entry 3202 (class 2606 OID 24614)
-- Name: Account cardtypefk; Type: FK CONSTRAINT; Schema: FraudRiskSchema; Owner: postgres
--

ALTER TABLE ONLY "FraudRiskSchema"."Account"
    ADD CONSTRAINT cardtypefk FOREIGN KEY (card_type_id) REFERENCES "FraudRiskSchema"."ProductData"(card_type_id) NOT VALID;


--
-- TOC entry 3201 (class 2606 OID 24609)
-- Name: Account custidfk; Type: FK CONSTRAINT; Schema: FraudRiskSchema; Owner: postgres
--

ALTER TABLE ONLY "FraudRiskSchema"."Account"
    ADD CONSTRAINT custidfk FOREIGN KEY (cust_id) REFERENCES "FraudRiskSchema"."CustomerProfile"(cust_id);


--
-- TOC entry 3204 (class 2606 OID 24642)
-- Name: Request sessionfk; Type: FK CONSTRAINT; Schema: FraudRiskSchema; Owner: postgres
--

ALTER TABLE ONLY "FraudRiskSchema"."Request"
    ADD CONSTRAINT sessionfk FOREIGN KEY (session_id) REFERENCES "FraudRiskSchema"."Session"(session_id);


-- Completed on 2022-01-21 19:24:36

--
-- PostgreSQL database dump complete
--

END IF;

END
$$;
