package controllers

import (
	"database/sql"
	"strings"

	"fmt"

	"log"

	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"

	"Doctor-Appointment-Project/models"

	helper "Doctor-Appointment-Project/helper"
)

func Add_lab_details() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}
		var data models.Lab
		err = c.BindJSON(&data)
		if err != nil {
			return
		}
		c.IndentedJSON(http.StatusCreated, data)
		query_data := fmt.Sprintf(`INSERT INTO Lab (Lab_Name,Lab_Operator,Phone,Address,City,Pin_code,Available_test_name,Opening_time,Closing_time,Availability,Availability_time_for_test) VALUES ( '%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')`, data.Lab_Name, data.Lab_Operator_Name, data.Phone, data.Address, data.City, data.Available_test_name, data.Opening_time, data.Closing_time, data.Availability, data.Availability_time_for_test)
		fmt.Println(query_data)
		//insert data
		insert, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		c.JSON(http.StatusOK, gin.H{"message": "Lab added successfully"})

		query_data2 := fmt.Sprintf(`SELECT Labid,Lab_name FROM Lab`)
		fmt.Println(query_data2)
		//insert data
		read, err := db.Query(query_data2)
		if err != nil {
			panic(err.Error())
		}
		defer read.Close()

		var output interface{}

		for read.Next() {

			var Labid int

			var Lab_Name string

			err = read.Scan(&Labid, &Lab_Name)

			if err != nil {

				panic(err.Error())

			}

			output = fmt.Sprintf("%d  '%s'", Labid, Lab_Name)

			c.JSON(http.StatusOK, gin.H{"Your Lab_Id and Name - ": output})

		}

	}
}

func Update_lab() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var data models.Lab
		var updateColumns []string
		var args []interface{}

		err = c.BindJSON(&data)

		if err != nil {

			return

		}
		if data.Address != "" {
			updateColumns = append(updateColumns, "Lab_Operator_Name = ?")
			args = append(args, data.Lab_Operator_Name)
		}
		if data.Phone != "" {
			updateColumns = append(updateColumns, "Phone = ?")
			args = append(args, data.Phone)
		}
		fmt.Println(updateColumns, args)
		fmt.Println(updateColumns, args)
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
		if data.Pin_code != "" {
			updateColumns = append(updateColumns, "Pin_code = ?")
			args = append(args, data.Pin_code)
		}
		fmt.Println(updateColumns, args)
		if data.Available_test_name != "" {
			updateColumns = append(updateColumns, "Available_test_name = ?")
			args = append(args, data.Available_test_name)
		}
		fmt.Println(updateColumns, args)
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

		if data.Availability_time_for_test != "" {
			updateColumns = append(updateColumns, "Availability_time = ?")
			args = append(args, data.Availability_time_for_test)
		}
		fmt.Println(updateColumns, args)
		if len(updateColumns) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No update data provided"})
			return
		}
		fmt.Println(updateColumns, args)
		updateQuery := "UPDATE Lab SET " + strings.Join(updateColumns, ", ") + " WHERE Labid = ?"
		args = append(args, data.Labid)
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

		c.JSON(http.StatusOK, gin.H{"message": "Lab Data updated successfully"})

	}

}

func Delete_lab() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var data models.Lab

		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		delete_query := fmt.Sprintf("DELETE FROM lab WHERE Labid= %d", data.Labid)

		delete, err := db.Query(delete_query)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return

		}

		defer delete.Close()

		c.JSON(http.StatusOK, gin.H{"message": "lab Deleted successfully"})

	}

}

func Get_Lab_Profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}
		var data models.Lab
		err = c.BindJSON(&data)
		if err != nil {
			return
		}
		get_detail := fmt.Sprintf("SELECT * FROM lab WHERE Labid = %d", data.Labid)
		detail, err := db.Query(get_detail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer detail.Close()

		var output interface{}
		for detail.Next() {

			var Labid int
			var Lab_Name string
			var Lab_Operator_Name string
			var Phone string
			var Address string
			var City string
			var Pin_code string
			var Available_test_name string
			var Opening_time string
			var Closing_time string
			var Availability string
			var Availability_time_for_test string
			err = detail.Scan(&Labid, &Lab_Name, &Lab_Operator_Name, &Phone, &Address, &City, &Phone, &Pin_code, &Available_test_name, &Opening_time, &Closing_time, &Availability, &Availability_time_for_test)

			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d  '%s' '%s'  '%s'  '%s'  '%s' '%s' '%s' '%s'  '%s'  '%s'", Labid, Lab_Name, Lab_Operator_Name, Phone, Address, City, Available_test_name, Opening_time, Closing_time, Availability, Availability_time_for_test)

			fmt.Println(output)

			c.JSON(http.StatusOK, gin.H{"Lab details": output})

		}

		c.JSON(http.StatusOK, gin.H{"message": "Your Lab details"})

	}
}

func GetLabByLocation() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("connection not created")
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		fmt.Println("connection is created")
		if err != nil {

			log.Fatal(err)

		}
		fmt.Println("Connection Created")

		var data models.Lab
		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		query_data := fmt.Sprintf("SELECT * FROM Lab WHERE City='%s'", data.City)
		fmt.Println(query_data)

		detail, err := db.Query(query_data)
		fmt.Println("Quary exicuted")

		if err != nil {

			panic(err.Error())

		}

		defer detail.Close()

		var output interface{}
		for detail.Next() {

			var Labid int
			var Lab_Name string
			var Lab_Operator_Name string
			var Phone string
			var Address string
			var City string
			var Pin_code string
			var Available_test_name string
			var Opening_time string
			var Closing_time string
			var Availability string
			var Availability_time_for_test string
			err = detail.Scan(&Labid, &Lab_Name, &Lab_Operator_Name, &Phone, &Address, &City, &Phone, &Pin_code, &Available_test_name, &Opening_time, &Closing_time, &Availability, &Availability_time_for_test)

			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d  '%s' '%s'  '%s'  '%s'  '%s' '%s' '%s' '%s'  '%s'  '%s'", Labid, Lab_Name, Lab_Operator_Name, Phone, Address, City, Available_test_name, Opening_time, Closing_time, Availability, Availability_time_for_test)

			fmt.Println(output)

			c.JSON(http.StatusOK, gin.H{"Lab details": output})

		}

	}
}

func Book_test() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}
		// var Doctor_str doctor
		var booking_data models.Lab_Appointment
		var lab_data models.Lab
		err = c.BindJSON(&booking_data)
		if err != nil {
			log.Fatal(err)
		}

		get_booking_time := fmt.Sprintf("SELECT Availability_time,Closing_time FROM Lab WHERE Labid = %d", lab_data.Labid)
		lab_result, err := db.Query(get_booking_time)
		// lab_result,err := db.Exec(get_booking_time)
		if err != nil {
			c.JSON(404, gin.H{"error": "Lab not found"})
			return
		}

		var people []models.TimeStr

		for lab_result.Next() {
			var p models.TimeStr
			if err := lab_result.Scan(&p.Availability_time, &p.Closing_time); err != nil {
				log.Fatal(err)
			}
			people = append(people, p)
		}

		if err := lab_result.Err(); err != nil {
			log.Fatal(err)
		}

		var booktime string = people[0].Availability_time
		var Closing_time string = people[1].Closing_time

		if Closing_time == booktime {
			c.JSON(http.StatusOK, gin.H{"message": "No slot available visit again thank you"})

		}

		c.IndentedJSON(http.StatusCreated, booking_data)

		booking_data.Booking_time = booktime

		query_data := fmt.Sprintf(`INSERT INTO TestAppointment(Patient_id,Doctor_id,Test_Name,Booking_time) VALUES(%d,%d,'%s','%s')`, booking_data.Patient_id, booking_data.Doctor_id, booking_data.Test_Name, booking_data.Booking_time)
		_, err = db.Exec(query_data)
		if err != nil {

			panic(err.Error())

		}
		t1 := helper.Add_time(booktime)

		query_data2 := fmt.Sprintf(`UPDATE Lab SET Availability_time = '%s' WHERE Labid = %d`, t1, lab_data.Labid)

		fmt.Println(query_data2)

		_, err = db.Query(query_data2)
		if err != nil {

			panic(err.Error())

		}

		if err != nil {

			panic(err.Error())

		}

		c.JSON(http.StatusOK, gin.H{"message": "Your Appointment successfully Booked"})

	}
}

func Lab_feedback() gin.HandlerFunc {
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

		query_data := fmt.Sprintf(`INSERT INTO Lab_feedback(Patient_id,Labid,Rating,Feedback_msg) VALUES(%d,%d,%d,'%s')`, data.Patient_id, data.Doctor_id, data.Rating, data.Feedback_msg)

		fmt.Println(query_data)

		//insert data

		insert, err := db.Query(query_data)

		if err != nil {

			panic(err.Error())

		}

		defer insert.Close()

	}
}
