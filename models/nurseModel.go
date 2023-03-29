package models

type Nurse struct {
	ID             int
	Name           string
	Gender         string
	Address        string
	City           string
	Phone          string
	Specialisation string
	Start_time     string
	End_time       string
	Charge_per_day int
	Availability   string
}
