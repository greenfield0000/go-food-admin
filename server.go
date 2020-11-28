package main

import (
	"github.com/greenfield0000/go-food/back/database"
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
	//// with test func header
	http.HandleFunc("/authtest", middleware(func(writer http.ResponseWriter, request *http.Request) {
	}))
	log.Fatalln(http.ListenAndServe(":8081", nil))
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Run middleware start")
		err := authMiddleWare(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Forbidden"))
			return
		}
		next.ServeHTTP(w, r)
		log.Println("Run middleware finish")
	}
}

func authMiddleWare(r *http.Request) error {
	_, err := secure.ExtractTokenMetadata(r)
	return err
}
