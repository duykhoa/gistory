package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func Download(equitySymbol string, period1 string, period2 string, interval string) {
  endpoint := os.Getenv("YAHOO_HISTORICAL_DATA_ENDPOINT")

  if interval != "1d" {
    log.Fatalf("ERROR! Sorry, the interval \"%s\" isn't yet supported", interval)
  }

  u, _ := url.Parse(endpoint)
  u.Path += equitySymbol
  q := u.Query()

  q.Set("period1", fmt.Sprintf("%d", parseTime(period1).Unix()))
  q.Set("period2", fmt.Sprintf("%d", parseTime(period2).Unix()))

  q.Set("interval", interval)
  q.Set("events", "history")
  q.Set("includeAdjustedClose", "true")

  u.RawQuery = q.Encode()
  log.Println("--- INFO: Calling api:", u.String())

  resp, err := http.Get(u.String())

  if err != nil {
    log.Fatal(err)
  }

  r := csv.NewReader(resp.Body)

  if resp.StatusCode != http.StatusOK {
    log.Println("--- ERROR: Error response", resp.StatusCode)
    log.Fatal(resp.StatusCode)
  } else {
    priceHistories := processCSV(r, equitySymbol, buildRecord)

    gormDbResult := Db.Create(&priceHistories)

    if (gormDbResult.Error != nil) {
      log.Fatal("--- ERROR: Persist HistoricalData failed", gormDbResult.Error)
    }

    log.Printf("--- INFO: Successfully process the history data. Symbol: \"%s\", Period1: \"%s\", Period2: \"%s\"\n", equitySymbol, period1, period2)
  }
}

func buildRecord(record []string, equitySymbol string) *HistoricalData {
  history := &HistoricalData{
    Symbol:        equitySymbol,
    Date:          parseDate(record[0]),
    Open:          parseFloat(record[1]),
    High:          parseFloat(record[2]),
    Low:           parseFloat(record[3]),
    Close:         parseFloat(record[4]),
    AdjustClose:   parseFloat(record[5]),
    Volume:        parseInt(record[6]),
  }

  return history
}

func processCSV(csvReader *csv.Reader, equitySymbol string, fun func(r []string, equitySymbol string) *HistoricalData) []*HistoricalData {
  // exclude header
  _, _ = csvReader.Read()

  result := make([]*HistoricalData, 0)

  for {
    record, err := csvReader.Read()

    if err == io.EOF {
      break
    }

    if err != nil {
      log.Fatal(err)
    } else {
      result = append(result, fun(record, equitySymbol))
    }
  }

  return result
}
