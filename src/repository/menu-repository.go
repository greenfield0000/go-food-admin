package repository

import (
	"github.com/greenfield0000/go-food/microservices/go-food-admin/database"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/handlers/dto"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/repository/customquery"
	"time"
)

type MenuRepository struct{}

func (r MenuRepository) All() ([]dto.MenuAllCustom, error) {
	var mList []dto.MenuAllCustom
	rows, err := database.DatabaseHolder.Db.Query(customquery.MenuAll)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var startDate *time.Time
		var finishDate *time.Time
		var dishName *string
		var dishCost *int64
		var dishWeight *int64
		var eatName *string
		var dishCategoryName *string
		err := rows.Scan(&startDate, &finishDate, &dishName, &dishCost, &dishWeight, &eatName, &dishCategoryName)
		if err != nil {
			return nil, err
		}

		mItem := dto.MenuAllCustom{
			StartDate:        startDate,
			FinishDate:       finishDate,
			DishName:         dishName,
			DishCost:         dishCost,
			DishWeight:       dishWeight,
			EatName:          eatName,
			DishCategoryName: dishCategoryName,
		}
		mList = append(mList, mItem)
	}

	return mList, nil
}
