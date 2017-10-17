package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/avecost/ibn/nines"
	"strconv"
	"github.com/avecost/ibn/config"
)

func processCSV(rc io.Reader) (ch chan []string) {
	ch = make(chan []string, 10)
	go func() {
		r := csv.NewReader(rc)
		if _, err := r.Read(); err != nil { // read header
			log.Fatal(err)
		}
		defer close(ch)

		for {
			rec, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			ch <- rec
		}
	}()
	return
}

func checkBool(s string) int {
	if s == "Y" || s == "y" {
		return 1
	}
	return 0
}

func ImportCSV(fCSV string) (int, error) {
	f, err := os.Open(fCSV)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	i := 0
	// prepare stmt
	stmt, err := config.DBPrepareStatement("INSERT INTO nines (login, game_name, bet_banker, bet_player, bet_tie, total_payout, game_number, game_time, dealer_cards, player_cards, banker_n9, player_n9) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return 0, err
	}
	// process CSV
	for row := range processCSV(f) {
		i++
		bb, _ := strconv.ParseFloat(row[2], 64)
		bp, _ := strconv.ParseFloat(row[3], 64)
		bt, _ := strconv.ParseFloat(row[4], 64)
		tp, _ := strconv.ParseFloat(row[5], 64)
		b9 := checkBool(row[12])
		p9 := checkBool(row[13])
		nines.InsertNine(stmt, nines.Nine{0, row[0], row[1], bb, bp, bt, tp, row[8], row[9], row[10], row[11], b9, p9})
	}
	// close stmt
	config.DBCloseStatement(stmt)
	return i, nil
}
