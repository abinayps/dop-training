Q1. Implement a Product Management module that supports the following operations:

Create a product

List all products

Get product by ID

Update product by ID

Delete product by ID


* Create a Product with the following data:

Name (string), Category (string), Price (float64), Stock (int)

🔹 APIs to implement
Method	Endpoint	Purpose
POST	/products	Create product
GET	/products	List all products
GET	/products/:id	Get product by ID
PUT	/products/:id	Update product by ID
DELETE	/products/:id	Delete product


CREATE SEQUENCE IF NOT EXISTS public.product_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


-- Table: public.product_details

-- DROP TABLE IF EXISTS public.product_details;

CREATE TABLE IF NOT EXISTS public.product_details
(
    id integer NOT NULL DEFAULT nextval('product_details_id_seq'::regclass),
    name character varying(150) COLLATE pg_catalog."default" NOT NULL,
    category character varying(100) COLLATE pg_catalog."default" NOT NULL,
    price numeric(10,2) NOT NULL,
    stock integer NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT product_details_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.product_details
    OWNER to postgres;


