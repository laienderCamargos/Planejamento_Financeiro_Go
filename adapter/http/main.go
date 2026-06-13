package http

import (
	"net/http"

	"github.com/laienderCamargos/Planejamento_Financeiro_Go/adapter/http/actuator"
	"github.com/laienderCamargos/Planejamento_Financeiro_Go/adapter/http/transaction"
)

// init
func Init() {
	http.HandleFunc("/transaction", transaction.GetTransactions)
	http.HandleFunc("/transaction/create", transaction.CreateTransaction)

	http.HandleFunc("/health", actuator.Health)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
