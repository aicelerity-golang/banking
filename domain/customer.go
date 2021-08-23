package domain

import "github.com/aicelerity-golang/banking/errs"

type Customer struct {
	Id          int    `json:"id" 		 xml:"id"`
	Name        string `json:"full_name" xml:"name"`
	City        string `json:"city"      xml:"city"`
	Zipcode     string `json:"zip_code"  xml:"zipcode"`
	DateofBirth string `json:"dob"  	 xml:"dob"`
	Status      int    `json:"status"  	 xml:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
