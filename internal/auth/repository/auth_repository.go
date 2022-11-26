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
	SAVE_USER       = "INSERT INTO users(full_name, email, role, password) VALUES (?, ?, ?, ?);"
	CHECK_EMAIL     = "SELECT id, full_name, email, password, role FROM users WHERE email = ?;"
	UPDATE_ACCOUNT  = "UPDATE users SET full_name = ?, email = ? WHERE id = ?;"
	DELETE_ACCOUNT  = "DELETE FROM users WHERE id = ?;"
	FIND_USER_BY_ID = "SELECT id FROM users WHERE id = ?;"
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

func (u UserRepository) UpdateAccount(ctx context.Context, data model.User, userID uint64) error {
	query := UPDATE_ACCOUNT

	_, err := u.db.ExecContext(ctx, query, data.FullName, data.Email, userID)
	if err != nil && err.(*mysql.MySQLError).Number == 1062 {
		log.Printf("[UpdateAccount] failed to insert data to the database, err : %v", err)
		return err
	}
	return nil
}

func (u UserRepository) DeleteAccount(ctx context.Context, userID uint64) (string, error) {
	query := DELETE_ACCOUNT

	_, err := u.db.ExecContext(ctx, query, userID)
	if err != nil {
		log.Printf("[DeleteAccount] failed to delete data from the database, err : %v", err)
		return "", err
	}
	msg := "Your account has been successfully deleted"
	return msg, nil
}

func (u UserRepository) FindUserByID(ctx context.Context, userID uint64) (bool, error) {
	query := FIND_USER_BY_ID

	res, err := u.db.QueryContext(ctx, query, userID)
	if err != nil {
		log.Printf("[FindUserByID] failed to query to the database, err: %v", err)
		return false, err
	}

	for res.Next() {
		return true, nil
	}
	return false, err
}
