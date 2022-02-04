package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kalberfon/buscaCep-go-react/api/controllers"
	"github.com/kalberfon/buscaCep-go-react/api/middlewares"
)

func main() {
	godotenv.Load(".env")
	router := gin.Default()

	router.Use(middlewares.AllowCorsAllOrigin)
	router.GET("/:cep", controllers.GetSearchCep)

	router.Run(":8000")
}
