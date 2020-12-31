package handlers

import (
	"encoding/json"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/handlers/dto"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/repository"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

	// Перед тем как добавить, проверим пересечения по дате
	menu := &createMenuRequest.Menu
	if menu == nil {
		w.Write([]byte("Отсутствует аттрибут menu!"))
		log.Println("Missed json attribute \"menu\"", err)
		return
	}

	if isValid(w, menu, err) {
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

func isValid(w http.ResponseWriter, menu *dto.Menu, err error) bool {
	// TODO Возможно у json кодогенератора есть такой чек

	property := &menu.Property
	if property == nil {
		w.Write([]byte("Отсутствует аттрибут menu.property!"))
		log.Println("Missed json attribute \"menu.property\"", err)
		return true
	}
	startDate := &property.StartDate
	if startDate == nil {
		w.Write([]byte("Отсутствует аттрибут menu.property.startDate!"))
		log.Println("Missed json attribute \"menu.property.startDate\"", err)
		return true
	}

	finishDate := &property.FinishDate
	if finishDate == nil {
		w.Write([]byte("Отсутствует аттрибут menu.property.finishDate!"))
		log.Println("Missed json attribute \"menu.property.finishDate\"", err)
		return true
	}

	if finishDate.Before(startDate.Time) {
		w.Write([]byte("Дата окончания не может быть раньше даты начала!"))
		log.Println("StartDate must be before finishDate", err)
		return true
	}

	dishes := &menu.Dishes
	if dishes == nil {
		w.Write([]byte("Отсутствует аттрибут menu.dishes!"))
		log.Println("Missed json attribute \"menu.dishes\"", err)
		return true
	}
	// TODO Возможно у json кодогенератора есть такой чек

	if menuRepository.IsExistDateCollision(startDate, finishDate) {
		w.Write([]byte("Существует пересечение с другими меню. Добавление невозможно!"))
		log.Println("Exist collision with other menu. Add impossible!", err)
		return true
	}

	return false
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

func processBundle(list []dto.MenuAllCustom) []dto.Bundle {
	if list == nil {
		return []dto.Bundle{}
	}
	// дата начала -> тип приема пищи (обед, завтрак) -> список блюд
	bundleMap := make(map[string]map[string]*dto.MenuItem)
	// старт дату и финиш дату буду хранить так
	dateMap := make(map[string]struct {
		StartDate  time.Time
		FinishDate time.Time
	})
	for i := range list {
		menuRow := list[i]

		startDate := menuRow.StartDate
		finishDate := menuRow.FinishDate

		startDateString := startDate.String()
		if _, ok := dateMap[startDateString]; !ok {
			dateMap[startDateString] = struct {
				StartDate  time.Time
				FinishDate time.Time
			}{StartDate: startDate, FinishDate: finishDate}
		}
		eatNameType := menuRow.EatName

		menuItem := getMenuItem(bundleMap, startDateString, eatNameType)
		menuItem.MenuDish = append(menuItem.MenuDish, &dto.MenuDish{
			Name:         menuRow.DishName,
			Cost:         menuRow.DishCost,
			Picture:      "",
			Weight:       menuRow.DishWeight,
			DishCategory: menuRow.DishCategoryName,
		})
	}

	bundles := make([]dto.Bundle, 0)
	for startDate, value := range dateMap {
		if _, ok := bundleMap[startDate]; ok {
			bundle := dto.Bundle{
				Property: dto.Property{
					StartDate: dto.Date{
						Time: value.StartDate,
					},
					FinishDate: dto.Date{
						Time: value.FinishDate,
					},
				},
				MenuItems: getMenuItems(bundleMap, startDate),
			}
			bundles = append(bundles, bundle)
		}
	}

	return bundles
}

func getMenuItems(bundleMap map[string]map[string]*dto.MenuItem, startDate string) []dto.MenuItem {
	if bundleMap == nil {
		return []dto.MenuItem{}
	}

	menuItems := make([]dto.MenuItem, 0)
	// Если такая категория есть, то вернем ее
	if menuItemMap, exist := bundleMap[startDate]; exist {
		for _, val := range menuItemMap {
			menuItem := dto.MenuItem{
				CategoryName: val.CategoryName,
				MenuDish:     val.MenuDish,
			}
			menuItems = append(menuItems, menuItem)
		}
	}

	return menuItems
}

func getMenuItem(bundleMap map[string]map[string]*dto.MenuItem, startDate string, eatType string) *dto.MenuItem {
	// Если такая категория есть, то вернем ее
	if menuItemMap, exist := bundleMap[startDate]; exist {
		if menuItem, ok := menuItemMap[eatType]; ok {
			return menuItem
		}
	}
	// Если такой категории нет, то создадим и вернем ее
	var menuItem = dto.MenuItem{
		CategoryName: eatType,
		MenuDish:     make([]*dto.MenuDish, 0),
	}
	if menuItemMap, ok := bundleMap[startDate]; !ok {
		if menuItemMap == nil {
			menuItemMap = make(map[string]*dto.MenuItem)
		}
		menuItemMap[eatType] = &menuItem
		bundleMap[startDate] = menuItemMap
	} else {
		if _, ok := menuItemMap[eatType]; !ok {
			menuItemMap[eatType] = &menuItem
		}
	}

	return &menuItem

}
