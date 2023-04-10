package main

import (
	"github.com/baguspanji/go-crud/initializers"
	"github.com/baguspanji/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
