package http

import (
	"net/http"

	"github.com/laienderCamargos/Planejamento_Financeiro_Go/adapter/http/actuator"
	"github.com/laienderCamargos/Planejamento_Financeiro_Go/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// init
func Init() {
	http.HandleFunc("/transaction", transaction.GetTransactions)
	http.HandleFunc("/transaction/create", transaction.CreateTransaction)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
