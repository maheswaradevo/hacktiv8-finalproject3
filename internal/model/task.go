package model

import "time"

type Task struct {
	TaskID      uint64    `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Status      bool      `db:"status"`
	UserID      uint64    `db:"user_id"`
	CategoryID  uint64    `db:"category_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type TaskUserJoined struct {
	Task       Task
	Categories Categories
	User       User
}

type PeopleTaskJoined []*TaskUserJoined
type Tasks []*Task
