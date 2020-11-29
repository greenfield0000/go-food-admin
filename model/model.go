package model

import (
	"time"
)

// Dish database entity
type Dish struct {
	Id      int64 `json:"id,omitempty"`
	Created time.Time
	Updated time.Time
	Uuid    string  `json:"uuid,omitempty"`
	Cost    float32 `json:"cost,omitempty"`
	Name    string  `json:"name,omitempty"`
	Picture string  `json:"picture,omitempty"`
	Weight  uint    `json:"weight,omitempty"`
}

// Ingridient database entity
type Ingridient struct {
	Id      int64 `json:"id,omitempty"`
	Created time.Time
	Updated time.Time
	Uuid    string `json:"uuid"`
	Name    string
	Weight  uint
}

// DishIngredient database entity
type DishIngredient struct {
	DishId       int64 `db:"dishid"`
	IngridientId int64 `db:"ingridientid"`
}
