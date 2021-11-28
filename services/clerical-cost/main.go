package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    _ "github.com/lib/pq"
)

type Env struct {
    db *sql.DB
}

func main() {
    db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/clericalcost?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }

    env := &Env{db: db}

    http.HandleFunc("/transactions", env.transactionIndex)
    http.ListenAndServe(":3000", nil)
}

func (env *Env) transactionIndex(w http.ResponseWriter, r *http.Request) {
    transactions, err := AllTransactions(env.db)
    if err != nil {
        log.Println(err)
        http.Error(w, http.StatusText(500), 500)
        return
    }

    for _, transaction := range transactions {
        fmt.Fprintf(w, "%s, %s, %s, %s, %s€\n",
			transaction.id,
			transaction.dateTime,
			transaction.category,
			transaction.memo,
			string(transaction.value),
		)
    }
}
