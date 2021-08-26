package service

import (
	"time"

	"github.com/aicelerity-golang/banking/domain"
	"github.com/aicelerity-golang/banking/dto"
	"github.com/aicelerity-golang/banking/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Accounts{
		Account_Id:   "",
		Customer_Id:  req.CustomerId,
		Opening_Date: time.Now().Format(dbTSLayout),
		Account_Type: req.AccountType,
		Amount:       req.Amount,
		Status:       1,
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToNewAccountResponseDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
