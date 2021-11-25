package structs

type PaymentStart struct {
	POSKey           string
	PaymentType      string
	GuestCheckOut    bool
	FundingSources   []string
	PaymentRequestId string
	RedirectUrl      string
	CallbackUrl      string
	Transactions     []PaymentTransaction
	OrderNumber      string
	Locale           string
	Currency         string
}

func (box *PaymentStart) AddTrans(item PaymentTransaction) []PaymentTransaction {
	box.Transactions = append(box.Transactions, item)
	return box.Transactions
}
func (box *PaymentStart) AddItem(item PaymentTransaction) []PaymentTransaction {
	box.Transactions = append(box.Transactions, item)
	return box.Transactions
}
