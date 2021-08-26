package service

import (
	"time"

	"github.com/aicelerity-golang/banking/domain"
	"github.com/aicelerity-golang/banking/dto"
	"github.com/aicelerity-golang/banking/errs"
)

type TransactionService interface {
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

func (s DefaultTransactionService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in your account")
		}
	}

	t := domain.Transactions{
		Account_Id:       req.AccountId,
		Amount:           req.Amount,
		Transaction_Type: req.TransactionType,
		Transaction_Date: time.Now().Format(dbTSLayout),
	}
	transaction, err := s.repo.SaveTransaction(t)
	if err != nil {
		return nil, err
	}
	response := transaction.ToTransactionResponseDto()
	return &response, nil
}

func NewTransactionService(repo domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repo}
}
