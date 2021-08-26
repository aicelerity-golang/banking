package domain

import (
	"fmt"

	"github.com/aicelerity-golang/banking/dto"
	"github.com/aicelerity-golang/banking/errs"
	"github.com/aicelerity-golang/banking/logger"
)

type Accounts struct {
	Account_Id   string
	Customer_Id  string
	Opening_Date string
	Account_Type string
	Amount       float64
	Status       int
}

type AccountRepository interface {
	Save(Accounts) (*Accounts, *errs.AppError)
}

func (a Accounts) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.Account_Id}
}

func (a Accounts) CanWithdraw(amount float64) bool {
	s1 := fmt.Sprintf("%f", a.Amount)
	logger.Info("Acc Bal Amount :" + s1)
	s2 := fmt.Sprintf("%f", amount)
	logger.Info("Transaction Amount :" + s2)
	if a.Amount > amount {
		return true
	} else {
		return false
	}
}
