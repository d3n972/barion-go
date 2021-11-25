package responses

type PaymentResponse struct {
	PaymentID        string `json:"PaymentId"`
	PaymentRequestID string `json:"PaymentRequestId"`
	Status           string `json:"Status"`
	QRURL            string `json:"QRUrl"`
	Transactions     []struct {
		POSTransactionID string      `json:"POSTransactionId"`
		TransactionID    string      `json:"TransactionId"`
		Status           string      `json:"Status"`
		Currency         string      `json:"Currency"`
		TransactionTime  string      `json:"TransactionTime"`
		RelatedID        interface{} `json:"RelatedId"`
	} `json:"Transactions"`
	RecurrenceResult      string        `json:"RecurrenceResult"`
	ThreeDSAuthClientData interface{}   `json:"ThreeDSAuthClientData"`
	GatewayURL            string        `json:"GatewayUrl"`
	RedirectURL           string        `json:"RedirectUrl"`
	CallbackURL           string        `json:"CallbackUrl"`
	TraceID               interface{}   `json:"TraceId"`
	Errors                []interface{} `json:"Errors"`
}
