package routes

import (
	"github.com/TaiwoDevOps/go-crud/pkg/controllers"
	"github.com/gorilla/mux"
)

// Create a named func
var RegisterBookStoreRoutes = func(router *mux.Router) {
	// Create the routes and call the associated controller
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
