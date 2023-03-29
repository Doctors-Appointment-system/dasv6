package routes

import (
	controller "Doctor-Appointment-Project/controllers"
	middleware "Doctor-Appointment-Project/middleware"

	"github.com/gin-gonic/gin"
)

func PatientRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())

	// Patient

	incomingRoutes.POST("/patient", controller.Add_patient())
	incomingRoutes.GET("/get_patient_details", controller.Get_patient_details())
	incomingRoutes.DELETE("/patient", controller.DeletePatient())

	//Doctor

	incomingRoutes.POST("/doctor_feedback", controller.Doctor_feedback())
	incomingRoutes.GET("/showall/doctors", controller.Get_docter())
	incomingRoutes.GET("/get_doctor_by_city", controller.GetDoctorByLocation())
	incomingRoutes.POST("/bookappointment", controller.Book_doctor_appointment())
	incomingRoutes.DELETE("/cancelAppointment", controller.Cancel_doctor_appointment())

	// Nurse
	incomingRoutes.POST("/nurse/feedback", controller.Nurse_feedback())
	incomingRoutes.POST("/nurse/book", controller.Book_nurse())
	incomingRoutes.DELETE("/nurse", controller.Cancel_nurse_appointment())
}
