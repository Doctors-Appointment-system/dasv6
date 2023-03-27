package models

type Doctor struct {
	ID                int
	Name              string
	Gender            string
	Address           string
	City              string
	Phone             string
	Specialisation    string
	Opening_time      string
	Closing_time      string
	Availability      string
	Availability_time string
	Fees              float64
}
