package main

import (
	"example.com/movie-app/db"
	"nikita sehgal/movie-app/router"
)

func main() {
	db.InitPostgresDB()
	router.InitRouter().Run()
}
