package models

type Feedback struct {
	ID           int
	Patient_id   int
	Doctor_id    int
	Rating       int
	Feedback_msg string
}

type Lab_feedback struct {
	ID           int
	Patient_id   int
	Lab_id       int
	Rating       int
	Feedback_msg string
}
