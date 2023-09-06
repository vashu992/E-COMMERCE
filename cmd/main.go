package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vashu992/E-Commerce/controllers"
	"github.com/vashu992/E-Commerce/database"
	"github.com/vashu992/E-Commerce/middleware"
	"github.com/vashu992/E-Commerce/routes"
)

func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Product"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}