package repository

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/database"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/model"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/repository/crudquery"
	"time"
)

// Dish repo
type CrudDishRepository struct{}

// Create create new dish in db
func (dishRepo *CrudDishRepository) Create(context context.Context, dish model.Dish) (bool, error) {
	//userId := context.Value("userId")
	genUUID, err := uuid.NewV4()
	if err != nil {
		return false, err
	}
	database.DatabaseHolder.Db.MustExec(crudquery.DishCreate,
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
func (dishRepo *CrudDishRepository) All() ([]model.Dish, error) {
	rows, err := database.DatabaseHolder.Db.Queryx(crudquery.DishAll)
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
			AuditEntity: model.AuditEntity{
				Created: dish.Created,
				Updated: dish.Updated,
				Uuid:    dish.Uuid,
			},
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
func (dishRepo *CrudDishRepository) Update(context context.Context, dish model.Dish) (bool, error) {
	//userId := context.Value("userId")
	res, err := database.DatabaseHolder.Db.Exec(crudquery.DishUpdate,
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
func (dishRepo *CrudDishRepository) Delete(uuid string) bool {
	_, err := database.DatabaseHolder.Db.Exec(crudquery.DishDelete, uuid)
	return err == nil
}

// FindByUUID find dish by uuid
func (dishRepo *CrudDishRepository) FindByUUID(uuid string) (*model.Dish, error) {
	row := database.DatabaseHolder.Db.QueryRow(crudquery.DishFindByUUID, uuid)
	var d model.Dish
	err := row.Scan(&d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
