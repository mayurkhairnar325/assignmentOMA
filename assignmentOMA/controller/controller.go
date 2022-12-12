package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/mayurkhairnar2525/restaurantManagement/database"
	"github.com/mayurkhairnar2525/restaurantManagement/modals"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RestaurRepo struct {
	Db *gorm.DB
}

func New() *RestaurRepo {
	db := database.InitDB()
	return &RestaurRepo{
		Db: db,
	}
}

var validate = validator.New()

func (r *RestaurRepo) GetAllOrders(c *gin.Context) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var orders []modals.OrderApp
	err := modals.GetAll(r.Db, &orders)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
		return
	}
	c.JSON(http.StatusOK, orders)

}

func (r *RestaurRepo) GetOrderByOrderID(c *gin.Context) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	id := c.Param("order_id")
	var orders modals.OrderApp

	err := modals.GetByOrderID(r.Db, &orders, id)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to convert user id from int to string",
		})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (r *RestaurRepo) CreateOrder(c *gin.Context) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var order modals.OrderApp

	if err := c.Bind(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	validationErr := validate.Struct(order)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validationErr.Error(),
		})
		return
	}

	// Creating the uid
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	order.OrderID = uuid

	err := modals.CreateOrder(r.Db, &order)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, order)
}

func (r *RestaurRepo) DeleteOrder(c *gin.Context) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("error", err)
	}
	err = modals.DeleteOrder(r.Db, int64(ID))
	defer cancel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, " Order deleted")

}

func (r *RestaurRepo) ModifyOrder(c *gin.Context) {
	var order modals.OrderApp

	i, _ := strconv.Atoi(c.Param("id"))
	err := modals.GetByID(r.Db, &order, i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.BindJSON(&order)
	err = modals.ModifyOrder(r.Db, &order)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (r *RestaurRepo) CreateOrderUsingUserID(c *gin.Context) {
	//var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	//
	//var user modals.User
	//var order modals.OrderApp
	//
	//if err := c.Bind(&order); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}
	//
	//validationErr := validate.Struct(order)
	//if validationErr != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": validationErr.Error(),
	//	})
	//	return
	//}
	//
	//userID := c.Param("user_id")
	//if user.User_id != "" {
	//	err := modals.GetUserByUserID(r.Db, &user, userID)
	//	defer cancel()
	//
	//	if err != nil {
	//		msg := fmt.Sprintf("message:user was not found")
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"error": msg,
	//		})
	//		return
	//	}
	//}
	//
	//// Creating the uid
	//uuidWithHyphen := uuid.New()
	//uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	//order.OrderID = uuid
	//
	//err := modals.CreateOrder(r.Db, &order)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}
	//defer cancel()
	//c.JSON(http.StatusOK, order)

	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var order modals.OrderApp
	var user modals.User

	if err := c.Bind(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validationErr := validate.Struct(order)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validationErr.Error(),
		})
		return
	}

	userID := c.Param("user_id")

	// Checking the user
	// Finding the user that exists with the help of userID
	// Here, we want to see that the user id, that I've sending
	// actually exists or not. If that exists then only it will
	// order items.
	// To create order user needs to enter the userID
	_, err := modals.GetUserByUserID(r.Db, &user, userID)
	defer cancel()
	if err != nil {
		msg := fmt.Sprintf("User not found")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": msg,
		})
		return
	}

	if *order.User_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user_ID is empty",
		})
		return
	}

	validationErr = validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validationErr.Error(),
		})
		return
	}

	//if sameUserID.User_id != *order.User_id {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": "user id mis-matched",
	//	})
	//	return
	//
	//}

	// Creating the uid
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	order.OrderID = uuid

	err = modals.CreateOrder(r.Db, &order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, order)
}
