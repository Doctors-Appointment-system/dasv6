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

type Doctor_appointment struct {
	Bookingid    int
	Patient_id   int
	Doctor_id    int
	Booking_time string
}

type Nurse_appointment struct {
	Bookingid    int
	Patient_id   int
	Nurse_id     int
	Booking_time string
}

type Home_visit_appointment struct {
	Bookingid                int
	Patient_id               int
	Doctor_id                int
	Available_for_home_visit string
}

type Prescription struct {
	Prescription_id int
	Patient_id      int
	Doctor_id       int
	Prescription    string
}

type Order_medicines struct {
	Order_id     int
	Patient_id   int
	Doctor_id    int
	Prescription string
}

type Online_consultancy struct {
	Bookingid    int
	Patient_id   int
	Doctor_id    int
	Booking_time string
}
