package main

import (
	"log"
	"net/http"

	"github.com/TaiwoDevOps/go-crud/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// creating a CRUD application

func main() {
	// create a router instance to pass to route method
	r := mux.NewRouter()
	// call the routes method from route package
	routes.RegisterBookStoreRoutes(r)
	// create a base route
	http.Handle("/", r)
	// call the listen and serve method
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
