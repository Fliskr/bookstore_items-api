package app

import (
	"net/http"

	"github.com/gervi/bookstore_items-api/src/controllers"
)

func mapUrls() {
	r := router

	r.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)

	r.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	r.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
	r.HandleFunc("/items/search", controllers.ItemsController.Search).Methods(http.MethodPost)
	r.HandleFunc("/items/{id}", controllers.ItemsController.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/items", controllers.ItemsController.Put).Methods(http.MethodPut)
}
