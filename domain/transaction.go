package domain

import (
	"github.com/aicelerity-golang/banking/dto"
	"github.com/aicelerity-golang/banking/errs"
)

const WITHDRAWAL = "withdrawal"

type Transactions struct {
	Transaction_Id   string
	Account_Id       string
	Amount           float64
	Transaction_Type string
	Transaction_Date string
}

type TransactionRepository interface {
	SaveTransaction(Transactions) (*Transactions, *errs.AppError)
	FindBy(accountId string) (*Accounts, *errs.AppError)
}

func (t Transactions) IsWithdrawal() bool {
	return t.Transaction_Type == WITHDRAWAL
}

func (t Transactions) ToTransactionResponseDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.Transaction_Id,
		AccountId:       t.Account_Id,
		Amount:          t.Amount,
		TransactionType: t.Transaction_Type,
		TransactionDate: t.Transaction_Date,
	}
}
