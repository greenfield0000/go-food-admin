package model

import (
	"time"
)

// Dish database entity
type Dish struct {
	Id         int64 `json:"id,omitempty"`
	Created    time.Time
	Updated    time.Time
	Cost       float32 `json:"cost,omitempty" db:"cost"`
	Uuid       string  `json:"uuid,omitempty" db:"uuid"`
	Name       string  `json:"name,omitempty" db:"name"`
	Picture    *string `json:"picture,omitempty" db:"picture"`
	Weight     int64   `json:"weight,omitempty" db:"weight"`
	CategoryId int64   `json:"category_id,omitempty" db:"category_id"`
}

// Ingridient database entity
type Ingridient struct {
	Id      int64 `json:"id,omitempty"`
	Created time.Time
	Updated time.Time
	Uuid    string `json:"uuid,omitempty" db:"uuid"`
	Name    string `json:"name,omitempty" db:"name"`
}

// DishIngredient database entity
type DishIngredient struct {
	DishId       int64 `db:"dishid"`
	IngridientId int64 `db:"ingridientid"`
}
