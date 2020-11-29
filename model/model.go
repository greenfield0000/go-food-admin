package model

import "time"

// Dish database entity
type Dish struct {
	Id      int64 `json:"id,omitempty"`
	Created time.Time
	Updated time.Time
	Uuid    string  `json:"uuid"`
	Cost    float32 `json:"cost"`
	Name    string  `json:"name"`
	Picture string  `json:"picture"`
	Weight  uint    `json:"weight"`
}

// Ingridient database entity
type Ingridient struct {
	Id      int64
	Created time.Time
	Updated time.Time
	Uuid    string
	Name    string
	Weight  uint
}

// DishIngredient database entity
type DishIngredient struct {
	DishId       int64 `db:"dishid"`
	IngridientId int64 `db:"ingridientid"`
}
