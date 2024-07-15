package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id          string
	Name        string
	Email       string
	Password    string
	Create_time time.Time
}

func NewUser(name string, email string, password string) *User {
	return &User{
		Id:          uuid.New().String(),
		Name:        name,
		Email:       email,
		Password:    password,
		Create_time: time.Now(),
	}
}
