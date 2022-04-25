package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpError struct {
	Message string `json:"message"`
}

func sendBadRequestError(c *gin.Context) {
	c.AbortWithStatus(http.StatusBadRequest)
}

func sendInternalServerError(c *gin.Context) {
	c.AbortWithStatus(http.StatusInternalServerError)
}
