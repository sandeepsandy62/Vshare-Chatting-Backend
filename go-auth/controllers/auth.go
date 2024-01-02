package controllers

import (
    "go-auth/models"
    "time"

    "go-auth/utils"

    "github.com/golang-jwt/jwt/v5"
    "github.com/gin-gonic/gin"
)


var jwtKey = []byte("my_secret_key")



// Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context.
// Any function takes gin.Context pointer is a handler  
func Login(c *gin.Context){
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nill {

		//Context.JSON will write the JSON data to the response body
		/* 
		he gin.H method in Golang is used to create a new H object. 
		An H object is a map of string keys to interface{} values.
		It is often used to represent the data returned by a Gin handler.
		*/
		c.JSON(400,gin.H{"error": err.Error()})
		return 
	}

	var existingUser models.User 

	models.DB.Where("email = ?" , user.Email).First(&existingUser)

	if existingUser.ID == 0 {
		c.JSON(400 , gin.H{"error":"user does not exist"})
		return
	}

	errHash := utils.CompareHashPassword(user.Password , existingUser.Password)

	if !errHash {
		c.JSON(400,gin.H{"error":"invalid password"})
		return 
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &models.Claims{
		Role : existingUser.Role, 
		StandardClaims : jwt.StandardClaims{
			Subject : existingUser.Email , 
			ExpiresAt : expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	tokenString , err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(500 , gin.H{"error" , "could not generate token"})
		return 
	}

	c.SetCookie("token",tokenString,int(expirationTime.Unix()),"/",false,true)
	c.JSON(200,gin.H{"success":"user logged in"})


}


