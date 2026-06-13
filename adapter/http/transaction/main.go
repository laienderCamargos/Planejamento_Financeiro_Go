package transaction

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"

	"github.com/laienderCamargos/Planejamento_Financeiro_Go/model/transaction"
	"github.com/laienderCamargos/Planejamento_Financeiro_Go/util"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return

	}

	w.Header().Set("Content-type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1200.29,
			Type:      0,
			CreatedAt: util.StringToTime("2020-04-05T11:45:26"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)

}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}
