package handlers

import (
	"encoding/json"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/model"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/repository"
	"io/ioutil"
	"log"
	"net/http"
)

var dishRepo repository.CrudDishRepository

// DishCreateHandler create dish
func DishCreateHandler(w http.ResponseWriter, r *http.Request) {
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

	ok, err := dishRepo.Create(r.Context(), dish)
	if err != nil || !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("dishRepo.Create is error = %s", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// DishAllHandler get list dishes
func DishAllHandler(w http.ResponseWriter, r *http.Request) {
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

// DishUpdateHandler update dish
func DishUpdateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Read body error %s", err)
		return
	}

	var dish model.Dish
	err = json.Unmarshal(body, &dish)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Unmarshal body error %s", err)
		return
	}

	ok, err := dishRepo.Update(r.Context(), dish)
	if err != nil || !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("dishRepo.Update error %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DishDeleteHandler delete dish
func DishDeleteHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var d model.Dish
	err = json.Unmarshal(body, &d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Unmarshal body error %s", err)
		return
	}

	if ok := dishRepo.Delete(d.Uuid); !ok {
		w.Write([]byte("Не удалось удалить объект " + d.Uuid))
		return
	}
	w.WriteHeader(http.StatusOK)
}