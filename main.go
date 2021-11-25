package main

import (
	"d3n/barion/structs"
)

func main() {

	poskey := "c5d86d2a03d940b9a7dbd7e6fb90fa39"
	ky := structs.BarionPayment{}
	ky.POSKey = poskey
	//	ky.InstantPay(poskey, 51)
	println(ky.GetPaymentStatus("a124f1fdee4dec118bdb001dd8b71cc4").Status)
}
