package main

import (
	"GoApp/initializers"
	"GoApp/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}
func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
