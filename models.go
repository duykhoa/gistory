package main

import (
	"time"

	"gorm.io/gorm"
)

type HistoricalData struct {
    gorm.Model

    Date time.Time
    Symbol string
    Open float64
    High float64
    Low float64
    Close float64
    AdjustClose float64
    Volume int64
}

type Ticker struct {
  gorm.Model

  Symbol string
}

