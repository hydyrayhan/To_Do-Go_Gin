package models

import "time"

type User struct {
	Id         string    `json:"id" db:"id"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"-" db:"password"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
}
