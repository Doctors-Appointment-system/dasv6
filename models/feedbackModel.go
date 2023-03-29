package models

type Doctor_feedback struct {
	ID           int
	Patient_id   int
	Doctor_id    int
	Rating       int
	Feedback_msg string
}

type Nurse_feedback struct {
	ID           int
	Patient_id   int
	Nurse_id     int
	Rating       int
	Feedback_msg string
}
