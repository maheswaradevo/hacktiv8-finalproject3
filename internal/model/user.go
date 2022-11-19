package model

import "time"

type User struct {
	UserID    uint64    `db:"id"`
	FullName  string    `db:"full_name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
