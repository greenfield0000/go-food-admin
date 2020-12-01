package menu_integration

import (
	"bytes"
	"github.com/tealeg/xlsx/v3"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
)

const (
	// Индекс колонки, в которой указано имя группы
	GROUPNAMEINDEX = 1
	// Полезные данные начинаются с
	FROMCOLUMN = 2
	// Полезные данные заканчиваются на
	TOCOLUMN = 5
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

		dg.dishes = append(dg.dishes, &dMapper)

		return nil
	}, xlsx.SkipEmptyRows)
	log.Println("aasd")
}

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
