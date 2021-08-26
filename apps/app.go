package apps

import (
	"log"
	"net/http"
	"time"

	"github.com/aicelerity-golang/banking/domain"
	"github.com/aicelerity-golang/banking/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()
	dbClient := getDbClient()

	//wiring
	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)
	transactionRepositoryDB := domain.NewTransactionRepositoryDB(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDB)}
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ah := AccountHandlers{service.NewAccountService(accountRepositoryDB)}
	th := TransactionHandlers{service.NewTransactionService(transactionRepositoryDB)}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", th.MakeTransaction).Methods(http.MethodPost)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}

func getDbClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "root:Xaq0!2021@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
