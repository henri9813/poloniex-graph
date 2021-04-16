package main

import (
	"fmt"

	"log"
	"os"
	"strconv"
	"strings"
	"time"

	influxdb "github.com/influxdata/influxdb1-client/v2"
	"github.com/jyap808/go-poloniex"
)

func main()  {
	client := poloniex.New(os.Getenv("POLONIEX_API_KEY"), os.Getenv("POLONIEX_API_SECRET"))

	for {
		total := 0.0

		tickers, err := client.GetTickers()
		if err != nil {
			log.Printf("An error occured, retrying in 10s: '%v'", err)
			time.Sleep(10*time.Second)
			continue
		}

		balances, err := client.GetBalances()
		if err != nil {
			log.Fatal(err)
		}

		batchPoints, err := influxdb.NewBatchPoints(influxdb.BatchPointsConfig{
			Database: influxDatabase,
		})

		for name, balance := range balances {
			available, _ := strconv.ParseFloat(balance.Available, 64)

			if available == 0 {
				continue
			}

			var ticker poloniex.Ticker
			if ticker, err = getTickerForBalance(tickers, name); err != nil {
				log.Print(err)
			}

			amount := ticker.Last*available

			cryptoPoint, _ := influxdb.NewPoint(
				"crypto",
				map[string]string{
					"name": name,
				},
				map[string]interface{}{
					"total": amount,
				},
				time.Now(),
			)
			batchPoints.AddPoint(cryptoPoint)

			total += amount
			log.Print(ticker)
			log.Printf("You have : %f %s because %f * %f ", ticker.Last*available, name, ticker.Last, available)
		}

		cryptoPoint, _ := influxdb.NewPoint(
			"crypto",
			map[string]string{
				"name": "total",
			},
			map[string]interface{}{
				"total": total,
			},
			time.Now(),
		)
		batchPoints.AddPoint(cryptoPoint)

		err = influxClient.Write(batchPoints)
		if err != nil {
			log.Fatalf("[InfluxDB] Impossible to write points: %v", err)
		}

		log.Printf("Total: %f", total)
		time.Sleep(time.Second*10)
	}
}

func getTickerForBalance(tickers map[string]poloniex.Ticker, balanceName string) (poloniex.Ticker, error){
	for name, ticker := range tickers {
		if strings.HasSuffix(name, balanceName) {
			if strings.Contains(name, "USDT") {
				//log.Printf("Found %s for %s", name, balanceName)
				return ticker, nil
			}
		}
	}

	return poloniex.Ticker{Last: 0}, fmt.Errorf("no ticker found for %s in USDT", balanceName)
}
