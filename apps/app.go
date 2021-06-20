package apps

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet) // explicit http Get
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]*}", getCustomer).Methods(http.MethodGet) // customer id numeric

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
