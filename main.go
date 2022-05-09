package main

import (
	"crud/models"
	"crud/routes"
)

func main() {

	db := models.SetUpDB()
	db.AutoMigrate(&models.Movie{})
	r := routes.SetupRoutes(db)
	r.Run()
}
