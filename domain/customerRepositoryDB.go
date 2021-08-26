package domain

import (
	"database/sql"

	"github.com/aicelerity-golang/banking/errs"
	"github.com/aicelerity-golang/banking/logger"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

// Function to Find All Customers by Status
func (d CustomerRepositoryDB) FindAll(status string) ([]Customers, *errs.AppError) {

	var err error
	customers := make([]Customers, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return customers, nil
}

// Function to Find a Customer by Id
func (d CustomerRepositoryDB) ById(id string) (*Customers, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customers
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning Customer" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

// func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {

// 	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

// 	rows, err := d.client.Query(findAllSql)

// 	if err != nil {
// 		logger.Error("Error while querying customer table" + err.Error())
// 		return nil, err
// 	}

// 	customers := make([]Customer, 0)
// 	for rows.Next() {
// 		var c Customer
// 		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
// 		if err != nil {
// 			logger.Error("Error while querying customer table" + err.Error())
// 			return nil, err
// 		}
// 		customers = append(customers, c)
// 	}
// 	return customers, nil
// }

// func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
// 	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

// 	row := d.client.QueryRow(customerSql, id)
// 	var c Customer
// 	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, errs.NewNotFoundError("Customer not found")
// 		} else {
// 			logger.Error("Error while scanning Customer" + err.Error())
// 			return nil, errs.NewUnexpectedError("Unexpected database error")
// 		}
// 	}
// 	return &c, nil

// }

func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{dbClient}
}
