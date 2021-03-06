package main

import (
	"os"
	"github.com/avecost/ibn/config"
	"github.com/avecost/ibn/csv"
	"log"
)

func main() {
	fToUpload := os.Args[1]
	if _, err := os.Stat(fToUpload); os.IsNotExist(err) {
		panic("File does not exist")
	}

	dbSourceName := os.Args[2]
	config.InitDB(dbSourceName)

	count, err := csv.ImportCSV(fToUpload)
	if err != nil {
		log.Panic("Error importing CSV")
	}
	log.Printf("Process %d records\n", count)
}
