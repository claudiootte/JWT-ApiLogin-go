package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/claudiootte/JWT-ApiLogin-go/handler"
	"github.com/gorilla/mux"
)

func InitializeRouter() {
	MyRouter := mux.NewRouter()

	MyRouter.HandleFunc("/login", handler.Login)
	MyRouter.HandleFunc("/home", handler.Home)
	MyRouter.HandleFunc("/refresh", handler.Refresh)

	fmt.Println("Connecting to server...")
	log.Fatal(http.ListenAndServe(":8080", MyRouter))
}
