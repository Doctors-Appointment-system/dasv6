package routes

import (
	controller "Doctor-Appointment-Project/controllers"
	middleware "Doctor-Appointment-Project/middleware"

	"github.com/gin-gonic/gin"
)

func NurseRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/nurse", controller.Get_nurse())
	incomingRoutes.GET("/nurse", controller.Get_nurse_by_city())
	incomingRoutes.GET("/nurse", controller.Get_nurse_by_specialisation())
	incomingRoutes.GET("/nurse", controller.Get_nurse_by_location())
	incomingRoutes.GET("/nurse", controller.Get_nurse_profile())
	incomingRoutes.GET("/nurse", controller.Nurse_checking_feedback())
	incomingRoutes.GET("/nurse", controller.Check_nurse_appointment())

	incomingRoutes.POST("/nurse", controller.Add_nurse())

	incomingRoutes.PUT("/nurse", controller.Update_nurse())
	incomingRoutes.DELETE("/nurse", controller.Delete_nurse())
}
