PGDMP     	    (            
    x         
   kriya_test %   10.15 (Ubuntu 10.15-0ubuntu0.18.04.1)    12.0     H           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            I           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            J           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            K           1262    17725 
   kriya_test    DATABASE     t   CREATE DATABASE kriya_test WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'C.UTF-8' LC_CTYPE = 'C.UTF-8';
    DROP DATABASE kriya_test;
                postgres    false            L           0    0    DATABASE kriya_test    ACL     *   GRANT ALL ON DATABASE kriya_test TO test;
                   postgres    false    2891            �            1259    17727    roles    TABLE     �   CREATE TABLE public.roles (
    id character(36) NOT NULL,
    data jsonb NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone
);
    DROP TABLE public.roles;
       public            test    false            �            1259    17737    users    TABLE       CREATE TABLE public.users (
    id character(36) NOT NULL,
    data jsonb NOT NULL,
    role_id character(36) NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone
);
    DROP TABLE public.users;
       public            test    false            D          0    17727    roles 
   TABLE DATA           M   COPY public.roles (id, data, created_at, updated_at, deleted_at) FROM stdin;
    public          test    false    196   �       E          0    17737    users 
   TABLE DATA           V   COPY public.users (id, data, role_id, created_at, updated_at, deleted_at) FROM stdin;
    public          test    false    197   �       �
           2606    17736    roles roles_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.roles DROP CONSTRAINT roles_pkey;
       public            test    false    196            �
           2606    17746    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            test    false    197            �
           2606    17747    users users_role_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.roles(id) ON UPDATE CASCADE ON DELETE CASCADE;
 B   ALTER TABLE ONLY public.users DROP CONSTRAINT users_role_id_fkey;
       public          test    false    2759    196    197            D   �   x�}����0��}���S2I����OPX3�l+i=���E��>��%I���G �A�Kv&s���C���?cX�J��ҙ�K�oK?�[8���uQ۷���¿��pQ�y<w�X��B�zlkgm��S�+w�DdH��O!�@bI1H���W�)��F�_Y%.�[���Y��������G�U�      E   7  x���ko�H�?7�"B��c�~�S!@b�BTi57'`��CI��ޡݮ�$Bjҕ,K�w4�����h�� H4ah�(�m"�u�;�?|���Ng��ae�����4�o�c�]�+V��^�E��k�J}�Z���.��"w�%��>B�(j��7W3;��]�����,��(�坆�H�Atv�ƨ���X6����_���3=�._�NW�}pL��$P%�*�C �!�<�1�D1�!^�,��IƞK��e�GA��?p�����Pi��+��PX��(/]��~/���ͫy�۬7c۰�]�*-�~޹��Еo�[|5<�����F&S�Я�JD"#� q8�A� ʡ���׻(��Rb*�X
"�I��u�+�p:�!��	uHI��vl�R�}���~�I���n +_�~NzK��6�F)��>)��<B�K�@P�<��z`��{O��n�_��ٛ3�.�n!L����?��������5,"�F|�� Ep�}��ӹC
4ZE���e3���������˲�=��bP���:�qZG��S��ե���6-r�����kN��H��*�F��E���}R ��pT)�$�*`�M ���8v��Bj��w��j?5����Ԑ�5�;��m[����~�M�O�Y9S���?ƭ}�B����P��e�+V^h協�)����	�ل�����d�^�o��Pm";I�M�y�hJ�n�Z��}�R�O�K�;��mn��,�'�T_t�?�JE��J��I�T�)z�4���2A��2�R�C�
� ⽨�+�]��Q��b$��'�(iZ6���RW��V��2��S�9(������
H"���|�Pa�C_�	�.�2��if�r�c���a�彤�
M;Y�>�YLT�t�w�.*a��V��v��7���Fg�%���A
ӈI��I�T�(aNb��9���@L��C���n�f��/�	���S{m��g�-����5å�Xv�h�f�F�����j|r֘MG'��^4����B[����H$(Ɣ�������wE�t     