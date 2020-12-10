package main

import (
	"log"
	"strconv"
	"time"
)

func parseInt(num string) int64 {
  numi, err := strconv.Atoi(num)

  if (err != nil) {
    log.Fatal(err)
  }

  return int64(numi)
}

func parseDate(date string) time.Time {
  d, err := time.Parse("2006-01-02", date)

  if (err != nil) {
    log.Fatal(err)
  }

  return d
}


func parseTime(date string) time.Time {
  d, err := time.Parse(time.RFC3339, date)

  if (err != nil) {
    log.Fatal(err)
  }

  return d
}

func parseTimeBasic(date string) time.Time {
  d, err := time.Parse("2006-01-02 03:04:05", date)

  if (err != nil) {
    log.Fatal(err)
  }

  return d
}

func parseFloat(number string) float64 {
  f, err := strconv.ParseFloat(number, 0)

  if (err != nil) {
    log.Fatal(err)
  }

  return f
}
