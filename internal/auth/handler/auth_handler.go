package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/auth"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/dto"
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
