package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/auth"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/global/middleware"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/global/utils"
	"github.com/maheswaradevo/hacktiv8-finalproject3/pkg/errors"
)

type UserHandler struct {
	r  *gin.RouterGroup
	us auth.UserService
}

func NewUserHandler(r *gin.RouterGroup, us auth.UserService) *gin.RouterGroup {
	delivery := UserHandler{
		r:  r,
		us: us,
	}
	userRoute := delivery.r.Group("/users")
	{
		userRoute.Handle(http.MethodPost, "/register", delivery.register)
		userRoute.Handle(http.MethodPost, "/login", delivery.login)
	}
	userProtectedRoute := delivery.r.Group("/users", middleware.AuthMiddleware())
	{
		userProtectedRoute.Handle(http.MethodPut, "/update-account", delivery.updateAccount)
		userProtectedRoute.Handle(http.MethodDelete, "/delete-account", delivery.deleteAccount)
	}
	return userRoute
}

func (u *UserHandler) register(c *gin.Context) {
	registerRequest := &dto.UserRegisterRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(registerRequest)
	if err != nil {
		log.Printf("[register] failed to parse json data: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Code, errResponse)
		return
	}

	res, err := u.us.RegisterUser(c, registerRequest)
	if err != nil {
		log.Printf("[registerUser] failed to register a user: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "Registrasi Akun Berhasil", http.StatusCreated, res)
	c.JSON(http.StatusCreated, response)
}

func (u *UserHandler) login(c *gin.Context) {
	loginRequest := &dto.UserSignInRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(loginRequest)
	if err != nil {
		log.Printf("[login] failed to parse json data: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Code, errResponse)
		return
	}

	res, err := u.us.Login(c, loginRequest)
	if err != nil {
		log.Printf("[login] user failed to login, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}

	response := utils.NewSuccessResponseWriter(c.Writer, "Login Sukses", http.StatusOK, res)
	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) updateAccount(c *gin.Context) {
	updateRequest := &dto.UserUpdateAccountRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(updateRequest)
	if err != nil {
		log.Printf("[updateAccount] failed to parse json data: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}

	userLoginData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userLoginData["user_id"].(float64))

	res, err := u.us.UpdateAccount(c, userID, updateRequest)
	if err != nil {
		log.Printf("[updateAccount] user failed to update account, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "Update Account Sukses", http.StatusOK, res)
	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) deleteAccount(c *gin.Context) {
	userLoginData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userLoginData["user_id"].(float64))

	res, err := u.us.DeleteAccount(c, userID)
	if err != nil {
		log.Printf("[deleteAccount] failed to delete account, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, "Delete Account Sukses", http.StatusOK, res)
	c.JSON(http.StatusOK, response)
}
