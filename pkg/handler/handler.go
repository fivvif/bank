<<<<<<< HEAD
package handler

import (
	"bank/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		transactions := api.Group("/transactions")
		{
			transactions.GET("/", h.giveMoney)
			transactions.POST("/", h.takeMoney)
		}
	}

	return router
}
=======
package handler

import (
	"bank/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		transactions := api.Group("/transactions")
		{
			transactions.GET("/give", h.giveMoney)
			transactions.POST("/", h.takeMoney)
		}
	}

	return router
}
>>>>>>> f48611f (first commit)