package handler

import "github.com/gin-gonic/gin"

func (h *Handler) apiMiddleware(c *gin.Context) {
	h.setUserIdentity(c)
}

func (h *Handler) setUserIdentity(c *gin.Context) {
	session := "asdasd"
	account, err := h.services.Session.GetAccount(session)
	if err != nil {
		sendUnauthorizedError(c, err)
	}
	c.Set("account", account)
}
