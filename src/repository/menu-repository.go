package repository

import (
	"context"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/model"
)

type MenuRepository struct{}

func (r MenuRepository) Create(context context.Context, menu model.Menu) (bool, error) {
	return false, nil
}

func (r MenuRepository) All() ([]model.Menu, error) {
	return nil, nil
}
