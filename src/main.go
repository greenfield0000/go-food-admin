package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/greenfield0000/go-food/microservices/go-food-admin/database"
	"github.com/greenfield0000/go-food/microservices/go-food-admin/handlers"
	menu_integration "github.com/greenfield0000/go-food/microservices/go-food-admin/integration/menu-integration"
	"github.com/greenfield0000/go-secure-microservice"
)

func init() {
	os.Setenv("ACCESS_SECRET", "mySecretTempKey") //this should be in an env file
	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "mySecretTempKey") //
	database.StartAutoMigrate()
}

// started server function
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server started"))
}

func main() {
	defer database.DatabaseHolder.Db.Close()

	//// with test func header
	http.HandleFunc("/", rootHandler)

	// dish
	http.HandleFunc("/dish/create", middleware(handlers.DishCreateHandler))
	http.HandleFunc("/dish/all", middleware(handlers.DishAllHandler))
	http.HandleFunc("/dish/update", middleware(handlers.DishUpdateHandler))
	http.HandleFunc("/dish/delete", middleware(handlers.DishDeleteHandler))
	http.HandleFunc("/dish/loadJournal", middleware(handlers.DishAllHandler))
	// ingridient
	http.HandleFunc("/ingridient/create", middleware(handlers.IngridientCreateHandler))
	http.HandleFunc("/ingridient/all", middleware(handlers.IngridientAllHandler))
	http.HandleFunc("/ingridient/update", middleware(handlers.IngridientUpdateHandler))
	http.HandleFunc("/ingridient/delete", middleware(handlers.IngridientDeleteHandler))
	// menu
	http.HandleFunc("/menu/create", middleware(handlers.MenuCreateHandler))
	http.HandleFunc("/menu/all", middleware(handlers.MenuAllHandler))
	// menu integration
	http.HandleFunc("/integration/menu", middleware(menu_integration.MenuIntegrationHandler))

	log.Fatalln(http.ListenAndServe(getServicePort(), nil))
}

// middleware function of proxy request, response mechanism
func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Run middleware start")
		defer log.Println("Run middleware finish")

		accessDetails, err := authMiddleWareDetails(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Forbidden"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		context := context.WithValue(r.Context(), "userId", strconv.FormatUint(accessDetails.UserId, 10))
		next.ServeHTTP(w, r.WithContext(context))
	}
}

// authMiddleWareDetails middleware function with auth protect request with details
func authMiddleWareDetails(r *http.Request) (*secure.AccessDetails, error) {
	return secure.ExtractTokenMetadata(r)
}

// getServicePort get port with service listen
func getServicePort() string {
	servicePort := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		servicePort = ":" + port
	}
	return servicePort
}
