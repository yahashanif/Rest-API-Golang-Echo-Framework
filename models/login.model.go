package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"rest-api-go-echo/db"
	"rest-api-go-echo/helpers"
	"strconv"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func RegisterUser(username, password string) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "INSERT INTO `users` (`username`, `password`) VALUES (?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}
	hash, _ := helpers.HashPassword(password)
	result, err := stmt.Exec(username, hash)

	if err != nil {
		return res, err

	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses Created User"
	res.Data = map[string]int64{
		"LastInsertID": lastInsertID,
	}

	return res, nil

}

func CheckLogin(username, password string) (Response, error) {
	var res Response
	var obj User

	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {

		fmt.Println("Username Not Found")
		return res, err

	}

	if err != nil {
		fmt.Print("Query Error")
		return res, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {

		fmt.Println("Hash and password doesn't match")
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Sukses Login"
	res.Data = map[string]string{
		"id":       strconv.Itoa(obj.Id),
		"username": obj.Username,
	}

	return res, nil
}
