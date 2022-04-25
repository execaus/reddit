package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reddit/models"
	"strconv"
)

func (h *Handler) GetPostById(c *gin.Context) {
	id := c.Param("item_id")
	if id == "" {
		sendBadRequestError(c)
		return
	}

	post, err := h.services.Post.GetById(id)
	if err != nil {
		sendInternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) GetList(c *gin.Context) {
	intPage, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		sendBadRequestError(c)
	}
	intLimit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		sendBadRequestError(c)
	}

	if intPage <= 0 || intLimit <= 0 {
		sendBadRequestError(c)
	}

	output, err := h.services.Post.GetList(intPage, intLimit)
	if err != nil {
		sendInternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) Create(c *gin.Context) {
	var input models.InputPost

	if err := c.BindJSON(&input); err != nil {
		sendBadRequestError(c)
		return
	}

	output, err := h.services.Post.Create(&input)
	if err != nil {
		sendInternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) Update(c *gin.Context) {
	var input models.InputUpdatePost

	id := c.Param("item_id")
	if id == "" {
		sendBadRequestError(c)
		return
	}
	input.Id = id

	if err := c.BindJSON(&input); err != nil {
		sendBadRequestError(c)
		return
	}

	if err := h.services.Post.Update(&input); err != nil {
		sendInternalServerError(c)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("item_id")
	if id == "" {
		sendBadRequestError(c)
		return
	}

	if err := h.services.Post.Delete(id); err != nil {
		sendInternalServerError(c)
		return
	}

	c.Status(http.StatusOK)
}
