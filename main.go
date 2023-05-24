package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/turuu/ecommerce/controllers"
	"github.com/turuu/ecommerce/database"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.Userdata(database.Client, "Products"))

	router := gin.New()
	router.Use(gin.Logger)
}
