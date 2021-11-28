package main

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	id        uuid.UUID
	dateTime  time.Time
	category  uuid.UUID
	memo      string
	value     []uint8
}

func AllTransactions(db *sql.DB) ([]Transaction, error) {
	rows, err := db.Query("SELECT * FROM transaction;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []Transaction

	for rows.Next() {
		var transaction Transaction

		err := rows.Scan(
			&transaction.id,
			&transaction.dateTime,
			&transaction.category,
			&transaction.memo,
			&transaction.value,
		)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
