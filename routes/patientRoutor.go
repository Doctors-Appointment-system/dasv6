package routes

import (
	controller "Doctor-Appointment-Project/controllers"
	middleware "Doctor-Appointment-Project/middleware"

	"github.com/gin-gonic/gin"
)

func PatientRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/patient", controller.Addpatient())
	incomingRoutes.GET("/get_my_details", controller.Get_my_details())
	incomingRoutes.DELETE("/patient", controller.DeletePatient())

	// Doctor
	incomingRoutes.GET("/showall/doctors", controller.Get_docter())
	incomingRoutes.GET("/get_doctor_by_city", controller.GetDoctorByLocation())
	incomingRoutes.POST("/bookappointment", controller.BookingAppointment())
	incomingRoutes.DELETE("/cancelAppointment", controller.Cancel_appointment())
	incomingRoutes.POST("/doctor_feedback", controller.Doctor_feedback())

	// Lab

	incomingRoutes.POST("/book_test", controller.Book_test())
	incomingRoutes.GET("/get_lab_by_location", controller.GetLabByLocation())
	incomingRoutes.POST("/lab_feedback", controller.Lab_feedback())

}
