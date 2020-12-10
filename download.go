package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"os"
)

type HistoricalData struct {
  Date string
  Open big.Float
  High big.Float
  Low big.Float
  Close big.Float
  AdjustClose big.Float
  Volume big.Int
}

func Download(equitySymbol string, period1 string, period2 string, interval string) {
  endpoint := os.Getenv("YAHOO_HISTORICAL_DATA_ENDPOINT")

  u, _ := url.Parse(endpoint)
  u.Path += equitySymbol
  q := u.Query()

  q.Set("period1", fmt.Sprintf("%d", parseTime(period1).Unix()))
  q.Set("period2", fmt.Sprintf("%d", parseTime(period2).Unix()))

  q.Set("interval", interval)
  q.Set("events", "history")
  q.Set("includeAdjustedClose", "true")

  u.RawQuery = q.Encode()
  fmt.Println("About to call api to", u.String())

  resp, err := http.Get(u.String())

  if err != nil {
    log.Fatal(err)
  }

  r := csv.NewReader(resp.Body)

  if resp.StatusCode != http.StatusOK {
    fmt.Println("Error response", resp.StatusCode)
    log.Fatal(resp.StatusCode)
  } else {
    processCSV(r, buildRecord)
  }
}

func buildRecord(record []string) *HistoricalData {
  history := &HistoricalData{
    Date:          record[0],
    Open:          parseFloat(record[1]),
    High:          parseFloat(record[2]),
    Low:           parseFloat(record[3]),
    Close:         parseFloat(record[4]),
    AdjustClose:   parseFloat(record[5]),
    Volume:        parseInt(record[6]),
  }

  fmt.Println(history.Date, &history.Open, &history.Close)

  return history
}

func processCSV(csvReader *csv.Reader, fun func(r []string) *HistoricalData) {
  // exclude header
  _, _ = csvReader.Read()

  for {
    record, err := csvReader.Read()

    if err == io.EOF {
      break
    }

    if err != nil {
      log.Fatal(err)
    } else {
      fun(record)
    }
  }
}
