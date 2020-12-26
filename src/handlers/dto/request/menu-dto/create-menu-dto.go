package menu_dto

import (
	"fmt"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/model"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

const layout = "02.01.2006"

func (c *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	if s == "null" {
		return
	}
	c.Time, err = time.Parse(layout, s)
	return
}

func (c Date) MarshalJSON() ([]byte, error) {
	if c.Time.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, c.Time.Format(layout))), nil
}

// Property Свойства меню
type Property struct {
	// Дата начала действия меню
	StartDate Date `json:"startdate"`
	// Дата окончания действия меню
	FinishDate Date `json:"finishdate"`
}

// Dish блюдо для меню
type Dish struct {
	// Уникальный идентификатор меню UUID
	Id string `json:"id"`
	// Уникальный идентификатор типа приема пищи UUID
	EatTypeId string `json:"eat_type_id"`
}

// Menu описатель меню
type Menu struct {
	// Свойства меню
	Property Property `json:"property"`
	// Список блюд для создания
	DishesCreate []model.Dish `json:"dishes_create"`
	// Список блюд для назначения
	Dishes []Dish `json:"dishes"`
}

// Запрос для создания меню
type MenuCreateRequest struct {
	Menu Menu `json:"menu"`
}
