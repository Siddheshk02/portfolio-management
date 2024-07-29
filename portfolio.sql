PGDMP         -                |            portfolio-management    15.3    15.3                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    16397    portfolio-management    DATABASE     �   CREATE DATABASE "portfolio-management" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_India.1252';
 &   DROP DATABASE "portfolio-management";
                postgres    false            �            1259    16419    assets    TABLE     �   CREATE TABLE public.assets (
    id integer NOT NULL,
    portfolio_id integer NOT NULL,
    name text NOT NULL,
    value numeric NOT NULL,
    created_at timestamp with time zone NOT NULL
);
    DROP TABLE public.assets;
       public         heap    postgres    false            �            1259    16418    assets_id_seq    SEQUENCE     �   CREATE SEQUENCE public.assets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.assets_id_seq;
       public          postgres    false    219                       0    0    assets_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.assets_id_seq OWNED BY public.assets.id;
          public          postgres    false    218            �            1259    16410 
   portfolios    TABLE     �   CREATE TABLE public.portfolios (
    id integer NOT NULL,
    user_id integer NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
    DROP TABLE public.portfolios;
       public         heap    postgres    false            �            1259    16409    portfolios_id_seq    SEQUENCE     �   CREATE SEQUENCE public.portfolios_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.portfolios_id_seq;
       public          postgres    false    217                       0    0    portfolios_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.portfolios_id_seq OWNED BY public.portfolios.id;
          public          postgres    false    216            �            1259    16399    users    TABLE     �   CREATE TABLE public.users (
    id integer NOT NULL,
    username text NOT NULL,
    password text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
    DROP TABLE public.users;
       public         heap    postgres    false            �            1259    16398    users_id_seq    SEQUENCE     �   CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          postgres    false    215                       0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          postgres    false    214            q           2604    16422 	   assets id    DEFAULT     f   ALTER TABLE ONLY public.assets ALTER COLUMN id SET DEFAULT nextval('public.assets_id_seq'::regclass);
 8   ALTER TABLE public.assets ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    218    219    219            p           2604    16413    portfolios id    DEFAULT     n   ALTER TABLE ONLY public.portfolios ALTER COLUMN id SET DEFAULT nextval('public.portfolios_id_seq'::regclass);
 <   ALTER TABLE public.portfolios ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    217    216    217            o           2604    16402    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    215    214    215                      0    16419    assets 
   TABLE DATA           K   COPY public.assets (id, portfolio_id, name, value, created_at) FROM stdin;
    public          postgres    false    219   9                 0    16410 
   portfolios 
   TABLE DATA           O   COPY public.portfolios (id, user_id, name, created_at, updated_at) FROM stdin;
    public          postgres    false    217   �       	          0    16399    users 
   TABLE DATA           O   COPY public.users (id, username, password, created_at, updated_at) FROM stdin;
    public          postgres    false    215   �                  0    0    assets_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.assets_id_seq', 3, true);
          public          postgres    false    218                       0    0    portfolios_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.portfolios_id_seq', 1, true);
          public          postgres    false    216                       0    0    users_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.users_id_seq', 1, true);
          public          postgres    false    214            y           2606    16426    assets assets_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.assets
    ADD CONSTRAINT assets_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.assets DROP CONSTRAINT assets_pkey;
       public            postgres    false    219            w           2606    16417    portfolios portfolios_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public.portfolios
    ADD CONSTRAINT portfolios_pkey PRIMARY KEY (id);
 D   ALTER TABLE ONLY public.portfolios DROP CONSTRAINT portfolios_pkey;
       public            postgres    false    217            s           2606    16406    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    215            u           2606    16408    users users_username_key 
   CONSTRAINT     W   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);
 B   ALTER TABLE ONLY public.users DROP CONSTRAINT users_username_key;
       public            postgres    false    215               J   x�3�4�.�O�.�440�4202�50�5�T04�25�26�326152�60r��	�71�31165�������� ?>p         5   x�3�4��4202�50�5�T04�21�2��362562�60�26 ��+F��� ��      	   s   x�3��L�T1JT14PI6�40�J6�,/�,�L(uM5��4�rMO,�,�*��ׯ�H/w
�
��2,7��4202�50�5�T04�26�26�3�43���60�26 ��+F��� ^;"L     