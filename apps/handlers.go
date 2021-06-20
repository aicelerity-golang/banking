package apps

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/aicelerity-golang/banking/service"
)

// type Customer struct {
// 	Name    string `json:"full_name" xml:"name"`
// 	City    string `json:"city"      xml:"city"`
// 	Zipcode string `json:"zip_code"  xml:"zipcode"`
// }

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"Rajesh", "Bangalore", "68001"},
	// 	{"Balaji", "Bangalore", "68001"},
	// }

	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
