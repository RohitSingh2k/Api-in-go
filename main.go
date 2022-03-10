package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)



type Account struct {
	Phone   int64  `json:"phone"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Balance int    `json:"balance"`
}

type Add struct {
	Address string `json:"address"`
}



// Our IN memory DB
var DB = make(map[int64]Account)




// This function create a new account ---------------------------------------------------
func CreateAccount(c *gin.Context) {

	var user Account
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Bad credential for creating an account",
		})
		return
	}

	DB[user.Phone] = user

	c.JSON(200, gin.H{
		"success": true,
		"message": "Account created successfully",
		"data":    DB[user.Phone],
	})

}

// This function get the account details -------------------------------------------------
func GetAccountDetails(c *gin.Context) {

	phone, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if value, ok := DB[phone]; ok {
		c.JSON(200, gin.H{
			"success": true,
			"data":    value,
		})
	} else {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Account does not exist into the db",
		})
	}

}

// This function get account balance of an account ------------------------------------------------
func GetAccountBalance(c *gin.Context) {
	phone, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if value, ok := DB[phone]; ok {
		c.JSON(200, gin.H{
			"success": true,
			"balance": value.Balance,
		})
	} else {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Account does not exist into the db",
		})
	}
}

// This function updates the address of an account -----------------------------------------
func UpdateAddress(c *gin.Context) {
	phone, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if _, ok := DB[phone]; ok {

		var address Add
		body := json.NewDecoder(c.Request.Body)

		err := body.Decode(&address)
		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": "Bad request",
			})

		} else {

			// here update the map address
			prev := DB[phone]
			prev.Address = address.Address
			DB[phone] = prev

			c.JSON(200, gin.H{
				"success": true,
				"message": "Address updated successfully",
			})
		}

	} else {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Account does not exist into the db",
		})
	}
}

func main() {
	app := gin.Default()

	app.POST("/account", CreateAccount)
	app.GET("/account/:id", GetAccountDetails)
	app.GET("/account/:id/balance", GetAccountBalance)
	app.PUT("/account/:id/address", UpdateAddress)

	app.Run("localhost:9000")

}
