package model

import "time"

type Categories struct {
	CategoryID uint64    `db:"id"`
	Type       string    `db:"type"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type CategoriesUserJoined struct {
	Categories Categories
	Task       Task
	User       User
}

type PeopleCategoriesJoined []*CategoriesUserJoined
type AllCategories []*Categories
