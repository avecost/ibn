package nines

import (
	"log"

	"database/sql"
)

type Nine struct {
	Id          int
	Login       string
	GameName    string
	BetBanker   float64
	BetPlayer   float64
	BetTie      float64
	TotalPayout float64
	GameNumber  string
	GameTime    string
	DealerCards string
	PlayerCards string
	BankerNine  int
	PlayerNine  int
}

func InsertNine(stmt *sql.Stmt, n Nine) {
	_, err := stmt.Exec(n.Login, n.GameName, n.BetBanker, n.BetPlayer, n.BetTie, n.TotalPayout, n.GameNumber, n.GameTime, n.DealerCards, n.PlayerCards, n.BankerNine, n.PlayerNine)
	if err != nil {
		log.Fatal(err.Error())
	}
}
