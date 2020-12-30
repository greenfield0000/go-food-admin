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

func (r MenuRepository) IsExistDateCollision(startDate *dto.Date, finishDate *dto.Date) bool {
	if startDate == nil || finishDate == nil {
		return true
	}
	var count int64
	err := database.DatabaseHolder.Db.QueryRow(customquery.MenuPropertyCheckDateCollision,
		startDate.Time,
		finishDate.Time,
	).Scan(&count)
	return err != nil || count != 0
}
