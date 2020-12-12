package main

import (
	"flag"

	"github.com/joho/godotenv"
)

var ticker string
var period1 string
var period2 string
var interval string

func init () {
  flag.StringVar(&ticker, "ticker", "AMZN", "Symbol")
  flag.StringVar(&period1, "period1", "2001-12-01T00:00:00+08:00", "Period1")
  flag.StringVar(&period2, "period2", "2020-12-31T23:59:59+08:00", "Period2")
  flag.StringVar(&interval, "interval", "1d", "Interval, only support '1d'")
}

func main() {
  flag.Parse()

  godotenv.Load()

  migrateDB()

  Download(ticker, period1, period2, interval)
}
