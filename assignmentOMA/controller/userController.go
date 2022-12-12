package controller

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mayurkhairnar2525/restaurantManagement/helpers"
	"github.com/mayurkhairnar2525/restaurantManagement/modals"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// HashPassword  will generate hash from the password
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal("error: while generating hash from the password")
	}
	return string(bytes)
}

func VerifyPassword(userPassword, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("email of password is incorrect")
		check = false
	}
	return check, msg
}

func (r *RestaurRepo) SignUp(c *gin.Context) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var user modals.User

	// Convert the JSON data coming from postman to something that golang understands
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Validate the data based on user struct
	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	// hash password
	password := HashPassword(user.Password)
	user.Password = password

	// Creating the uid
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	user.User_id = uuid

	// Generate token and refresh token (generate all tokens function from helper)
	token, refreshToken, _ := helpers.GenerateAllTokens(user.Email, user.First_name, user.Last_name, user.User_id)

	user.Token = &token
	user.Refresh_Token = &refreshToken

	// if all ok, then you insert this new user into the user collection
	err := modals.CreateUser(r.Db, &user)

	defer cancel()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "User created")
}

func (r *RestaurRepo) Login(c *gin.Context) {
	var user = &modals.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp := r.FindOne(user.Email, user.Password)
	c.JSON(http.StatusOK, resp)
}

func (r *RestaurRepo) FindOne(email, password string) map[string]interface{} {
	var user = &modals.User{}
	var foundUser modals.User

	if err := r.Db.Where("email=?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}

	expiresAt := time.Now().Add(time.Minute * 1000).Unix()

	// Checking the password
	// If found.user password is same as user password
	_, _ = VerifyPassword(user.Password, foundUser.Password)

	tkn := &helpers.SignedDetails{
		Email: user.Email,

		Uid: user.User_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tkn)
	_, err := token.SignedString([]byte("SECRET_KEY"))
	if err != nil {
		log.Fatal("error", err)
	}
	var resp = map[string]interface{}{"status": true, "message": "logged in"}

	// Storing the token in the response
	//	resp["token"] = tokenString
	resp["user"] = user
	return resp
}

func (r *RestaurRepo) GetUserByID(c *gin.Context) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user modals.User
	Id := c.Param("id")
	userID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to convert user id from int to string",
		})
		return
	}
	err = modals.GetUserByID(r.Db, &user, uint(userID))
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error occurred while listing user items",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (r *RestaurRepo) GetUserBy_userID(c *gin.Context) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user modals.User
	userID := c.Param("user_id")

	_, err := modals.GetUserByUserID(r.Db, &user, userID)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error occurred while listing user items",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (r *RestaurRepo) GetUsers(c *gin.Context) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var users []modals.User
	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	// Query for with page you want
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	// Skipping index
	_ = (page - 1) * recordPerPage
	_, err = strconv.Atoi(c.Query("startIndex"))

	err = modals.GetAllUsers(r.Db, &users)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing user items"})
	}
	c.JSON(http.StatusOK, users)
}
