package storage

import (
	"github.com/jmoiron/sqlx"
)

type (
	MatchesStorage interface {
		GetMatches() ([]Match, error)
		CreateMatch(date string, length string, team1 int, team2 int, score string, inProgress bool) error
		GetMatchByTeam(team string) (Match, error)
	}

	Matches struct {
		*sqlx.DB
	}

	Match struct {
		ID         int     `db:"id" json:"id"`
		Date       string  `db:"date" json:"date"`
		Length     *string `db:"length" json:"length"`
		Team1      string  `db:"team1" json:"team1"`
		Team2      string  `db:"team2" json:"team2"`
		Score      *string `db:"score" json:"score"`
		InProgress bool    `db:"inprogress" json:"inprogress"`
	}
)

func (db Matches) GetMatches() (matches []Match, err error) {
	const q = `select matches.id, date, length, t1.name as team1, t2.name as team2, score, inprogress from matches
join teams t1 on matches.team1 = t1.id
join teams t2 on matches.team2 = t2.id
ORDER by date asc;`
	return matches, db.Select(&matches, q)
}

func (db Matches) CreateMatch(date string, length string, team1 int, team2 int, score string, inProgress bool) error {
	const q = `insert into matches(date, length, team1, team2, score, inprogress) values ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(q, date, length, team1, team2, score, inProgress)
	return err
}

func (db Matches) GetMatchByTeam(team string) (matches Match, err error) {
	const q = `select * from matches where team1 = $1 OR team2 = $1`
	return matches, db.Get(&matches, q, team)
}
