package models

import "time"

type User struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Address   string `json:"address" binding:"required"`
	Age       int    `json:"age" binding:"required"`
}

type UserDto struct {
	User      User
	Timestamp time.Time
	TrxId     string
}
