package main

import "GoApp/initializers"

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}
func main() {
	initializers.DB.AutoMigrate("&models.Post{}")
}
