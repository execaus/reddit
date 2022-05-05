package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"reddit/models"
)

const headerAuthorization = "Authorization"

func (h *Handler) apiMiddleware(c *gin.Context) {
	h.setUserIdentity(c)
}

func (h *Handler) setUserIdentity(c *gin.Context) {
	session := c.Request.Header.Get(headerAuthorization)
	if session == "" {
		sendUnauthorizedError(c, errors.New("header authorization is empty"))
		return
	}
	account, err := h.services.Session.GetAccount(session)
	if err != nil {
		sendUnauthorizedError(c, err)
		return
	}
	c.Set(gin.AuthUserKey, account)
}

func getAuthAccount(c *gin.Context) *models.Account {
	account, ok := c.Get(gin.AuthUserKey)
	if !ok {
		sendInternalServerErrorGetAuthUser(c)
		return nil
	}
	return account.(*models.Account)
}
