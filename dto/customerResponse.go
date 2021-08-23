package dto

type CustomerResponse struct {
	Id          int    `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     string `json:"zip_code"`
	DateofBirth string `json:"dob"`
	Status      string `json:"status"`
}
