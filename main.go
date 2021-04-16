package main

import (
	"github.com/jyap808/go-poloniex"
	"log"
	"os"
	"strconv"
	"time"
)

func main()  {
	client := poloniex.New(os.Getenv("POLONIEX_API_KEY"), os.Getenv("POLONIEX_API_SECRET"))


	for {
		tickers, err := client.GetTickers()
		if err != nil {
			log.Fatal(err)
		}

		//log.Print(tickers)

		usdtBTCValue := tickers["USDT_BTC"].Last

		balances, err := client.GetBalances()
		if err != nil {
			log.Fatal(err)
		}

		for name, balance := range balances {
			available, _ := strconv.ParseFloat(balance.Available, 64)
			btcValue, _ := strconv.ParseFloat(balance.BtcValue, 64)

			if available == 0 {
				continue
			}

			log.Print(name, available*btcValue*usdtBTCValue)
		}

		log.Print(usdtBTCValue)

		time.Sleep(time.Second*10)
	}


}
