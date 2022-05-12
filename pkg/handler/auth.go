package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reddit/models"
)

func (h *Handler) SignIn(c *gin.Context) {
	var input models.InputSignIn

	if err := c.BindJSON(&input); err != nil {
		sendBadRequestError(c, err)
		return
	}

	if err := input.Validate(); err != nil {
		sendBadRequestWithMessage(c, err)
		return
	}

	output, err := h.services.Auth.SignIn(&input)
	if err != nil {
		sendInternalServerError(c, err)
		return
	}

	output.Session, err = h.services.Session.Generate(output.Account.Login)
	if err != nil {
		sendInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) SignUp(c *gin.Context) {
	var input models.InputSignUp
	var output models.OutputSignUp

	if err := c.BindJSON(&input); err != nil {
		sendBadRequestError(c, err)
		return
	}

	if err := input.IsValid(); err != nil {
		sendBadRequestWithMessage(c, err)
		return
	}

	if err := h.services.Auth.SignUp(&input); err != nil {
		sendInternalServerError(c, err)
		return
	}

	hash, err := h.services.Session.Generate(input.Login)
	if err != nil {
		sendInternalServerError(c, err)
		return
	}

	output.Session = hash

	c.JSON(http.StatusOK, output)
}
