package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kalberfon/buscaCep-go-react/api/usecases"
)

func GetSearchCep(c *gin.Context) {
	cep := []string{c.Param("cep")}

	respostaCep := usecases.BuscaCepCase(cep)

	c.JSON(http.StatusOK, respostaCep)
}
