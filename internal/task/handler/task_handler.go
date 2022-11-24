package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/global/middleware"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/global/utils"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/task"
	"github.com/maheswaradevo/hacktiv8-finalproject3/pkg/errors"
)

type TaskHandler struct {
	r  *gin.RouterGroup
	ts task.TaskService
}

func NewTaskHandler(r *gin.RouterGroup, ts task.TaskService) *gin.RouterGroup {
	delivery := TaskHandler{
		r:  r,
		ts: ts,
	}
	taskRoute := delivery.r.Group("/tasks")
	taskProtectedRoute := delivery.r.Group("/tasks", middleware.AuthMiddleware())
	{
		taskProtectedRoute.Handle(http.MethodPost, "/", delivery.createTask)
		taskProtectedRoute.Handle(http.MethodGet, "/", delivery.viewTask)
		taskProtectedRoute.Handle(http.MethodPatch, "/:taskId", delivery.updateTaskStatus)
		taskProtectedRoute.Handle(http.MethodDelete, "/:taskId", delivery.deleteTask)
	}
	return taskRoute
}

func (cmth *TaskHandler) createTask(c *gin.Context) {
	var requestBody dto.CreateTaskRequest
	err := c.BindJSON(&requestBody)
	if err != nil {
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["user_id"].(float64))

	res, err := cmth.ts.CreateTask(c, &requestBody, userID)
	if err != nil {
		log.Printf("[createTask] failed to create user, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusCreated, res)
	c.JSON(http.StatusCreated, response)
}

func (cmth *TaskHandler) viewTask(c *gin.Context) {
	res, err := cmth.ts.ViewTask(c)
	if err != nil {
		log.Printf("[viewTask] failed to view task, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusOK, res)
	c.JSON(http.StatusOK, response)
}

func (cmth *TaskHandler) updateTaskStatus(c *gin.Context) {
	data := dto.EditTaskStatusRequest{}

	err := c.BindJSON(&data)
	if err != nil {
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["user_id"].(float64))
	taskID := c.Param("taskId")
	TaskIDConv, _ := strconv.ParseUint(taskID, 10, 64)

	res, err := cmth.ts.UpdateTaskStatus(c, TaskIDConv, userID, &data)
	if err != nil {
		log.Printf("[updateTask] failed to update task, id: %v, err: %v", TaskIDConv, err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusOK, res)
	c.JSON(http.StatusOK, response)
}

func (cmth *TaskHandler) deleteTask(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["user_id"].(float64))
	taskID := c.Param("taskId")
	taskIDConv, _ := strconv.ParseUint(taskID, 10, 64)

	res, err := cmth.ts.DeleteTask(c, taskIDConv, userID)
	if err != nil {
		log.Printf("[deleteTask] failed to delete task, id: %v, err: %v", taskID, err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusCreated, res)
	c.JSON(http.StatusOK, response)
}
