package repository

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/database"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/model"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/repository/crudquery"
	"time"
)

type CrudIngridientRepository struct{}

// Create create new ingr in db
func (ingridientRepo *CrudIngridientRepository) Create(context context.Context, ingr model.Ingridient) (bool, error) {
	genUUID, err := uuid.NewV4()
	if err != nil {
		return false, err
	}
	database.DatabaseHolder.Db.MustExec(crudquery.IngridientCreate,
		time.Now(),
		time.Now(),
		genUUID.String(),
		ingr.Name,
	)
	return true, nil
}

// All get list ingres
func (ingridientRepo *CrudIngridientRepository) All() ([]model.Ingridient, error) {
	rows, err := database.DatabaseHolder.Db.Queryx(crudquery.IngridientAll)
	ingres := make([]model.Ingridient, 0)

	if err != nil {
		return ingres, err
	}
	for rows.Next() {
		var ingr model.Ingridient
		err := rows.StructScan(&ingr)
		if err != nil {
			return ingres, err
		}

		mappedingr := model.Ingridient{
			AuditEntity: model.AuditEntity{
				Created: ingr.Created,
				Updated: ingr.Updated,
				Uuid:    ingr.Uuid,
			},
			Name: ingr.Name,
		}

		ingres = append(ingres, mappedingr)
	}

	return ingres, nil
}

// Update update ingr
func (ingridientRepo *CrudIngridientRepository) Update(ctx context.Context, ingr model.Ingridient) (bool, error) {
	res, err := database.DatabaseHolder.Db.Exec(crudquery.IngridientUpdate,
		time.Now(),
		ingr.Name,
		ingr.Uuid,
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

// Delete delete ingr by uuid
func (ingridientRepo *CrudIngridientRepository) Delete(uuid string) bool {
	_, err := database.DatabaseHolder.Db.Exec(crudquery.IngridientDelete, uuid)
	return err == nil
}

// FindByUUID find ingr by uuid
func (ingridientRepo *CrudIngridientRepository) FindByUUID(uuid string) (*model.Ingridient, error) {
	row := database.DatabaseHolder.Db.QueryRow(crudquery.IngridientFindByUUID, uuid)
	var d model.Ingridient
	err := row.Scan(&d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
