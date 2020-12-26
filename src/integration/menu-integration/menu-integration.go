package menu_integration

import (
	"bytes"
	"github.com/gofrs/uuid"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/database"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/model/handbook/dishtype"
	"github.com/tealeg/xlsx/v3"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

const (
	// Индекс колонки, в которой указано имя группы
	GROUPNAMEINDEX = 1
	// Полезные данные начинаются с
	FROMCOLUMN = 2
	// Полезные данные заканчиваются на
	TOCOLUMN = 5

	INSERT_DISH              = "INSERT INTO k_dish (created, updated, uuid, cost, name, weight, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7) returning id;"
	INSERT_INGRIDIDIENT      = "INSERT INTO k_ingridient (created, updated, uuid, name) VALUES ($1, $2, $3, $4) returning id;"
	INSERT_DISH_INGRIDIDIENT = "INSERT INTO k_dish_ingredient (dishid, ingridientid) VALUES ($1, $2);"
)

type ingridientMapper struct {
	Name string
}

type dishMapper struct {
	Name     string
	Weight   int64
	Cost     int64
	ingrList []*ingridientMapper
}

type dishGroup struct {
	Name   string
	dishes []*dishMapper
}

type dishIngridientLink struct {
	InsertedDishId       int64
	InsertedIngridientId int64
}

// MenuIntegrationHandler integrate menu xls file from old standard
func MenuIntegrationHandler(w http.ResponseWriter, r *http.Request) {
	book, _, err := getWordBook(r)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Не удалось прочитать файл!"))
		return
	}

	sheetName := "Набор блюд на неделю"
	sheet := book.Sheet[sheetName]

	lastGroup := ""
	mp := make(map[string]*dishGroup)
	var dg *dishGroup

	_ = sheet.ForEachRow(func(r *xlsx.Row) error {
		// Группу перебираем отдельно
		cell := r.GetCell(GROUPNAMEINDEX)
		if cell.Value == "" {
			return nil
		}
		// Если это строка, то скорее всего это название группы
		if cellType := cell.Type(); cellType == xlsx.CellTypeString {
			lastGroup = cell.Value
			dg := dishGroup{
				Name: lastGroup,
			}
			mp[lastGroup] = &dg
			return nil
		} else {
			dg = mp[lastGroup]
		}
		//Создадим блюдо
		dMapper := dishMapper{}

		for i := FROMCOLUMN; i <= TOCOLUMN; i++ {
			//var group dishGroup
			cell := r.GetCell(i)
			switch i {
			case 2:
				// Наименование блюда
				dMapper.Name = cell.Value
			case 3:
				// Состав
				if dMapper.ingrList == nil {
					dMapper.ingrList = make([]*ingridientMapper, 0)
				}
				ingMapper := ingridientMapper{
					Name: cell.Value,
				}

				dMapper.ingrList = append(dMapper.ingrList, &ingMapper)
			case 4:
				// Вес
				parseUint, _ := strconv.ParseUint(cell.Value, 10, 64)
				dMapper.Weight = int64(parseUint)
			case 5:
				// Стоимость
				cost, _ := strconv.ParseInt(cell.Value, 10, 64)
				dMapper.Cost = cost
			}
		}

		if dg.dishes == nil {
			dg.dishes = make([]*dishMapper, 0)
		}

		// Похеренные записи не добавляем!
		if dMapper.Name != "" {
			dg.dishes = append(dg.dishes, &dMapper)
		}

		return nil
	}, xlsx.SkipEmptyRows)
	writeToDB(mp)
}

// writeToDB function to write dish group data into db
func writeToDB(dGropMap map[string]*dishGroup) {
	if dGropMap != nil {
		links := make([]dishIngridientLink, 0)
		for k, dg := range dGropMap {
			dishes := dg.dishes

			categoryId := dishtype.GetCategoryIndexByName(k)
			if categoryId == -1 {
				log.Println("Не удалось определить категорию!")
				return
			}
			// Вставка блюд
			for _, dish := range dishes {
				insertedDishId := insertDish(dish, categoryId)
				ingList := dish.ingrList
				// Вставка ингридиентов
				for _, ingridient := range ingList {
					insertedIngridientId := insertIngridient(ingridient)

					if insertedDishId != -1 && insertedIngridientId != -1 {
						links = append(links, dishIngridientLink{
							InsertedDishId:       insertedDishId,
							InsertedIngridientId: insertedIngridientId,
						})
					}
					insertedIngridientId = -1
				}
			}
		}

		for _, link := range links {
			database.DatabaseHolder.Db.Exec(INSERT_DISH_INGRIDIDIENT,
				link.InsertedDishId,
				link.InsertedIngridientId,
			)
		}
	}
}

func insertIngridient(ingridient *ingridientMapper) int64 {
	var insertedIngridientId int64 = -1

	database.DatabaseHolder.Db.QueryRowx(INSERT_INGRIDIDIENT,
		time.Now(),
		time.Now(),
		genUUID(),
		ingridient.Name,
	).Scan(&insertedIngridientId)

	return insertedIngridientId
}

func insertDish(dish *dishMapper, categoryId int) int64 {
	var insertedDishId int64 = -1

	database.DatabaseHolder.Db.QueryRowx(INSERT_DISH,
		time.Now(),
		time.Now(),
		genUUID(),
		dish.Cost,
		dish.Name,
		dish.Weight,
		categoryId,
	).Scan(&insertedDishId)

	return insertedDishId
}

func genUUID() string {
	genUUID, _ := uuid.NewV4()
	return genUUID.String()
}

// getWordBook read excel file from request
func getWordBook(r *http.Request) (*xlsx.File, *multipart.FileHeader, error) {
	// 5мб
	r.ParseMultipartForm(5000000)

	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, header, err
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, header, err
	}

	wb, err := xlsx.OpenBinary(buf.Bytes())
	if _, err := io.Copy(buf, file); err != nil {
		return nil, header, err
	}
	return wb, header, nil
}
