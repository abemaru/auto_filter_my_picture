package main

import (
	"log"
	"os"
)

func main() {
	if err := os.MkdirAll("case/case01", os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
