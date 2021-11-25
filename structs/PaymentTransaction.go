package structs

type PaymentTransaction struct {
	POSTransactionId  string
	Payee             string
	Total             float32
	Comment           string
	PayeeTransactions []PayeeTransaction
	Items             []Item
}

func (box *PaymentTransaction) AddPayeeTrans(item PayeeTransaction) []PayeeTransaction {
	box.PayeeTransactions = append(box.PayeeTransactions, item)
	return box.PayeeTransactions
}
func (box *PaymentTransaction) AddItem(item Item) []Item {
	box.Items = append(box.Items, item)
	return box.Items
}
