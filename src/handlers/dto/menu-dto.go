package dto

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
	StartDate  Date `json:"startdate"`  // Дата начала действия меню
	FinishDate Date `json:"finishdate"` // Дата окончания действия меню
}

// Dish блюдо для меню
type Dish struct {
	Id        string `json:"id"`          // Уникальный идентификатор меню UUID
	EatTypeId string `json:"eat_type_id"` // Уникальный идентификатор типа приема пищи UUID
}

// Menu описатель меню
type Menu struct {
	Property     Property     `json:"property"`      // Свойства меню
	DishesCreate []model.Dish `json:"dishes_create"` // Список блюд для создания
	Dishes       []Dish       `json:"dishes"`        // Список блюд для назначения
}

type MenuDish struct {
	Name    string `json:"name"`
	Cost    int64  `json:"cost"`
	Picture string `json:"picture"`
	Weight  int64  `json:"weight"`
}

type MenuItem struct {
	Category string     `json:"category"`
	MenuDish []MenuDish `json:"menu_dish"`
}

type Bundle struct {
	Property Property   `json:"property"`
	Menu     []MenuItem `json:"menu_items"`
}

// kd.weight as dish_weight, ke.name as eat_name, kdc.name as dish_category_name
type MenuAllCustom struct {
	StartDate        *time.Time   `db:"startdate"`
	FinishDate       *time.Time   `db:"finishdate"`
	DishName         *string `db:"dish_name"`
	DishCost         *int64  `db:"dish_cost"`
	DishWeight       *int64  `db:"dish_weight"`
	EatName          *string `db:"eat_name"`
	DishCategoryName *string `db:"dish_category_name"`
}

// MenuCreateRequest Запрос для создания меню
type MenuCreateRequest struct {
	Menu Menu `json:"menu"`
}

// MenuAllResponse ответ списка меню
type MenuAllResponse struct {
	Bundle []Bundle `json:"bundle"`
}
