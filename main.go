package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/turuu/ecommerce/controllers"
	"github.com/turuu/ecommerce/database"
	"github.com/turuu/ecommerce/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.Userdata(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.User(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.instantbuy())

	log.Fatal(router.Run(":" + port))
}
