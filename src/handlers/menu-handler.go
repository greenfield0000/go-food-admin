package handlers

import (
	"encoding/json"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/handlers/dto"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/repository"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	crudMenuRepository repository.CrudMenuRepository
	menuRepository     repository.MenuRepository
)

// MenuCreateHandler create menu
func MenuCreateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Read body error %s", err)
		return
	}

	var createMenuRequest dto.MenuCreateRequest
	err = json.Unmarshal(body, &createMenuRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Read body error %s", err)
		return
	}

	ok, err := crudMenuRepository.Create(r.Context(), createMenuRequest)
	if err != nil || !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("crudMenuRepository.Create is error = %s", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// MenuAllHandler get list menu
func MenuAllHandler(w http.ResponseWriter, r *http.Request) {
	menuList, err := menuRepository.All()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Read body error %s", err)
		return
	}

	var response = dto.MenuAllResponse{
		Bundle: processBundle(menuList),
	}

	marshal, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("menuRepository.All is error = %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
}

func processBundle(list []dto.MenuAllCustom) []*dto.Bundle {
	if list == nil {
		return []*dto.Bundle{}
	}
	bundles := []*dto.Bundle{}
	setsDate := make(map[string]bool)

	// Как только встречаем новую дату, считаем, что это новый бандл
	var bundleItem dto.Bundle
	var menuCategoryMap map[string]dto.MenuItem
	for i := range list {
		menuRow := list[i]

		dateTime := menuRow.StartDate
		dateTimeString := dateTime.String()
		if _, exist := setsDate[dateTimeString]; !exist {
			bundleItem = dto.Bundle{}
			bundles = append(bundles, &bundleItem)

			setsDate[dateTimeString] = true
			// Установим свойства
			bundleItem.Property = dto.Property{
				StartDate: dto.Date{
					Time: menuRow.StartDate,
				},
				FinishDate: dto.Date{
					menuRow.FinishDate,
				},
			}
		}

		// Если меню еще не создано, создаем
		if bundleItem.MenuItems == nil {
			bundleItem.MenuItems = make([]dto.MenuItem, 0)
		}
		menuItems := bundleItem.MenuItems
		// Категория - список блюд
		if menuCategoryMap == nil {
			menuCategoryMap = make(map[string]dto.MenuItem, 0)
		}
		eatNameType := menuRow.EatName

		menuItem := getMenuItemByEatNameType(menuCategoryMap, eatNameType)
		menuItem.MenuDish = append(menuItem.MenuDish, dto.MenuDish{
			Name:         menuRow.DishName,
			Cost:         menuRow.DishCost,
			Picture:      "",
			Weight:       menuRow.DishWeight,
			DishCategory: menuRow.DishCategoryName,
		})

		menuItems = append(menuItems, menuItem)
		bundleItem.MenuItems = menuItems
	}

	return bundles
}

func getMenuItemByEatNameType(menuCategoryMap map[string]dto.MenuItem, eatNameType string) dto.MenuItem {
	// Если такая категория есть, то вернем ее
	if menuItem, exist := menuCategoryMap[eatNameType]; exist && len(eatNameType) > 0 {
		return menuItem
	}
	// Если такой категории нет, то создадим и вернем ее
	menuItem := dto.MenuItem{CategoryName: eatNameType, MenuDish: make([]dto.MenuDish, 0)}
	menuCategoryMap[eatNameType] = menuItem
	return menuItem
}
