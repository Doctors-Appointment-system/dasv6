package models

type Lab struct {
	Labid                      int
	Lab_Name                   string
	Lab_Operator_Name          string
	Phone                      string
	Address                    string
	City                       string
	Pin_code                   string
	Available_test_name        string
	Opening_time               string
	Closing_time               string
	Availability               string
	Availability_time_for_test string
}

type Lab_Appointment struct {
	TestAppointmentBookingid int
	Patient_id               int
	Doctor_id                int
	Test_Name                string
	Booking_time             string
}
