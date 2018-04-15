package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/harsain/input-svc/internal/db"
	"github.com/harsain/input-svc/internal/route"
)

func main() {
	db.Init("saml", "us-west-2")

	var router = mux.NewRouter()

	route.RegisterRoutes(router)

	headersOK := handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "Content-Type", "Origin", "Accept-Encoding", "Accept-Language", "Authorization"})
	originOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST"})
	credentialsAllowed := handlers.AllowCredentials()

	fmt.Println("Running Server")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originOK, headersOK, methodsOK, credentialsAllowed)(router)))
}
