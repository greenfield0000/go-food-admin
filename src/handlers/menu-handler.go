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
	menuRepository repository.MenuRepository
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

func processBundle(list []dto.MenuAllCustom) []dto.Bundle {
	if list == nil {
		return []dto.Bundle{}
	}
	bundles := make([]dto.Bundle, 0)

	// TODO формирование бандла !

	return bundles
}
