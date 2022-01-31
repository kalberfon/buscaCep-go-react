package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kalberfon/buscaCep-go-react/api/controllers"
)

func main() {
	godotenv.Load(".env")
	r := gin.Default()
	r.GET("/:cep", controllers.GetSearchCep)

	r.Run()
}
