package responses

import "time"

type PaymentStatus struct {
	PaymentID          string      `json:"PaymentId"`
	PaymentRequestID   string      `json:"PaymentRequestId"`
	OrderNumber        string      `json:"OrderNumber"`
	POSID              string      `json:"POSId"`
	POSName            string      `json:"POSName"`
	POSOwnerEmail      string      `json:"POSOwnerEmail"`
	POSOwnerCountry    string      `json:"POSOwnerCountry"`
	Status             string      `json:"Status"`
	PaymentType        string      `json:"PaymentType"`
	FundingSource      string      `json:"FundingSource"`
	RecurrenceType     interface{} `json:"RecurrenceType"`
	TraceID            interface{} `json:"TraceId"`
	FundingInformation struct {
		BankCard struct {
			MaskedPan      string `json:"MaskedPan"`
			BankCardType   string `json:"BankCardType"`
			ValidThruYear  string `json:"ValidThruYear"`
			ValidThruMonth string `json:"ValidThruMonth"`
		} `json:"BankCard"`
		AuthorizationCode string `json:"AuthorizationCode"`
		ProcessResult     string `json:"ProcessResult"`
	} `json:"FundingInformation"`
	AllowedFundingSources []string    `json:"AllowedFundingSources"`
	GuestCheckout         bool        `json:"GuestCheckout"`
	CreatedAt             time.Time   `json:"CreatedAt"`
	ValidUntil            time.Time   `json:"ValidUntil"`
	CompletedAt           time.Time   `json:"CompletedAt"`
	ReservedUntil         interface{} `json:"ReservedUntil"`
	DelayedCaptureUntil   interface{} `json:"DelayedCaptureUntil"`
	Transactions          []struct {
		TransactionID    string    `json:"TransactionId"`
		POSTransactionID string    `json:"POSTransactionId"`
		TransactionTime  time.Time `json:"TransactionTime"`
		Total            float64   `json:"Total"`
		Currency         string    `json:"Currency"`
		Payer            struct {
			Name struct {
				LoginName        string      `json:"LoginName"`
				FirstName        interface{} `json:"FirstName"`
				LastName         interface{} `json:"LastName"`
				OrganizationName interface{} `json:"OrganizationName"`
			} `json:"Name"`
			Email string `json:"Email"`
		} `json:"Payer"`
		Payee struct {
			Name struct {
				LoginName        string      `json:"LoginName"`
				FirstName        string      `json:"FirstName"`
				LastName         string      `json:"LastName"`
				OrganizationName interface{} `json:"OrganizationName"`
			} `json:"Name"`
			Email string `json:"Email"`
		} `json:"Payee"`
		Comment         string `json:"Comment"`
		Status          string `json:"Status"`
		TransactionType string `json:"TransactionType"`
		Items           []struct {
			Name        string  `json:"Name"`
			Description string  `json:"Description"`
			Quantity    float64 `json:"Quantity"`
			Unit        string  `json:"Unit"`
			UnitPrice   float64 `json:"UnitPrice"`
			ItemTotal   float64 `json:"ItemTotal"`
			Sku         string  `json:"SKU"`
		} `json:"Items"`
		RelatedID interface{} `json:"RelatedId"`
		POSID     string      `json:"POSId"`
		PaymentID string      `json:"PaymentId"`
	} `json:"Transactions"`
	Total           float64       `json:"Total"`
	SuggestedLocale string        `json:"SuggestedLocale"`
	FraudRiskScore  interface{}   `json:"FraudRiskScore"`
	RedirectURL     string        `json:"RedirectUrl"`
	CallbackURL     string        `json:"CallbackUrl"`
	Currency        string        `json:"Currency"`
	Errors          []interface{} `json:"Errors"`
}
