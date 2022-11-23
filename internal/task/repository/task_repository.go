package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/model"
)

type TaskImplRepo struct {
	db *sql.DB
}

func ProvideTaskRepository(db *sql.DB) *TaskImplRepo {
	return &TaskImplRepo{
		db: db,
	}
}

var (
	CREATE_TASK = "INSERT INTO `tasks`(user_id, category_id, title, description) VALUES (?, ?, ?, ?);"
	CHECK_CATEGORY   = "SELECT id FROM categories;"
)

func (tsk TaskImplRepo) CreateTask(ctx context.Context, data model.Task) (taskID uint64, err error) {
	query := CREATE_TASK
	stmt, err := tsk.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[CreateTask] failed to prepare statement: %v", err)
		return
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, data.UserID, data.CategoryID, data.Title, data.Description)
	if err != nil {
		log.Printf("[CreateComment] failed to insert user to the database: %v", err)
		return
	}

	id, _ := res.LastInsertId()
	taskID = uint64(id)

	return taskID, nil
}

func (tsk TaskImplRepo) CheckTask(ctx context.Context, categoryID uint64) (bool, error) {
	query := CHECK_CATEGORY
	stmt, err := tsk.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[CheckCategory] failed to prepare the statement, err: %v", err)
		return false, err
	}
	rows, err := stmt.QueryContext(ctx, categoryID)
	if err != nil {
		log.Printf("[CheckCategory] failed to query to the database, err: %v", err)
		return false, err
	}
	for rows.Next() {
		return true, nil
	}
	return false, nil
}