package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reddit/models"
)

func (h *Handler) SignIn(c *gin.Context) {

}

func (h *Handler) SignUp(c *gin.Context) {
	var input models.InputSignUp

	if err := c.BindJSON(&input); err != nil {
		sendBadRequestError(c, err)
		return
	}

	if err := input.IsValid(); err != nil {
		sendBadRequestWithMessage(c, err)
		return
	}

	output, err := h.services.Auth.SignUp(&input)
	if err != nil {
		sendInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}
