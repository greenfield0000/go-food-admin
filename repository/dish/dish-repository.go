package dish

import (
	"github.com/gofrs/uuid"
	"github.com/greenfield0000/go-food/back/database"
	"github.com/greenfield0000/go-food/back/model"

	"time"
)

// Dish repo
type DishRepository struct{}

// Create create new dish in db
func (dishRepo *DishRepository) Create(dish model.Dish) (bool, error) {
	genUUID, err := uuid.NewV4()
	if err != nil {
		return false, err
	}
	database.DatabaseHolder.Db.MustExec(Dish_create,
		time.Now(),
		time.Now(),
		genUUID.String(),
		dish.Cost,
		dish.Name,
		dish.Picture,
		dish.Weight,
	)
	return true, nil
}

//All get list dishes
func (dishRepo *DishRepository) All() ([]model.Dish, error) {
	rows, err := database.DatabaseHolder.Db.Queryx(Dish_All)
	dishes := make([]model.Dish, 0)

	if err != nil {
		return dishes, err
	}
	for rows.Next() {
		var dish model.Dish
		err := rows.StructScan(&dish)
		if err != nil {
			return dishes, err
		}

		mappedDish := model.Dish{
			Created: dish.Created,
			Updated: dish.Updated,
			Uuid:    dish.Uuid,
			Cost:    dish.Cost,
			Name:    dish.Name,
			Picture: dish.Picture,
			Weight:  dish.Weight,
		}

		dishes = append(dishes, mappedDish)
	}

	return dishes, nil
}

// Update update dish
func (dishRepo *DishRepository) Update(dish model.Dish) (bool, error) {
	return true, nil
}

// Delete delete dish by id
func (dishRepo *DishRepository) Delete(id uint64) bool {
	return true
}
