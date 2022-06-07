package main

import (
	"encoding/json"
	"erik-petrov/footballWebsite/storage"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func dbMiddleware(db *storage.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("__db", db)
			return next(c)
		}
	}
}

func getMatches(c echo.Context) error {
	db := c.Get("__db").(*storage.DB)
	matches, err := db.Matches.GetMatches()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, matches)
}

func getUsers(c echo.Context) error {
	db := c.Get("__db").(*storage.DB)
	emails, err := db.Users.GetAllEmails()
	m := make(map[string][]string)
	m["emails"] = emails
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, m)
}

func login(c echo.Context) error {
	db := c.Get("__db").(*storage.DB)
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	fmt.Println(jsonMap["email"].(string))
	fmt.Println(jsonMap["password"].(string))
	if err != nil {
		fmt.Println(err)
		return err
	}
	user, err := db.Users.GetUserByEmail(jsonMap["email"].(string))
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = checkPassword(user.Password, jsonMap["password"].(string))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("success")
	user.Password = "" //на всякий случай, я не знаю сильно ли плохо давать хэш юзеру
	return c.JSON(http.StatusOK, user)
}

func signup(c echo.Context) error {
	db := c.Get("__db").(*storage.DB)
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		fmt.Println(err)
		return err
	}
	emails, err := db.Users.GetAllEmails()
	if stringInSlice(jsonMap["email"].(string), emails) {
		err := errors.New("duplicate email")
		return err
	}
	u := storage.User{}
	u.Password = hashPassword(jsonMap["password"].(string))
	u.Birthday, _ = time.Parse("2006-02-05", jsonMap["date"].(string))
	u.Name = jsonMap["name"].(string) + ":" + jsonMap["sname"].(string)
	u.Gender = jsonMap["gender"].(string)
	u.Email = jsonMap["email"].(string)
	err = db.Users.CreateUser(u)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, "lol")
}

func hashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashed)
}

func checkPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func changePassword(c echo.Context) error {
	db := c.Get("__db").(*storage.DB)
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		fmt.Println(err)
		return err
	}
	password := hashPassword(jsonMap["pswd"].(string))
	id := jsonMap["id"].(string)
	err = db.Users.ChangeUserPassword(id, password)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, "lol")
}

func editUser(c echo.Context) error {
	db := c.Get("__db").(*storage.DB)
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		fmt.Println(err)
		return err
	}
	u := storage.User{}
	u.Birthday, _ = time.Parse("2006-01-02", jsonMap["date"].(string))
	fmt.Println(u.Birthday)
	u.Name = jsonMap["name"].(string) + ":" + jsonMap["sname"].(string)
	u.Gender = jsonMap["gender"].(string)
	u.Email = jsonMap["email"].(string)
	id := jsonMap["id"].(string)
	err = db.Users.ChangeUser(u, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func main() {
	db, err := storage.Open("postgres://postgres:123123@localhost:5432/football")
	if err != nil {
		panic(err)
	}
	e := echo.New()
	e.Use(dbMiddleware(db))
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/matches", getMatches)
	e.GET("/users", getUsers)
	e.POST("/login", login)
	e.POST("/reg", signup)
	e.POST("/pswdChange", changePassword)
	e.POST("/edit", editUser)
	e.Logger.Fatal(e.Start(":9000"))
}
