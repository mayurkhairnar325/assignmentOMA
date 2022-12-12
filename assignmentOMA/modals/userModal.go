package modals

import (
	"github.com/jinzhu/gorm"
	"log"
)

type User struct {
	gorm.Model
	First_name    string  `json:"first_name,omitempty"`
	Last_name     string  `json:"last_name,omitempty"`
	Password      string  `json:"password,omitempty"`
	Email         string  `gorm:"type:varchar(100);unique_index"`
	Avatar        string  `json:"avatar,omitempty"`
	Phone         string  `json:"phone,omitempty"`
	Token         *string `json:"token,omitempty"`
	Refresh_Token *string `json:"refresh_token,omitempty"`
	User_id       string  `json:"user_id,omitempty"`
}

func GetUserByID(db *gorm.DB, user *User, id uint) error {
	err := db.Where("id=?", id).First(user)
	if err == nil {
		log.Fatal("error", err)
	}
	return nil
}

func GetUserByUserID(db *gorm.DB, user *User, userID string) (User, error) {
	err := db.Where("user_id=?", userID).First(user)
	if err == nil {
		log.Fatal("error", err)
	}
	return *user, nil

}

func GetAllUsers(db *gorm.DB, users *[]User) error {
	err := db.Find(users)
	if err == nil {
		log.Fatal("error", err)
	}
	return nil
}

func CreateUser(db *gorm.DB, user *User) error {
	err := db.Create(&user)
	if err == nil {
		log.Fatal("error")
	}
	return nil
}

func UpdateUser(db *gorm.DB, user *User) error {
	err := db.Save(user)
	if err == nil {
		log.Fatal("error", err)
	}
	return nil
}
