package apps

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city"      xml:"city"`
	Zipcode string `json:"zip_code"  xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello goWorld")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// fmt.Fprint(w, vars["customer_id"])

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vars["customer_id"])
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Rajesh", "Bangalore", "68001"},
		{"Balaji", "Bangalore", "68001"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}
