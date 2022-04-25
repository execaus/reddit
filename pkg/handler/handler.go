package handler

import (
	"github.com/gin-gonic/gin"
	"reddit/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h Handler) GetRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		posts := api.Group("/posts")
		{
			posts.GET("/:item_id", h.GetPostById)
			posts.GET("/:page/:limit", h.GetList)
			posts.POST("", h.Create)
			posts.PUT("/:item_id", h.Update)
			posts.DELETE("/:item_id", h.Delete)
		}
	}

	return router
}
