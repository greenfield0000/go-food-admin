package model

import (
	"time"
)

// AuditEntity database entity
type AuditEntity struct {
	Id      int64     `json:"id,omitempty"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Uuid    string    `json:"uuid,omitempty" db:"uuid"`
	UserId  int64     `json:"userid,omitempty" db:"userid"`
}

// Menu database entity
type Menu struct {
	AuditEntity
}

// MenuProperty database entity
type MenuProperty struct {
	AuditEntity
}

// Dish database entity
type Dish struct {
	AuditEntity
	Cost       float32 `json:"cost,omitempty" db:"cost"`
	Name       string  `json:"name,omitempty" db:"name"`
	Picture    *string `json:"picture,omitempty" db:"picture"`
	Weight     int64   `json:"weight,omitempty" db:"weight"`
	CategoryId int64   `json:"category_id,omitempty" db:"category_id"`
}

// Ingridient database entity
type Ingridient struct {
	AuditEntity
	Name string `json:"name,omitempty" db:"name"`
}

// DishIngredient database entity
type DishIngredient struct {
	AuditEntity
	DishId       int64 `db:"dishid"`
	IngridientId int64 `db:"ingridientid"`
}
