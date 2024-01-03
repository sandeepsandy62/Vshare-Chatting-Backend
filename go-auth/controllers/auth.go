package controllers

import (
    "go-auth/models"
	"go-auth/utils"
    "time"

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
		StandardClaims : jwt.RegisteredClaims{
			Subject : existingUser.Email , 
			ExpiresAt : jwt.NumericDate(expirationTime.Unix()),
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

func Signup(c *gin.Context){
	var user models.User 

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400,gin.H{"error":err.Error()})
		return 
	}

	var existingUser models.User

	models.DB.Where("email=?",user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		c.JSON(400,gin.H{"error":"user already exists"})
		return 
	}

	var errHash error
	user.Password,errHash = utils.GenerateHashPassword(user.Password)

	if errHash != nil {
		c.JSON(500,gin.H{"error":"could not generate password hash"})
		return 
	}

	models.DB.Create(&user)

	c.JSON(200,gin.H{"success":"user created"})
}

func Home(c *gin.Context){
	cookie , err := c.Cookie("token")

	if err != nil {
		c.JSON(401,gin.H{"error":"unauthorized"})
		return
	}

	claims , err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401,gin.H{"error":"unauthorized"})
		return
	}

	if claims.Role != "user" && claims.Role != "admin"{
		c.JSON(401,gin.H{"error":"unauthorized"})
		return 
	}

	c.JSON(200,gin.H{"success":"home page" , "role" : claims.Role})
}

func Premium(c *gin.Context){
	cookie , err := c.Cookie("token")

	if err != nil {
		c.JSON(401,gin.H{"error":"unauthorized"})
		return 
	}

	claims,err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401,gin.H{"error":"unauthorized"})
		return 
	}

	if claims.Role != "admin" {
		c.JSON(401,gin.H{"error":"unauthorized"})
		return
	}

	c.JSON(200,gin.H{"success":"premium page" , "role":claims.Role})
}

func Logout(c *gin.Context){
	c.SetCookie("token","",-1,"/","localhost",false,true)
	c.JSON(200,gin.H{"success":"user logged out"})
}