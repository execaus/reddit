package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"reddit/models"
	"strconv"
)

func (h *Handler) GetPostById(c *gin.Context) {
	id := c.Param("item_id")
	if id == "" {
		sendBadRequestError(c, errors.New("invalid item id"))
		return
	}

	post, err := h.services.Post.GetById(id)
	if err != nil {
		sendInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) GetList(c *gin.Context) {
	var uriParams models.UriGetPostList
	if err := c.ShouldBindUri(&uriParams); err != nil {
		sendBadRequestError(c, err)
		return
	}
	intPage, err := strconv.Atoi(uriParams.Page)
	if err != nil {
		sendInternalServerError(c, err)
		return
	}
	intLimit, err := strconv.Atoi(uriParams.Limit)
	if err != nil {
		sendInternalServerError(c, err)
		return
	}

	if intPage <= 0 || intLimit <= 0 {
		sendBadRequestError(c, err)
		return
	}

	output, err := h.services.Post.GetList(intPage, intLimit)
	if err != nil {
		sendInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) Create(c *gin.Context) {
	var input models.InputPost

	if err := c.BindJSON(&input); err != nil {
		sendBadRequestError(c, err)
		return
	}

	output, err := h.services.Post.Create(&input)
	if err != nil {
		sendInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *Handler) Update(c *gin.Context) {
	var input models.InputUpdatePost

	id := c.Param("item_id")
	if id == "" {
		sendBadRequestError(c, errors.New("invalid item id"))
		return
	}
	input.Id = id

	if err := c.BindJSON(&input); err != nil {
		sendBadRequestError(c, err)
		return
	}

	if err := h.services.Post.Update(&input); err != nil {
		sendInternalServerError(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("item_id")
	if id == "" {
		sendBadRequestError(c, errors.New("invalid item id"))
		return
	}

	if err := h.services.Post.Delete(id); err != nil {
		sendInternalServerError(c, err)
		return
	}

	c.Status(http.StatusOK)
}
