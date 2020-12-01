package migration

var Schema = `

create table if not exists k_dish
(
    id      bigserial not null constraint k_dish_pkey primary key,
    created timestamp,
    updated timestamp,
    uuid    varchar(255),
    cost    double precision,
    name    varchar(255),
    picture varchar(255),
    weight  integer
);

alter table k_dish owner to admin;

create table if not exists k_ingridient
(
    id      bigserial not null constraint k_ingridient_pkey primary key,
    created timestamp,
    updated timestamp,
    uuid    varchar(255),
    name    varchar(255)
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
