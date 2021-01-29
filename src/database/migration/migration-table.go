package migration

var Schema = `


create table if not exists k_dish_category 
(
    id      bigserial not null constraint k_dish_category_pkey primary key,
    sysname varchar(25),
    name    varchar(100)
);

comment on table k_dish_category is 'Категории блюд';

alter table k_dish_category
    owner to admin;

INSERT INTO k_dish_category (id, sysname, name) VALUES (1, 'Salad', 'Салаты') on conflict (id) do nothing;;
INSERT INTO k_dish_category (id, sysname, name) VALUES (2, 'FirstMeal', 'Первые блюда') on conflict (id) do nothing;;
INSERT INTO k_dish_category (id, sysname, name) VALUES (3, 'SecondMeal', 'Вторые блюда') on conflict (id) do nothing;;
INSERT INTO k_dish_category (id, sysname, name) VALUES (4, 'SideDish', 'Гарниры') on conflict (id) do nothing;;
INSERT INTO k_dish_category (id, sysname, name) VALUES (5, 'Bread', 'Хлеб') on conflict (id) do nothing;;
INSERT INTO k_dish_category (id, sysname, name) VALUES (6, 'Bake', 'Выпечка') on conflict (id) do nothing;;
INSERT INTO k_dish_category (id, sysname, name) VALUES (7, 'Confectionery', 'Кондитерские изделия') on conflict (id) do nothing;;
INSERT INTO k_dish_category (id, sysname, name) VALUES (8, 'Cakes', 'Торты') on conflict (id) do nothing;;

create  table if not exists k_dish
(
    id      bigserial not null constraint k_dish_pkey primary key,
    created timestamp,
    updated timestamp,
    uuid    varchar(255),
    cost    double precision,
    name    varchar(255) unique,
    picture varchar(255),
    weight  integer,
	category_id bigint not null constraint category_id_fk references k_dish_category(id)
);

alter table k_dish owner to admin;

create  table if not exists k_ingridient
(
    id      bigserial not null constraint k_ingridient_pkey primary key,
    created timestamp,
    updated timestamp,
    uuid    varchar(255),
    name    varchar(255) unique
);

alter table k_ingridient owner to admin;

create table if not exists k_dish_ingredient
(
    dishid       bigint not null constraint fkhup5cd1ujgsf64u3jarlbimbf references k_dish,
    ingridientid bigint not null constraint fkp1hsrracg6eswaj4gelwpmr2q references k_ingridient,
    constraint k_dish_ingredient_pkey primary key (dishid, ingridientid)
);

alter table k_dish_ingredient owner to admin;

create table if not exists k_eattype
(
    id      bigserial not null
        constraint k_eattype_pk
            primary key,
    uuid    varchar(40),
    created timestamp,
    updated timestamp,
    userid  bigint,
    name    varchar(255),
    sysname varchar(255)
);

alter table k_eattype
    owner to admin;

INSERT INTO k_eattype (id, uuid, created, updated, userid, name, sysname) VALUES (1, '93dc4394-47c9-11eb-ad37-430f9829063b', '2020-12-27 01:28:19.000000', '2020-12-27 01:28:19.000000', 1, 'Завтрак', 'Breakfast') on conflict (id) do nothing;;
INSERT INTO k_eattype (id, uuid, created, updated, userid, name, sysname) VALUES (2, '9aca1a32-47c9-11eb-b944-f375365165a0', '2020-12-27 01:28:19.000000', '2020-12-27 01:28:19.000000', 1, 'Обед', 'Lunch') on conflict (id) do nothing;;
INSERT INTO k_eattype (id, uuid, created, updated, userid, name, sysname) VALUES (3, 'a07c6502-47c9-11eb-bccb-cfad46344a6a', '2020-12-27 01:28:19.000000', '2020-12-27 01:28:19.000000', 1, 'Ужин', 'Dinner') on conflict (id) do nothing;;

create table if not exists k_menu_property
(
    id         bigserial not null
        constraint k_menu_property_pk
            primary key,
    uuid       varchar(40),
    created    timestamp,
    updated    timestamp,
    userid     bigint,
    startdate  timestamp,
    finishdate timestamp
);

alter table k_menu_property
    owner to admin;

create table if not exists k_dayofweek
(
    id      bigserial not null
        constraint k_dayofweek_pk
            primary key,
    uuid    varchar(40),
    created timestamp,
    updated timestamp,
    userid  bigint,
    name    varchar(255),
    sysname varchar(255)
);

alter table k_dayofweek
    owner to admin;

INSERT INTO k_dayofweek (id, sysname, name) VALUES (1, 'Monday', 'Понедельник') on conflict (id) do nothing;;
INSERT INTO k_dayofweek (id, sysname, name) VALUES (2, 'Tuesday', 'Вторник') on conflict (id) do nothing;;
INSERT INTO k_dayofweek (id, sysname, name) VALUES (3, 'Wednesday', 'Среда') on conflict (id) do nothing;;
INSERT INTO k_dayofweek (id, sysname, name) VALUES (4, 'Thursday', 'Четверг') on conflict (id) do nothing;;
INSERT INTO k_dayofweek (id, sysname, name) VALUES (5, 'Friday', 'Пятница') on conflict (id) do nothing;;
INSERT INTO k_dayofweek (id, sysname, name) VALUES (6, 'Saturday', 'Суббота') on conflict (id) do nothing;;
INSERT INTO k_dayofweek (id, sysname, name) VALUES (7, 'Sunday', 'Воскресенье') on conflict (id) do nothing;;

create table if not exists k_menu
(
    id               bigserial not null
        constraint k_menu_pk
            primary key,
    uuid             varchar(40),
    created          timestamp,
    updated          timestamp,
    userid           bigint,
    dish_id          bigint
        constraint k_menu_k_dish_id_fk
            references k_dish,
    eat_type_id      bigint
        constraint k_menu_k_eattype_id_fk
            references k_eattype,
    menu_property_id bigint
        constraint k_menu_k_menu_property_id_fk
            references k_menu_property,
    day_of_week_id bigint
        constraint day_of_week_id_fk
            references k_dayofweek,
	constraint k_menu_bundle
        unique (dish_id, eat_type_id, menu_property_id)
);

comment on table k_menu is 'Меню';

alter table k_menu
    owner to admin;

create unique index if not exists k_menu_uuid_uindex
    on k_menu (uuid);
`
