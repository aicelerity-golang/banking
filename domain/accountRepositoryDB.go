package domain

import (
	"strconv"

	"github.com/aicelerity-golang/banking/errs"
	"github.com/aicelerity-golang/banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Accounts) (*Accounts, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, a.Customer_Id, a.Opening_Date, a.Account_Type, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	a.Account_Id = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
