package main

import (
	"github.com/greenfield0000/go-food/back/database"
	"github.com/greenfield0000/go-food/back/handlers/dish"
	"github.com/greenfield0000/go-secure-microservice"
	"log"
	"net/http"
	"os"
)

func init() {
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //
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
	http.HandleFunc("/dish/create", middleware(dish.CreateHandler))
	http.HandleFunc("/dish/all", middleware(dish.AllHandler))
	//http.HandleFunc("/dish/update", middleware())
	//http.HandleFunc("/dish/delete", middleware())
	// ingridient
	//http.HandleFunc("/ingridient", middleware())
	//http.HandleFunc("/ingridient/create", middleware())
	//http.HandleFunc("/ingridient/all", middleware())
	//http.HandleFunc("/ingridient/update", middleware())
	//http.HandleFunc("/ingridient/delete", middleware())

	log.Fatalln(http.ListenAndServe(getServicePort(), nil))
}

// middleware function of proxy request, response mechanism
func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Run middleware start")
		defer log.Println("Run middleware finish")

		err := authMiddleWare(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Forbidden"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
}

// authMiddleWare middleware function with auth protect request
func authMiddleWare(r *http.Request) error {
	_, err := secure.ExtractTokenMetadata(r)
	return err
}

// getServicePort get port with service listen
func getServicePort() string {
	servicePort := ":8081"
	if port := os.Getenv("PORT"); port != "" {
		servicePort = ":" + port
	}
	return servicePort
}
