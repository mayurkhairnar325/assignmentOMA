package modals

import (
	"github.com/jinzhu/gorm"
	"log"
)

type OrderApp struct {
	gorm.Model

	OrderID    string  `json:"order_id,omitempty"`
	Address    string  `json:"address,omitempty"`
	Menu       string  `json:"menu,omitempty"`
	TotalItems int     `json:"total_items,omitempty"`
	Payment    string  `json:"payment,omitempty"`
	User_id    *string `json:"user_id,omitempty"`
}

func GetAll(db *gorm.DB, orders *[]OrderApp) error {
	err := db.Find(orders)
	if err == nil {
		log.Fatal("error", err)
	}
	return nil
}

func GetByOrderID(db *gorm.DB, order *OrderApp, orderID string) error {
	err := db.Where("order_id=?", orderID).First(order)
	if err == nil {
		log.Fatal("error", err)
	}
	return nil
}

func GetByID(db *gorm.DB, order *OrderApp, ID int) error {
	err := db.Where("id=?", ID).First(order)
	if err == nil {
		log.Fatal("error", err)
	}
	return nil
}

func CreateOrder(db *gorm.DB, order *OrderApp) error {
	err := db.Create(order)
	if err == nil {
		log.Fatal("error", err)
	}
	return nil
}

func DeleteOrder(db *gorm.DB, id int64) error {
	var order OrderApp
	err := db.Where("id=?", id).Delete(order)
	if err == nil {
		log.Fatal("error", err)
	}
	return nil
}

func ModifyOrder(db *gorm.DB, order *OrderApp) error {
	// func UpdateUser(db *gorm.DB, user *User) error {
	err := db.Save(order)
	if err == nil {
		log.Fatal("error", err)
	}
	return nil
}
