package model

import "time"

// Dish database entity
type Dish struct {
	Id      uint64    `db:"id"`
	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`
	Uuid    string    `db:"uuid"`
	Cost    float32   `db:"cost"`
	Name    string    `db:"name"`
	Picture string    `db:"picture"`
	Weight  uint      `db:"weight"`
}

// Ingridient database entity
type Ingridient struct {
	Id      int64     `db:"id"`
	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`
	Uuid    string    `db:"uuid"`
	Name    string    `db:"name"`
	Weight  uint      `db:"weight"`
}

// DishIngredient database entity
type DishIngredient struct {
	DishId       int64 `db:"dishid"`
	IngridientId int64 `db:"ingridientid"`
}
