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
`