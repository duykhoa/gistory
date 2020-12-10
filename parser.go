package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"
)

func parseInt(num string) big.Int {
  numi, err := strconv.Atoi(num)

  if (err != nil) {
    log.Fatal(err)
  }

  return *big.NewInt(int64(numi))
}

func parseTime(date string) time.Time {
  d, err := time.Parse(time.RFC3339, date)

  if (err != nil) {
    log.Fatal(err)
  }

  return d
}

func parseTimeBasic(date string) time.Time {
  fmt.Println(date)
  d, err := time.Parse("2006-01-02 03:04:05", date)

  if (err != nil) {
    log.Fatal(err)
  }

  return d
}

func parseFloat(number string) big.Float {
  f, err := strconv.ParseFloat(number, 0)

  if (err != nil) {
    log.Fatal(err)
  }

  return *big.NewFloat(f)
}
