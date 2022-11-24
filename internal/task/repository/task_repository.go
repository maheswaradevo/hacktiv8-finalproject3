package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/model"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/dto"
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
	CREATE_TASK        = "INSERT INTO `tasks`(user_id, category_id, title, description) VALUES (?, ?, ?, ?);"
	CHECK_CATEGORY     = "SELECT id FROM categories;"
	VIEW_TASK          = "SELECT t.id, t.title, t.description, t.status, t.category_id, t.user_id, t.created_at, t.updated_at, u.email, u.full_name FROM tasks t JOIN `users` u ON u.id = t.user_id  ORDER BY t.created_at DESC;"
	COUNT_TASK         = "SELECT COUNT(*) FROM tasks;"
	CHECK_TASK         = "SELECT id FROM tasks WHERE id = ? AND user_id = ?;"
	UPDATE_TASK_STATUS = "UPDATE tasks SET status = ? WHERE id = ?;"
	DELETE_TASK        = "DELETE FROM tasks WHERE id = ? AND user_id = ?;"
	GET_TASK_BY_ID     = "SELECT t.id, t.title, t.description, t.status, t.category_id, t.updated_at FROM `tasks` t WHERE t.id = ? AND t.user_id = ?;"
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
		log.Printf("[CreateTask] failed to insert user to the database: %v", err)
		return
	}

	id, _ := res.LastInsertId()
	taskID = uint64(id)

	return taskID, nil
}

func (tsk TaskImplRepo) CheckCategory(ctx context.Context, categoryID uint64) (bool, error) {
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
		log.Printf("[ViewTask] failed to prepare the statement, err: %v", err)
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("[ViewTask] failed to query to the database, err: %v", err)
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
			log.Printf("[ViewTask] failed to scan the data from the database, err: %v", err)
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

func (tsk TaskImplRepo) CheckTask(ctx context.Context, taskID uint64, userID uint64) (bool, error) {
	query := CHECK_TASK
	stmt, err := tsk.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[CheckTask] failed to prepare the statement, err: %v", err)
		return false, err
	}
	rows, err := stmt.QueryContext(ctx, taskID, userID)
	if err != nil {
		log.Printf("[CheckTask] failed to query to the database, err: %v", err)
		return false, err
	}
	for rows.Next() {
		return true, nil
	}
	return false, nil
}

func (tsk TaskImplRepo) UpdateTaskStatus(ctx context.Context, reqData model.TaskUserJoined, taskID uint64, userID uint64) error {
	query := UPDATE_TASK_STATUS

	stmt, err := tsk.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[UpdateTaskStatus] failed to prepare the statement, err: %v", err)
		return err
	}
	_, err = stmt.ExecContext(ctx, reqData.Task.Status, taskID)
	if err != nil {
		log.Printf("[UpdateTaskStatus] failed to store data to the database, err: %v", err)
		return err
	}
	return nil
}

func (tsk TaskImplRepo) GetTaskByID(ctx context.Context, taskID uint64, userID uint64) (*dto.EditTaskStatusResponse, error) {
	query := GET_TASK_BY_ID
	stmt, err := tsk.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[GetTaskByID] failed to prepare the statement, err: %v", err)
		return nil, err
	}
	rows := stmt.QueryRowContext(ctx, taskID, userID)
	if err != nil {
		log.Printf("[GetTaskByID] failed to query to the database, err: %v", err)
		return nil, err
	}
	personTask := model.TaskUserJoined{}
	err = rows.Scan(
		&personTask.Task.TaskID,
		&personTask.Task.Title,
		&personTask.Task.Description,
		&personTask.Task.Status,
		&personTask.Task.CategoryID,
		&personTask.Task.UpdatedAt,
	)
	if err != nil {
		log.Printf("[GetTaskByID] failed to scan the data from the database, err: %v", err)
		return nil, err
	}
	return dto.NewEditTaskResponse(personTask.Task, userID), err
}

func (tsk TaskImplRepo) DeleteTask(ctx context.Context, taskID uint64, userID uint64) error {
	query := DELETE_TASK

	stmt, err := tsk.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[DeleteTask] failed to prepare the statement, err: %v", err)
		return err
	}

	_, err = stmt.QueryContext(ctx, taskID, userID)
	if err != nil {
		log.Printf("[DeleteTask] failed to delete the task, err: %v", err)
		return err
	}
	return nil
}
