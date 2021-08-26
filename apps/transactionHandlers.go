package apps

import (
	"encoding/json"
	"net/http"

	"github.com/aicelerity-golang/banking/dto"
	"github.com/aicelerity-golang/banking/service"
	"github.com/gorilla/mux"
)

type TransactionHandlers struct {
	service service.TransactionService
}

func (th TransactionHandlers) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	var request dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		request.CustomerId = customerId
		// logger.Info("Account id :" + accountId)
		// logger.Info("Customer id :" + customerId)
		account, appError := th.service.MakeTransaction(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
