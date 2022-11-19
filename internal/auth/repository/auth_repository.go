package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/model"
	"github.com/maheswaradevo/hacktiv8-finalproject3/pkg/errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

var (
	SAVE_USER   = "INSERT INTO users(full_name, email, role, password) VALUES (?, ?, ?, ?);"
	CHECK_EMAIL = "SELECT id, full_name, email, password, role FROM users WHERE email = ?;"
)

func (u UserRepository) Save(ctx context.Context, data model.User) (uint64, error) {
	query := SAVE_USER
	res, err := u.db.ExecContext(ctx, query, data.FullName, data.Email, data.Role, data.Password)
	if err != nil && err.(*mysql.MySQLError).Number == 1062 {
		log.Printf("[SaveUser] failed to insert the data to the database, err: %v", err)
		return 0, err
	}
	userID, _ := res.LastInsertId()
	return uint64(userID), nil
}

func (u UserRepository) CheckEmail(ctx context.Context, email string) (*model.User, error) {
	query := CHECK_EMAIL
	rows := u.db.QueryRowContext(ctx, query, email)

	user := &model.User{}

	err := rows.Scan(&user.UserID, &user.Email, &user.FullName, &user.Password, &user.Role)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[CheckEmail] failed to scan the data: %v", err)
		return nil, err
	} else if err == sql.ErrNoRows {
		log.Printf("[CheckEmail] no data existed in the database\n")
		return nil, errors.ErrInvalidResources
	}
	return user, nil
}
