package handler

import (
	"github.com/elizanotbeth/test_const/pckg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/", h.createPayment)
		api.GET("/byuserid", h.getPaymentByIdUser)
		api.GET("/byemail", h.getPaymentByEmail)
		api.PUT("/change", h.updatePayment)

		//api.PATCH("/deletepay/:id", h.updatePaymentById)
	}
	return router
}
