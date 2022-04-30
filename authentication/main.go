package main 

import (
	"fmt"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main(){
	r := gin.Default()
	r.GET("/rest/v1/login/:email", login)
	r.Run()
}

func login(c *gin.Context) {
	user := c.Param("email")
	fmt.Printf("User: %v", user)
	password := c.Query("password")
	fmt.Printf("Password: %v", password)
	ok, err := checkPassword(user, password)
	if err != nil {
		c.String(404, "No such user")
		return 
	}else if !ok {
		c.String(401, "Login Failed")
		return 
	}
	c.JSON(200, gin.H{
		"status": "success",
		"token": uuid.New().String(),
		"user": user, 
		"password": password,
		})
}

func checkPassword(user string, password string) (bool, error){
	if user != "marvin@example.com"{
		return false, errors.New("Cannot find user")
	}
	if password == "wrong" {
		return false, nil
	}
	return true, nil 
}