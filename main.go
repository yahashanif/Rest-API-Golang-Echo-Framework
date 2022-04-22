package main

import (
	"rest-api-go-echo/db"
	"rest-api-go-echo/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":9000"))
}
