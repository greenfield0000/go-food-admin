package repository

import (
	"database/sql"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/database"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/handlers/dto"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/repository/customquery"
)

type MenuRepository struct{}

func (r MenuRepository) All() ([]dto.MenuAllCustom, error) {
	var mList []dto.MenuAllCustom
	rows, err := database.DatabaseHolder.Db.Query(customquery.MenuAll)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var startDate sql.NullTime
		var finishDate sql.NullTime
		var dishName sql.NullString
		var dishCost sql.NullInt64
		var dishWeight sql.NullInt64
		var eatName sql.NullString
		var dishCategoryName sql.NullString
		err := rows.Scan(&startDate, &finishDate, &dishName, &dishCost, &dishWeight, &eatName, &dishCategoryName)
		if err != nil {
			return nil, err
		}

		mItem := dto.MenuAllCustom{
			StartDate:        startDate.Time,
			FinishDate:       finishDate.Time,
			DishName:         dishName.String,
			DishCost:         dishCost.Int64,
			DishWeight:       dishWeight.Int64,
			EatName:          eatName.String,
			DishCategoryName: dishCategoryName.String,
		}
		mList = append(mList, mItem)
	}

	return mList, nil
}
