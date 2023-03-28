package routes

import (
	controller "Doctor-Appointment-Project/controllers"
	middleware "Doctor-Appointment-Project/middleware"

	"github.com/gin-gonic/gin"
)

func LabtRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/lab", controller.Add_lab_details())
	incomingRoutes.PUT("/lab", controller.Update_lab())
	incomingRoutes.GET("/lab", controller.Get_Lab_Profile())

}
