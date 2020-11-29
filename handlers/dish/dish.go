package dish

import (
	"encoding/json"
	"github.com/greenfield0000/go-food/back/model"
	"github.com/greenfield0000/go-food/back/repository/dish"
	"io/ioutil"
	"log"
	"net/http"
)

var dishRepo dish.DishRepository

// CreateHandler create dish
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Read body error %s", err)
		return
	}

	var dish model.Dish
	err = json.Unmarshal(body, &dish)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Read body error %s", err)
		return
	}

	ok, err := dishRepo.Create(dish)
	if err != nil || !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("dishRepo.Create is error = %s", err)
		return
	}

	if ok {
		w.WriteHeader(http.StatusCreated)
		return
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// AllHandler get list dishes
func AllHandler(w http.ResponseWriter, r *http.Request) {
	dishes, err := dishRepo.All()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("dishRepo.All is error = %s", err)
		return
	}

	marshal, err := json.Marshal(dishes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("dishRepo.All is error = %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
}
