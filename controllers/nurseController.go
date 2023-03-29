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
)

func Add_nurse() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}
		var data models.Nurse
		err = c.BindJSON(&data)
		if err != nil {
			return
		}
		c.IndentedJSON(http.StatusCreated, data)
		query_data := fmt.Sprintf(`INSERT INTO Nurse(Name,Gender,Address,City,Phone,Specialisation,Start_time,End_time,Charge_per_day,Availability,) VALUES('%s','%s','%s','%s','%s','%s','%s','%s','%d','%s')`, data.Name, data.Gender, data.Address, data.City, data.Phone, data.Specialisation, data.Start_time, data.End_time, data.Charge_per_day, data.Availability)
		fmt.Println(query_data)

		// Data insertion operation

		insert, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		c.JSON(http.StatusOK, gin.H{"message": "Nurse Added Successfully"})

		// METHOD TO SHOW THAT HE/SHE IS ADDED AND WHAT THEIR ID IS

		query_data2 := fmt.Sprintf(`SELECT ID,Name FROM Nurse`)
		fmt.Println(query_data2)

		//insert data

		read, err := db.Query(query_data2)
		if err != nil {
			panic(err.Error())
		}
		defer read.Close()
		var output interface{}
		for read.Next() {
			var ID int
			var Name string
			err = read.Scan(&ID, &Name)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d  '%s'", ID, Name)
			c.JSON(http.StatusOK, gin.H{"Your Id and Name is  ": output})
		}
	}
}

// Method to get Profile of Nurse

func Get_nurse_profile() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			return
		}
		var mob models.Nurse
		err = c.BindJSON(&mob)
		if err != nil {
			log.Fatal(err)
		}
		get_details := fmt.Sprintf("SELECT * FROM NURSE WHERE Phone='%s'", mob.Phone)
		details, err := db.Query(get_details)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer details.Close()

		var output interface{}
		for details.Next() {
			var ID int
			var Name string
			var Gender string
			var Address string
			var City string
			var Phone string
			var Specialisation string
			var Start_time string
			var End_time string
			var Charge_per_day int
			var Availability string

			err = details.Scan(&ID, &Name, &Gender, &Address, &City, &Phone, &Specialisation, &Start_time, &End_time, &Charge_per_day, &Availability)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d '%s' '%s' '%s' '%s' '%s' '%s' '%s' '%s' %d '%s'", ID, Name, Gender, Address, City, Phone, Specialisation, Start_time, End_time, Charge_per_day, Availability)
			fmt.Println(output)
			c.JSON(http.StatusOK, gin.H{" Nurse Details ": output})
		}
	}
}

// Method to get the list of Nurse

func Get_nurse() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			return
		}
		fmt.Println("Connection Established")
		result, err := db.Query("SELECT * FROM Nurse")
		if err != nil {
			panic(err.Error())
		}
		defer result.Close()

		var output interface{}
		for result.Next() {
			var ID int
			var Name string
			var Gender string
			var Address string
			var City string
			var Phone string
			var Specialisation string
			var Start_time string
			var End_time string
			var Charge_per_day int
			var Availability string

			err = result.Scan(&ID, &Name, &Gender, &Address, &City, &Phone, &Specialisation, &Start_time, &End_time, &Charge_per_day, &Availability)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d '%s' '%s' '%s' '%s' '%s' '%s' '%s' '%s' %d '%s'", ID, Name, Gender, Address, City, Phone, Specialisation, Start_time, End_time, Charge_per_day, Availability)
			fmt.Println(output)
			c.JSON(http.StatusOK, gin.H{" All Registerd Nurse are ": output})

		}
	}
}

// Method to update the info about the nurse

func Update_nurse() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}

		var data models.Nurse
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

		if data.Start_time != "" {
			updateColumns = append(updateColumns, "Start_time  = ?")
			args = append(args, data.Start_time)
		}
		fmt.Println(updateColumns, args)
		if data.End_time != "" {
			updateColumns = append(updateColumns, "End_time  = ?")
			args = append(args, data.End_time)
		}
		fmt.Println(updateColumns, args)

		if data.Charge_per_day != 0 {
			updateColumns = append(updateColumns, "Fees = ?")
			args = append(args, data.Charge_per_day)
		}
		fmt.Println(updateColumns, args)

		if data.Availability != "" {
			updateColumns = append(updateColumns, "Availability = ?")
			args = append(args, data.Availability)
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

		c.JSON(http.StatusOK, gin.H{"message": "Nurse Information updated successfully"})

	}
}

// Method to delete the nurese from the database

func Delete_nurse() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}
		var data models.Nurse
		err = c.BindJSON(&data)
		if err != nil {
			return
		}

		delete_query := fmt.Sprintf("DELETE FROM Nurse WHERE ID = %d", data.ID)
		delete, err := db.Query(delete_query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer delete.Close()
		c.JSON(http.StatusOK, gin.H{"message": "Nurse removed from database successfully"})
	}
}

// method to list out nurses based on location

func Get_nurse_by_city() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306),das")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection Established")

		var data models.Nurse
		err = c.BindJSON(&data)
		if err != nil {
			panic(err.Error())
		}
		query_data := fmt.Sprintf("SELECT * FROM Nurse WHERE City='%s'", data.City)

		result, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}
		defer result.Close()

		var output interface{}
		for result.Next() {
			var ID int
			var Name string
			var Gender string
			var Address string
			var City string
			var Phone string
			var Specialisation string
			var Start_time string
			var End_time string
			var Charge_per_day int
			var Availability string

			err = result.Scan(&ID, &Name, &Gender, &Address, &City, &Phone, &Specialisation, &Start_time, &End_time, &Charge_per_day, &Availability)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d '%s' '%s' '%s' '%s' '%s' '%s' '%s' '%s' %d '%s'", ID, Name, Gender, Address, City, Phone, Specialisation, Start_time, End_time, Charge_per_day, Availability)
			fmt.Println(output)
			c.JSON(http.StatusOK, gin.H{" All Registerd Nurse are ": output})
		}
	}
}

// method to list all nusrse based on specialisation

func Get_nurse_by_specialisation() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306),das")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection Established")

		var data models.Nurse
		err = c.BindJSON(&data)
		if err != nil {
			panic(err.Error())
		}
		query_data := fmt.Sprintf("SELECT * FROM Nurse WHERE Specialisation='%s'", data.Specialisation)

		result, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}
		defer result.Close()

		var output interface{}
		for result.Next() {
			var ID int
			var Name string
			var Gender string
			var Address string
			var City string
			var Phone string
			var Specialisation string
			var Start_time string
			var End_time string
			var Charge_per_day int
			var Availability string

			err = result.Scan(&ID, &Name, &Gender, &Address, &City, &Phone, &Specialisation, &Start_time, &End_time, &Charge_per_day, &Availability)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d '%s' '%s' '%s' '%s' '%s' '%s' '%s' '%s' %d '%s'", ID, Name, Gender, Address, City, Phone, Specialisation, Start_time, End_time, Charge_per_day, Availability)
			fmt.Println(output)
			c.JSON(http.StatusOK, gin.H{" All Registerd Nurse are ": output})
		}
	}
}

// GET NURSE BY LOCATION

func Get_nurse_by_location() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306),das")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection Established")

		var data models.Nurse
		err = c.BindJSON(&data)
		if err != nil {
			panic(err.Error())
		}
		query_data := fmt.Sprintf("SELECT * FROM Nurse WHERE City='%s',Specialisation='%s'", data.City, data.Specialisation)

		result, err := db.Query(query_data)
		if err != nil {
			panic(err.Error())
		}
		defer result.Close()

		var output interface{}
		for result.Next() {
			var ID int
			var Name string
			var Gender string
			var Address string
			var City string
			var Phone string
			var Specialisation string
			var Start_time string
			var End_time string
			var Charge_per_day int
			var Availability string

			err = result.Scan(&ID, &Name, &Gender, &Address, &City, &Phone, &Specialisation, &Start_time, &End_time, &Charge_per_day, &Availability)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d '%s' '%s' '%s' '%s' '%s' '%s' '%s' '%s' %d '%s'", ID, Name, Gender, Address, City, Phone, Specialisation, Start_time, End_time, Charge_per_day, Availability)
			fmt.Println(output)
			c.JSON(http.StatusOK, gin.H{" All Registerd Nurse are ": output})
		}
	}
}

// method for nurse to check his/her appointment

func Check_nurse_appointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection Established")
		var data models.Nurse
		err = c.BindJSON(&data)
		if err != nil {
			return
		}
		query_data := fmt.Sprintf("SELECT Patient.Name,Patient.Age,Patient.Gender,Patient.Address,Patient.City,Patient.Phone,Patient.Disease,Patient.Patient_history,appointment.Booking_time FROM appointment INNER JOIN Patient ON appointment.Patient_id = Patient.id where appointment.Doctor_id = %d", data.ID)
		result, err := db.Query(query_data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer result.Close()
		c.JSON(http.StatusOK, gin.H{"message": "Your Appointment"})

		var output interface{}

		for result.Next() {
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
			err = result.Scan(&Name, &Age, &Gender, &Address, &City, &Phone, &Disease, &Patient_history, &Booking_time)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("'%s' %d  '%s'  %s  '%s'  '%s'  '%s' '%s' '%s','%s", Name, Age, Gender, Address, City, Phone, Disease, Selected_Specialisation, Patient_history, Booking_time)
			c.JSON(http.StatusOK, gin.H{"Data": output})
		}
	}
}

// method to check feedback by nurse

func Nurse_checking_feedback() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}
		var data models.Nurse_feedback
		err = c.BindJSON(&data)
		if err != nil {
			return
		}

		Get_query := fmt.Sprintf("SELECT Patient.Name,feedback.Rating,feedback.feedback_msg FROM feedback INNER JOIN Patient ON feedback.Patient_id = Patient.id where appointment.Doctor_id = %d", data.ID)
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
