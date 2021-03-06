package repository

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/database"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/handlers/dto"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/model"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/repository/crudquery"
	"strconv"
	"time"
)

type CrudMenuRepository struct{}

func (r CrudMenuRepository) Create(context context.Context, menuReq dto.MenuCreateRequest) (bool, error) {
	userId := context.Value("userId").(string)
	// Вставляем свойства
	menu := menuReq.Menu
	property := menu.Property
	parseUint, _ := strconv.ParseUint(userId, 10, 64)
	propertyId, err := createPropertyMenu(parseUint, property)
	if err != nil {
		return false, err
	}
	db := database.DatabaseHolder.Db
	// Вставляем блюда в меню
	for _, d := range menu.Dishes {
		var dish model.Dish
		var eatTypeId uint64

		db.QueryRowx(crudquery.DishFindByUUID, d.Id).StructScan(&dish)
		db.QueryRow(crudquery.EatTypeFindByUUID, d.EatTypeId).Scan(&eatTypeId)

		dishId := dish.Id
		if dishId == 0 || eatTypeId == 0 {
			return false, err
		}

		genUUID, err := uuid.NewV4()
		if err != nil {
			deletePropertyMenu(propertyId)
			return false, err
		}
		_, err = db.Exec(crudquery.MenuInsert,
			genUUID,
			time.Now(),
			time.Now(),
			userId,
			dishId,
			eatTypeId,
			propertyId,
		)
		if err != nil {
			deletePropertyMenu(propertyId)
			return false, err
		}
	}

	return true, nil
}

func (r CrudMenuRepository) All() ([]model.Menu, error) {
	return nil, nil
}

func createPropertyMenu(userId uint64, property dto.Property) (int64, error) {
	genUUID, err := uuid.NewV4()
	if err != nil {
		return 0, err
	}

	db := database.DatabaseHolder.Db
	var id int64
	db.QueryRowx(crudquery.MenuPropetyCreate,
		genUUID.String(),
		time.Now(),
		time.Now(),
		userId,
		property.StartDate.Time,
		property.FinishDate.Time,
	).Scan(&id)

	return id, err
}

func deletePropertyMenu(propertyId int64) {
	database.DatabaseHolder.Db.QueryRow("delete k_menu_property where id = $1",
		propertyId,
	)
}
