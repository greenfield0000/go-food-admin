package model

//create table if not exists k_dish
//(
//id      bigserial not null constraint k_dish_pkey primary key,
//created timestamp,
//status  varchar(255),
//updated timestamp,
//uuid    varchar(255),
//cost    double precision,
//name    varchar(255),
//picture varchar(255),
//weight  integer
//)

type Dish struct {
	Id int64
}

//
//create table if not exists k_ingridient
//(
//id      bigserial not null constraint k_ingridient_pkey primary key,
//created timestamp,
//status  varchar(255),
//updated timestamp,
//uuid    varchar(255),
//name    varchar(255),
//weight  integer
//)
type Ingridient struct {
	Id int64
}

type DishIngredient struct {
	Dishid       int64 `db:"dishid"`
	IngridientId int64 `db:"ingridientid"`
}
