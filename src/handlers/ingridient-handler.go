package handlers

import (
	"encoding/json"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/model"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/repository"
	"io/ioutil"
	"log"
	"net/http"
)

var ingridientRepo repository.CrudIngridientRepository

// IngridientCreateHandler create ingr
func IngridientCreateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Read body error %s", err)
		return
	}

	var ingr model.Ingridient
	err = json.Unmarshal(body, &ingr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Read body error %s", err)
		return
	}

	ok, err := ingridientRepo.Create(r.Context(), ingr)
	if err != nil || !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ingridientRepo.Create is error = %s", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// IngridientAllHandler get list dishes
func IngridientAllHandler(w http.ResponseWriter, r *http.Request) {
	dishes, err := ingridientRepo.All()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ingridientRepo.All is error = %s", err)
		return
	}

	marshal, err := json.Marshal(dishes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ingridientRepo.All is error = %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
}

// IngridientUpdateHandler update ingr
func IngridientUpdateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Read body error %s", err)
		return
	}

	var ingr model.Ingridient
	err = json.Unmarshal(body, &ingr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Unmarshal body error %s", err)
		return
	}

	ok, err := ingridientRepo.Update(r.Context(), ingr)
	if err != nil || !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ingridientRepo.Update error %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// IngridientDeleteHandler delete ingr
func IngridientDeleteHandler(w http.ResponseWriter, r *http.Request) {
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

	if ok := ingridientRepo.Delete(d.Uuid); !ok {
		w.Write([]byte("Не удалось удалить объект " + d.Uuid))
		return
	}
	w.WriteHeader(http.StatusOK)
}
