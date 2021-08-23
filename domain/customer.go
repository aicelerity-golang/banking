package domain

import (
	"github.com/aicelerity-golang/banking/dto"
	"github.com/aicelerity-golang/banking/errs"
)

type Customers struct {
	Customer_Id   int
	Name          string
	City          string
	Zipcode       string
	Date_of_Birth string
	Status        int
}

func (c Customers) statusAsText() string {
	statusAsText := "active"
	if c.Status == 0 {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customers) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Customer_Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.Date_of_Birth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	// status == 1, status == 0, status ==""
	FindAll(status string) ([]Customers, *errs.AppError)
	ById(string) (*Customers, *errs.AppError)
}
