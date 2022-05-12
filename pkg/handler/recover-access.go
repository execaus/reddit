package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"reddit/models"
)

func (h *Handler) GenerateRecoverAccessLink(c *gin.Context) {
	var input models.InputRecoverAccessLink
	if err := c.BindJSON(&input); err != nil {
		sendBadRequestError(c, err)
		return
	}

	if err := h.services.RecoverAccess.GenerateLink(&input); err != nil {
		sendInternalServerError(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) RegisterNewPassword(c *gin.Context) {
	var input models.InputRecoverAccessRegister
	if err := c.BindJSON(&input); err != nil {
		sendBadRequestError(c, err)
		return
	}

	message, err := h.services.RecoverAccess.RegisterNewPassword(&input)
	if err != nil {
		sendBadRequestWithMessage(c, errors.New(message))
		return
	}

	c.Status(http.StatusOK)
}
