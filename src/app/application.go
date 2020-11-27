package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gervi/bookstore_items-api/src/clients/elastic_client"
	// _ "github.com/gervi/bookstore_items-api/src/clients/mongo"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
	r      = router
)

func StartApplication() {
	elastic_client.Init()
	mapUrls()

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:3335",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Server started at %s\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
