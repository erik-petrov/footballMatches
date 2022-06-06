package storage

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	UsersStorage interface {
		GetUsers() ([]User, error)
		CreateUser(user User) error
		GetUserByID(id string) (User, error)
		GetUserByEmail(mail string) (User, error)
		GetAllEmails() ([]string, error)
		ChangeUserPassword(id string, password string) error
		ChangeUser(user User, id string) error
	}

	Users struct {
		*sqlx.DB
	}

	User struct {
		ID        int       `db:"id" json:"id"`
		Name      string    `db:"name" json:"name"`
		Email     string    `db:"email" json:"email"`
		Password  string    `db:"password" json:"password"`
		CreatedAt time.Time `db:"created_at" json:"created_at"`
		UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
		Birthday  time.Time `db:"birthday" json:"birthday"`
		Gender    string    `db:"gender" json:"gender"`
	}
)

func (db Users) GetUsers() (users []User, err error) {
	const q = `select * from users`
	return users, db.Select(&users, q)
}

func (db Users) CreateUser(user User) error {
	const q = `insert into users(name, email, created_at, updated_at, birthday, gender, password) values ($1, $2, $3, $4, $5, $6, $7)`
	_, err := db.Exec(q, user.Name, user.Email, time.Now(), time.Now(), user.Birthday, user.Gender, user.Password)
	return err
}

func (db Users) GetUserByID(id string) (users User, err error) {
	idInt, _ := strconv.Atoi(id)
	const q = `select * from users where id = $1`
	return users, db.Get(&users, q, idInt)
}

func (db Users) GetUserByEmail(mail string) (user User, err error) {
	const q = `select id, name, password, email, created_at, birthday, gender from users where email = $1`
	return user, db.Get(&user, q, mail)
}

func (db Users) GetAllEmails() (mails []string, err error) {
	const q = `select email from users;`
	return mails, db.Select(&mails, q)
}

func (db Users) ChangeUserPassword(id string, password string) error {
	idInt, _ := strconv.Atoi(id)
	const q = `update users set password = $1 where id = $2`
	_, err := db.Exec(q, password, idInt)
	return err
}

func (db Users) ChangeUser(user User, id string) error {
	idInt, _ := strconv.Atoi(id)
	const q = `update users set name = $1, email = $2, birthday = $3, gender = $4, updated_at = $5 where id = $6`
	output, err := db.Exec(q, user.Name, user.Email, user.Birthday, user.Gender, time.Now(), idInt)
	fmt.Println(output)
	return err
}
