package routes

import (
	controller "Doctor-Appointment-Project/controllers"
	middleware "Doctor-Appointment-Project/middleware"

	"github.com/gin-gonic/gin"
)

func DoctorRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())                // this protect our routes
	incomingRoutes.POST("/doctor", controller.Add_docter())      //done
	incomingRoutes.PUT("/doctor", controller.Update_docter())    //done
	incomingRoutes.DELETE("/doctor", controller.Delete_docter()) //done
	incomingRoutes.GET("/doctor/MyAppointment", controller.CheckMyAppointment())
}
