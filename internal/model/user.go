package model

import (
	_ "gorm.io/gorm"
	"time"
)

type User struct {
	UserId    int       `db:"user_id" json:"id"`
	Email     string    `db:"email" json:"email" gorm:"unique"`
	Password  string    `db:"password" json:"password"`
	Role      string    `db:"role" json:"role"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
