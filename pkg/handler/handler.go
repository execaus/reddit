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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}

	recoverAccess := router.Group("/recover-access")
	{
		recoverAccess.POST("/", h.GenerateRecoverAccessLink)
		recoverAccess.PUT("/", h.RegisterNewPassword)
	}

	api := router.Group("/api", h.apiMiddleware)
	{
		posts := api.Group("/posts")
		{
			posts.GET("/:item_id", h.GetPostById)
			posts.GET("/list/:page/:limit", h.GetList)
			posts.POST("", h.Create)
			posts.PUT("/:item_id", h.Update)
			posts.DELETE("/:item_id", h.Delete)
		}
	}

	return router
}
