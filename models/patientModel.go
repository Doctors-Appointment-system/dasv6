package models

type Patient struct {
	ID                      int
	Name                    string
	Age                     int
	Gender                  string
	Address                 string
	City                    string
	Phone                   string
	Disease                 string
	Selected_specialisation string
	Patient_history         string
}

type Appointment struct {
	Bookingid    int
	Patient_id   int
	Doctor_id    int
	Booking_time string
}
