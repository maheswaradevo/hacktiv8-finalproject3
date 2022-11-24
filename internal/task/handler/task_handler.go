package handler

import (
	"log"
	"net/http"

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
		log.Printf("[createComment] failed to create user, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusCreated, res)
	c.JSON(http.StatusCreated, response)
}

func (cmth *TaskHandler) viewTask(c *gin.Context) {
	res, err := cmth.ts.ViewTask(c)
	if err != nil {
		log.Printf("[viewComment] failed to view comment, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusOK, res)
	c.JSON(http.StatusOK, response)
}