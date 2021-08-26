package domain

import (
	"strconv"

	"github.com/aicelerity-golang/banking/errs"
	"github.com/aicelerity-golang/banking/logger"
	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryDB struct {
	client *sqlx.DB
}

func (d TransactionRepositoryDB) SaveTransaction(t Transactions) (*Transactions, *errs.AppError) {
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	result, _ := tx.Exec("INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)", t.Account_Id, t.Amount, t.Transaction_Type, t.Transaction_Date)

	// updating account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec("UPDATE accounts SET amount = amount - ? where account_id = ?", t.Amount, t.Account_Id)
	} else {
		_, err = tx.Exec("UPDATE accounts SET amount = amount + ? where account_id = ?", t.Amount, t.Account_Id)
	}

	//In case of error RollBack
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction to bank account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// If all good commit
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction to bank account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	transactionId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while getting last  transaction Id" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	account, appErr := d.FindBy(t.Account_Id)
	if err != nil {
		return nil, appErr
	}
	t.Transaction_Id = strconv.FormatInt(transactionId, 10)

	t.Amount = account.Amount
	return &t, nil
}

func (d TransactionRepositoryDB) FindBy(accountId string) (*Accounts, *errs.AppError) {
	sqlSelect := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	var account Accounts
	err := d.client.Get(&account, sqlSelect, accountId)
	if err != nil {
		logger.Error("Error while getting account Id" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}

func NewTransactionRepositoryDB(dbClient *sqlx.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{dbClient}
}
