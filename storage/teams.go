package storage

import (
	"github.com/jmoiron/sqlx"
)

type (
	TeamsStorage interface {
		GetTeams() ([]Team, error)
		CreateTeam(name string, coach string) error
		GetTeamByName(name string) (Team, error)
	}

	Teams struct {
		*sqlx.DB
	}

	Team struct {
		ID    int    `db:"id" json:"id"`
		Name  string `db:"name" json:"name"`
		Coach string `db:"coach" json:"coach"`
	}
)

func (db Teams) GetTeams() (teams []Team, err error) {
	const q = `select * from teams`
	return teams, db.Select(&teams, q)
}

func (db Teams) CreateTeam(name string, coach string) error {
	const q = `insert into teams(name, coach) values ($1, $2)`
	_, err := db.Exec(q, name, coach)
	return err
}

func (db Teams) GetTeamByName(name string) (teams Team, err error) {
	const q = `select * from teams where name = $1`
	return teams, db.Get(&teams, q, name)
}
