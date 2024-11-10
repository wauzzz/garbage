package main

import (
	"fmt"
	"time"
	"mexc-sdk/mexcsdk"
)

func main() {
	apiKey := ""
	apiSecret := ""
	client := mexcsdk.NewSpot(&apiKey, &apiSecret)

	symbol := "KANGOUSDT"
	side := "BUY"
	orderType := "LIMIT"
	t := time.Now().UnixMilli()
	tstart := time.Date(2024, time.November, 8, 15, 00, 0, 0, time.UTC).UnixMilli() - 350
	twait := tstart - t
	fmt.Println(twait)
	tm := time.NewTimer(time.Duration(twait)*time.Millisecond)
	<-tm.C
	options := map[string]interface{}{"quantity": 100000, "price": 0.00002, "recvWindow": 5000}
	res := client.NewOrder(&symbol, &side, &orderType, &options)
	fmt.Println(res)
}
