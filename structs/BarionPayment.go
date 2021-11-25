package structs

import (
	"d3n/barion/responses"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gojek/heimdall/httpclient"
	"github.com/google/uuid"
)

type PaymentInterface interface {
	Pay()
}
type BarionPayment struct {
	PaymentInterface
	K          PaymentStart
	httpclient httpclient.Client
	POSKey     string
}

func (B *BarionPayment) InstantPay(poskey string, amount int32) {
	B.httpclient = *httpclient.NewClient()
	B.K.PaymentType = "Immediate"
	B.K.Currency = "HUF"
	B.K.GuestCheckOut = true
	B.K.POSKey = poskey
	B.K.RedirectUrl = "http://127.0.0.1:52300/redir"
	B.K.CallbackUrl = "http://127.0.0.1:52300/ipc"
	B.K.PaymentRequestId = "PRI-" + uuid.NewString()
	k := PaymentTransaction{}
	k.AddItem(Item{
		Name:        "Instant Payment",
		Description: "Instant payment for xyz",
		Quantity:    1,
		Unit:        "pc",
		UnitPrice:   float32(amount),
		ItemTotal:   float32(amount),
		SKU:         "INSTP-1",
	})
	k.POSTransactionId = "TRA-" + uuid.NewString()
	k.Total = float32(amount)
	k.Payee = "d3n+bariontest@d3n.it"

	B.K.FundingSources = append(B.K.FundingSources, "All")

	B.K.AddItem(k)
	B.K.OrderNumber = "ORD-" + uuid.NewString()
	foo_marshalled, _ := json.Marshal(B.K)
	sjson := string(foo_marshalled)
	println(sjson)
	h := http.Header{}
	h.Add("Content-Type", "application/json")
	h.Add("Accept", "application/json")
	reader := strings.NewReader(sjson)
	if reader == nil {
		panic("reader is nil")
	}
	r, err := B.httpclient.Post(
		"https://api.test.barion.com/v2/Payment/Start",
		reader,
		h,
	)
	if err != nil {
		panic(err)
	}
	if b, err := io.ReadAll(r.Body); err == nil {
		rM := responses.PaymentResponse{}
		_ = json.Unmarshal(b, &rM)
		println(rM.GatewayURL)

	}
}
func (B *BarionPayment) GetPaymentStatus(statusId string) responses.PaymentStatus {
	B.httpclient = *httpclient.NewClient()
	ret := responses.PaymentStatus{}
	u, err := url.Parse("https://api.test.barion.com/v2/Payment/GetPaymentState")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("POSKey", B.POSKey)
	q.Set("PaymentId", statusId)
	u.RawQuery = q.Encode()
	h := http.Header{}
	h.Add("Content-Type", "application/json")
	h.Add("Accept", "application/json")
	r, err := B.httpclient.Get(
		u.String(),
		h,
	)
	if err != nil {
		panic(err)
	}
	if b, err := io.ReadAll(r.Body); err == nil {
		_ = json.Unmarshal(b, &ret)

	}
	return ret
}
