package controllers

import (
	"database/sql"
	"net/http"

	"fmt"

	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"

	"Doctor-Appointment-Project/models"
)

func Get_order() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {
			log.Fatal(err)
		}

		var data models.Order_medicines
		err = c.BindJSON(&data)

		if err != nil {
			return
		}
		query := fmt.Sprintf("SELECT Patient.Name,Patient.Address,Patient.City,Patient.Phone,Order_medicines.Prescription FROM Order_medicines INNER JOIN Patient ON Order_medicines.Patient_id = Patient.id where Order_medicines.Order_id = %d", data.Order_id)
		result, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer result.Close()

		var output interface{}

		for result.Next() {
			var Name string

			var Address string

			var City string

			var Phone string

			var Prescription string

			err = result.Scan(&Name, &Address, &City, &Phone, &Prescription)

			if err != nil {
				panic(err.Error())
			}

			output = fmt.Sprintf("'%s' '%s' '%s' '%s'  '%s' ", Name, Address, City, Phone, Prescription)

			c.JSON(http.StatusOK, gin.H{"Data": output})

		}

	}

}
