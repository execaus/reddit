package models

import (
	"database/sql"
	"time"
)

type Account struct {
	Login      string       `json:"login" db:"login"`
	Password   string       `json:"-" db:"password"`
	Name       string       `json:"name" db:"name"`
	Email      string       `json:"email" db:"email"`
	Role       string       `json:"role" db:"role"`
	CreateDate time.Time    `json:"create_date" db:"create_date"`
	DeadDate   sql.NullTime `json:"dead_date" db:"dead_date"`
}
