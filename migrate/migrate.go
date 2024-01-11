package main

import (
	"server-test/web-service-gin/initializers"
	"server-test/web-service-gin/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
