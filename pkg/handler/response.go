package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type HttpError struct {
	Message string `json:"message"`
}

func sendBadRequestError(c *gin.Context, err error) {
	sendError(c, http.StatusBadRequest, err)
}

func sendInternalServerError(c *gin.Context, err error) {
	sendError(c, http.StatusInternalServerError, err)
}

func sendBadRequestWithMessage(c *gin.Context, err error) {
	sendErrorWithMessage(c, http.StatusBadRequest, err)
}

func sendError(c *gin.Context, status int, err error) {
	c.AbortWithStatus(status)
	log.Println(err.Error())
}

func sendErrorWithMessage(c *gin.Context, status int, err error) {
	errorJson := HttpError{
		Message: err.Error(),
	}
	c.AbortWithStatusJSON(status, errorJson)
}
