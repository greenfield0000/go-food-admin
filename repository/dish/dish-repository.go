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
	database.DatabaseHolder.Db.MustExec(DishCreate,
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
	rows, err := database.DatabaseHolder.Db.Queryx(DishAll)
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
	res, err := database.DatabaseHolder.Db.Exec(DishUpdate,
		time.Now(),
		dish.Cost,
		dish.Name,
		dish.Picture,
		dish.Weight,
		dish.Uuid,
	)
	if err != nil {
		return false, err
	}
	// смотрим, было ли обновление на самом деле
	if affected, _ := res.RowsAffected(); affected == 0 {
		return false, nil
	}
	return true, nil
}

// Delete delete dish by uuid
func (dishRepo *DishRepository) Delete(uuid string) bool {
	_, err := database.DatabaseHolder.Db.Exec(DishDelete, uuid)
	return err == nil
}

// FindByUUID find dish by uuid
func (dishRepo *DishRepository) FindByUUID(uuid string) (*model.Dish, error) {
	row := database.DatabaseHolder.Db.QueryRow(DishFindByUUID, uuid)
	var d model.Dish
	err := row.Scan(&d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
