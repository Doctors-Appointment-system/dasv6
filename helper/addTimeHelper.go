package helper

import "time"

func Add_time(s string) string {
	timeStr := s
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		panic(err)
	}

	t = t.Add(30 * time.Minute)
	newTimeStr := t.Format("15:04")

	return newTimeStr
}
