package handlers

import (
	"encoding/json"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/model"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/repository"
	"io/ioutil"
	"log"
	"net/http"
)

var menuRepo repository.MenuRepository

// MenuCreateHandler create menu
func MenuCreateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Read body error %s", err)
		return
	}

	var menu model.Menu
	err = json.Unmarshal(body, &menu)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Read body error %s", err)
		return
	}

	ok, err := menuRepo.Create(r.Context(), menu)
	if err != nil || !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("menuRepo.Create is error = %s", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// MenuAllHandler get list menu
func MenuAllHandler(w http.ResponseWriter, r *http.Request) {
	menuList, err := menuRepo.All()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("menuRepo.All is error = %s", err)
		return
	}

	marshal, err := json.Marshal(menuList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("menuRepo.All is error = %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
}
