package controllers

import (
	"database/sql"

	"fmt"

	"log"

	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"

	"Doctor-Appointment-Project/models"

	helper "Doctor-Appointment-Project/helper"
)

func Addpatient() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("add patient")

		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var data models.Patient

		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		c.IndentedJSON(http.StatusCreated, data)

		query_data := fmt.Sprintf(`INSERT INTO patient(Name,Age,Gender,Address,City,Phone,Disease,Selected_Specialisation,Patient_history) VALUES('%s',%d,'%s','%s','%s','%s','%s','%s','%s')`, data.Name, data.Age, data.Gender, data.Address, data.City, data.Phone, data.Disease, data.Selected_specialisation, data.Patient_history)

		fmt.Println(query_data)

		//insert data

		insert, err := db.Query(query_data)

		if err != nil {

			panic(err.Error())

		}

		defer insert.Close()

	}
}
func Get_my_details() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}
		var mob models.Patient
		err = c.BindJSON(&mob)
		if err != nil {
			return
		}
		get_detail := fmt.Sprintf("SELECT * FROM Patient WHERE Phone = '%s'", mob.Phone)
		detail, err := db.Query(get_detail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer detail.Close()

		var output interface{}
		for detail.Next() {

			var ID int
			var Name string
			var Age int
			var Gender string
			var Address string
			var City string
			var Phone string
			var Disease string
			var Selected_specialisation string
			var Patient_history string
			err = detail.Scan(&ID, &Name, &Age, &Gender, &Address, &City, &Phone, &Disease, &Selected_specialisation, &Patient_history)

			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d  '%s'  %d  '%s'  '%s'  '%s'  '%s' '%s' %s' '%s'  ", ID, Name, Age, Gender, Address, City, Phone, Disease, Selected_specialisation, Patient_history)

			fmt.Println(output)

			c.JSON(http.StatusOK, gin.H{"Patient details": output})

		}

		c.JSON(http.StatusOK, gin.H{"message": "Your details"})

	}
}

// func Getpatient() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
// 		if err != nil {

// 			log.Fatal(err)

// 		}
// 		fmt.Println("Connection Created")
// 		results, err := db.Query("SELECT * FROM Patient")
// 		fmt.Println("Quary exicuted")
//

// 		if err != nil {

// 			panic(err.Error())

// 		}

// 		defer results.Close()

// 		var output interface{}

// 		for results.Next() {

// 			var ID int
// 			var Name string
// 			var Age int
// 			var Gender string
// 			var Address string
// 			var City string
// 			var Phone string
// 			var Disease string
// 			var Selected_specialisation string
// 			var Patient_history string
// 			err = results.Scan(&ID, &Name, &Age, &Gender, &Address, &City, &Phone, &Disease, &Selected_specialisation, &Patient_history)

// 			if err != nil {

// 				panic(err.Error())

// 			}

// 			output = fmt.Sprintf("%d  '%s'  %d  '%s'  '%s'  '%s'  '%s' '%s' %s' '%s'  ", ID, Name, Age, Gender, Address, City, Phone, Disease, Selected_specialisation, Patient_history)

// 			fmt.Println(output)

// 			c.JSON(http.StatusOK, gin.H{"Data": output})

// 		}

// 	}
// }

func DeletePatient() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var data models.Patient

		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		// _, err = db.Exec("DELETE FROM Dost WHERE id = 10")

		delete_query := fmt.Sprintf("DELETE FROM patient WHERE ID= %d", data.ID)

		delete, err := db.Query(delete_query)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return

		}

		defer delete.Close()

		c.JSON(http.StatusOK, gin.H{"message": "Patient Deleted successfully"})

	}
}

func BookingAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}
		// var Doctor_str doctor
		var booking_data models.Appointment
		var doctor_data models.Doctor
		err = c.BindJSON(&booking_data)
		if err != nil {
			log.Fatal(err)
		}

		get_booking_time := fmt.Sprintf("SELECT Availability_time,Closing_time FROM Doctor WHERE id = %d", doctor_data.ID)
		doctor_result, err := db.Query(get_booking_time)
		// doctor_result,err := db.Exec(get_booking_time)
		if err != nil {
			c.JSON(404, gin.H{"error": "Doctor not found"})
			return
		}

		var people []models.TimeStr

		for doctor_result.Next() {
			var p models.TimeStr
			if err := doctor_result.Scan(&p.Availability_time, &p.Closing_time); err != nil {
				log.Fatal(err)
			}
			people = append(people, p)
		}

		if err := doctor_result.Err(); err != nil {
			log.Fatal(err)
		}

		var booktime string = people[0].Availability_time
		var Closing_time string = people[1].Closing_time

		if Closing_time == booktime {
			c.JSON(http.StatusOK, gin.H{"message": "No slot available visit again thank you"})

		}

		c.IndentedJSON(http.StatusCreated, booking_data)

		booking_data.Booking_time = booktime

		query_data := fmt.Sprintf(`INSERT INTO appointment (Patient_id,Doctor_id,Booking_time) VALUES(%d,%d,'%s')`, booking_data.Patient_id, booking_data.Doctor_id, booking_data.Booking_time)
		_, err = db.Exec(query_data)
		if err != nil {

			panic(err.Error())

		}
		t1 := helper.Add_time(booktime)

		query_data2 := fmt.Sprintf(`UPDATE Doctor SET Availability_time = '%s' WHERE ID = %d`, t1, booking_data.Doctor_id)

		fmt.Println(query_data2)

		_, err = db.Query(query_data2)
		if err != nil {

			panic(err.Error())

		}

		c.JSON(http.StatusOK, gin.H{"message": "Your Appointment successfully Booked"})

	}
}

func Cancel_appointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var data models.Appointment

		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		c.IndentedJSON(http.StatusCreated, data)

		query_data := fmt.Sprintf("DELETE FROM appointment WHERE id =%d", data.Bookingid)

		_, err = db.Exec(query_data)

		if err != nil {

			panic(err.Error())

		}

		c.JSON(http.StatusOK, gin.H{"message": "Cancel Appointment successfully"})

	}
}

func Doctor_feedback() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("add feedback")

		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var data models.Feedback

		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		c.IndentedJSON(http.StatusCreated, data)

		query_data := fmt.Sprintf(`INSERT INTO Feedback(Patient_id,Doctor_id,Rating,Feedback_msg) VALUES(%d,%d,%d,'%s')`, data.Patient_id, data.Doctor_id, data.Rating, data.Feedback_msg)

		fmt.Println(query_data)

		//insert data

		insert, err := db.Query(query_data)

		if err != nil {

			panic(err.Error())

		}

		defer insert.Close()

	}
}
