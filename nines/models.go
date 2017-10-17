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
	//stmt, err := config.DB.Prepare("INSERT INTO nines (login, game_name, bet_banker, bet_player, bet_tie, total_payout, game_number, game_time, dealer_cards, player_cards, banker_n9, player_n9) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer stmt.Close()

	_, err := stmt.Exec(n.Login, n.GameName, n.BetBanker, n.BetPlayer, n.BetTie, n.TotalPayout, n.GameNumber, n.GameTime, n.DealerCards, n.PlayerCards, n.BankerNine, n.PlayerNine)
	if err != nil {
		log.Fatal(err.Error())
	}
}
