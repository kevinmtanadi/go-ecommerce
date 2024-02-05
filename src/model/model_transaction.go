package model

type Transaction struct {
}

func (Transaction) TableName() string {
	return "transaction"
}

type TransactionDetail struct {
}

func (TransactionDetail) TableName() string {
	return "transaction_detail"
}
