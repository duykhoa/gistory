package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
  godotenv.Load()

  equitySymbol := os.Args[1]

  Download(
    equitySymbol,
    "2019-12-19T00:00:00+08:00",
    "2019-12-20T23:00:00+08:00",
    "1d",
  )
}
