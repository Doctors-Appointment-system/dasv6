package controllers

import (
	"database/sql"

	"fmt"

	"log"

	"net/http"

	"strings"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"

	"Doctor-Appointment-Project/models"
)

func Add_docter() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}
		var data models.Doctor
		err = c.BindJSON(&data)
		if err != nil {
			return
		}
		c.IndentedJSON(http.StatusCreated, data)
		query_data := fmt.Sprintf(`INSERT INTO Doctor (Name,Gender,Address,City,Phone,Specialisation,Opening_time,Closing_time,Availability,Availability_time,Fees) VALUES ( '%s','%s','%s','%s','%s','%s','%s','%s','%s','%s',%f)`, data.Name, data.Gender, data.Address, data.City, data.Phone, data.Specialisation, data.Opening_time, data.Closing_time, data.Availability, data.Availability_time, data.Fees)
		fmt.Println(query_data)
		//insert data
		insert, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		c.JSON(http.StatusOK, gin.H{"message": "doctor added successfully"})
	}
}

func Get_docter() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("connection not created")
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		fmt.Println("connection is created")
		if err != nil {

			log.Fatal(err)

		}
		fmt.Println("Connection Created")
		results, err := db.Query("SELECT * FROM Doctor")
		fmt.Println("Quary exicuted")

		if err != nil {

			panic(err.Error())

		}

		defer results.Close()

		var output interface{}

		for results.Next() {

			var ID int

			var Name string

			var Gender string

			var Address string

			var City string

			var Phone string

			var Specialisation string

			var Opening_time string

			var Closing_time string

			var Availability string

			var Availability_time string

			var Fees float64

			err = results.Scan(&ID, &Name, &Gender, &Address, &City, &Phone, &Specialisation, &Opening_time, &Closing_time, &Availability, &Availability_time, &Fees)

			if err != nil {

				panic(err.Error())

			}

			output = fmt.Sprintf("%d  '%s'  '%s'  %s  '%s'  '%s'  '%s' '%s' '%s' '%s' '%s' %f", ID, Name, Gender, Address, City, Phone, Specialisation, Opening_time, Closing_time, Availability, Availability_time, Fees)

			c.JSON(http.StatusOK, gin.H{"Data": output})

		}

	}
}

func Update_docter() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var data models.Doctor
		var updateColumns []string
		var args []interface{}

		err = c.BindJSON(&data)

		if err != nil {

			return

		}
		fmt.Println(data)
		if data.Address != "" {
			updateColumns = append(updateColumns, "Address = ?")
			args = append(args, data.Address)
		}
		fmt.Println(updateColumns, args)

		if data.City != "" {
			updateColumns = append(updateColumns, "City = ?")
			args = append(args, data.City)
		}
		fmt.Println(updateColumns, args)
		if data.Phone != "" {
			updateColumns = append(updateColumns, "Phone = ?")
			args = append(args, data.Phone)
		}
		fmt.Println(updateColumns, args)
		if data.Specialisation != "" {
			updateColumns = append(updateColumns, "Specialisation = ?")
			args = append(args, data.Specialisation)
		}

		if data.Opening_time != "" {
			updateColumns = append(updateColumns, "Opening_time = ?")
			args = append(args, data.Opening_time)
		}
		fmt.Println(updateColumns, args)
		if data.Closing_time != "" {
			updateColumns = append(updateColumns, "Closing_time = ?")
			args = append(args, data.Closing_time)
		}
		fmt.Println(updateColumns, args)
		if data.Availability != "" {
			updateColumns = append(updateColumns, "Availability = ?")
			args = append(args, data.Availability)
		}
		fmt.Println(updateColumns, args)

		if data.Availability_time != "" {
			updateColumns = append(updateColumns, "Availability_time = ?")
			args = append(args, data.Availability_time)
		}
		fmt.Println(updateColumns, args)
		if data.Fees != 0 {
			updateColumns = append(updateColumns, "Fees = ?")
			args = append(args, data.Fees)
		}
		fmt.Println(updateColumns, args)
		if len(updateColumns) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No update data provided"})
			return
		}
		fmt.Println(updateColumns, args)
		updateQuery := "UPDATE Doctor SET " + strings.Join(updateColumns, ", ") + " WHERE id = ?"
		args = append(args, data.ID)
		fmt.Println(updateQuery)
		stmt, err := db.Prepare(updateQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer stmt.Close()
		if _, err := stmt.Exec(args...); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusCreated, data)

		c.JSON(http.StatusOK, gin.H{"message": "Doctor updated successfully"})

	}
}

func Delete_docter() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var data models.Doctor

		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		// _, err = db.Exec("DELETE FROM Dost WHERE id = 10")

		delete_query := fmt.Sprintf("DELETE FROM doctor WHERE ID = %d", data.ID)

		delete, err := db.Query(delete_query)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return

		}

		defer delete.Close()

		c.JSON(http.StatusOK, gin.H{"message": "Doctor Deleted successfully"})

	}
}

func GetDoctorByLocation() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("connection not created")
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		fmt.Println("connection is created")
		if err != nil {

			log.Fatal(err)

		}
		fmt.Println("Connection Created")

		var data models.Doctor
		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		query_data := fmt.Sprintf("SELECT * FROM Doctor WHERE City='%s' AND Specialisation='%s'", data.City, data.Specialisation)
		fmt.Println(query_data)

		results, err := db.Query(query_data)
		fmt.Println("Quary exicuted")

		if err != nil {

			panic(err.Error())

		}

		defer results.Close()

		var output interface{}

		for results.Next() {

			var ID int

			var Name string

			var Gender string

			var Address string

			var City string

			var Phone string

			var Specialisation string

			var Opening_time string

			var Closing_time string

			var Availability string

			var Availability_time string

			var Fees float64

			err = results.Scan(&ID, &Name, &Gender, &Address, &City, &Phone, &Specialisation, &Opening_time, &Closing_time, &Availability, &Availability_time, &Fees)

			if err != nil {

				panic(err.Error())

			}

			output = fmt.Sprintf("%d  '%s'  '%s'  %s  '%s'  '%s'  '%s' '%s' '%s' '%s' '%s' %f", ID, Name, Gender, Address, City, Phone, Specialisation, Opening_time, Closing_time, Availability, Availability_time, Fees)

			c.JSON(http.StatusOK, gin.H{"Data": output})

		}

	}
}

func CheckMyAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var data models.Doctor

		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		// _, err = db.Exec("DELETE FROM Dost WHERE id = 10")

		Get_query := fmt.Sprintf("SELECT patient.Name,patient.Age,patient.Gender,patient.Address,patient.City,patient.Phone,patient.Disease,patient.Patient_history,appointment.Booking_time FROM appointment INNER JOIN patient ON appointment.Patient_id = patient.id where appointment.Doctor_id = %d", data.ID)

		GetResult, err := db.Query(Get_query)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return

		}

		defer GetResult.Close()

		c.JSON(http.StatusOK, gin.H{"message": "Your Appointment"})

		var output interface{}

		for GetResult.Next() {

			var Name string

			var Age int

			var Gender string

			var Address string

			var City string

			var Phone string

			var Disease string

			var Selected_Specialisation string

			var Patient_history string

			var Booking_time string

			err = GetResult.Scan(&Name, &Age, &Gender, &Address, &City, &Phone, &Disease, &Patient_history, &Booking_time)

			if err != nil {

				panic(err.Error())

			}

			output = fmt.Sprintf("'%s' %d  '%s'  %s  '%s'  '%s'  '%s' '%s' '%s','%s", Name, Age, Gender, Address, City, Phone, Disease, Selected_Specialisation, Patient_history, Booking_time)

			c.JSON(http.StatusOK, gin.H{"Data": output})

		}

	}
}

func Doctor_Checking_Feedback() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var data models.Feedback

		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		// _, err = db.Exec("DELETE FROM Dost WHERE id = 10")

		Get_query := fmt.Sprintf("SELECT patient.Name,feedback.Rating,feedback.feedback_msg FROM feedback INNER JOIN patient ON feedback.Patient_id = patient.id where appointment.Doctor_id = %d", data.ID)

		GetResult, err := db.Query(Get_query)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return

		}

		defer GetResult.Close()

		c.JSON(http.StatusOK, gin.H{"message": "Your Appointment"})

		var output interface{}

		for GetResult.Next() {

			var Name string
			var Rating int
			var Feedback_msg string

			err = GetResult.Scan(&Name, &Rating, &Feedback_msg)

			if err != nil {

				panic(err.Error())

			}

			output = fmt.Sprintf("'%s' %d '%s'", Name, Rating, Feedback_msg)

			c.JSON(http.StatusOK, gin.H{"Data": output})

		}
	}
}
