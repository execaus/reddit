package models

import "time"

type Account struct {
	Login      string    `json:"login" db:"login"`
	Password   string    `json:"-" db:"password"`
	Name       string    `json:"name" db:"name"`
	Email      string    `json:"email" db:"email"`
	Role       string    `json:"role" db:"role"`
	CreateDate time.Time `json:"create_date" db:"create_date"`
	DeadDate   time.Time `json:"dead_date" db:"dead_date"`
}
