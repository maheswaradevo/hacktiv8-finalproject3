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
	CREATE_TASK    = "INSERT INTO `tasks`(user_id, category_id, title, description) VALUES (?, ?, ?, ?);"
	CHECK_CATEGORY = "SELECT id FROM categories;"
	VIEW_TASK      = "SELECT t.id, t.title, t.description, t.status, t.created_at, t.updated_at, u.id, u.email, u.full_name, c.id FROM tasks t INNER JOIN `users` u ON u.id = t.user_id  INNER JOIN `categories` c ON c.id = t.category_id ORDER BY t.created_at DESC;"
	COUNT_TASK     = "SELECT COUNT(*) FROM tasks;"
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

func (tsk TaskImplRepo) ViewTask(ctx context.Context) (model.PeopleTaskJoined, error) {
	query := VIEW_TASK
	stmt, err := tsk.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[ViewPhoto] failed to prepare the statement, err: %v", err)
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("[ViewPhoto] failed to query to the database, err: %v", err)
		return nil, err
	}
	var peopleTask model.PeopleTaskJoined
	for rows.Next() {
		personTask := model.TaskUserJoined{}
		err := rows.Scan(
			&personTask.Task.TaskID,
			&personTask.Task.Title,
			&personTask.Task.Description,
			&personTask.Task.Status,
			&personTask.Task.CategoryID,
			&personTask.Task.UserID,
			&personTask.Task.UpdatedAt,
			&personTask.Task.CreatedAt,
			&personTask.User.Email,
			&personTask.User.FullName,
		)
		if err != nil {
			log.Printf("[ViewComment] failed to scan the data from the database, err: %v", err)
			return nil, err
		}
		peopleTask = append(peopleTask, &personTask)
	}
	return peopleTask, nil
}

func (tsk TaskImplRepo) CountTask(ctx context.Context) (int, error) {
	query := COUNT_TASK
	rows := tsk.db.QueryRowContext(ctx, query)
	var count int
	err := rows.Scan(&count)
	if err != nil {
		log.Printf("[CountTask] failed to scan the data from the database, err: %v", err)
		return 0, err
	}
	return count, nil
}
