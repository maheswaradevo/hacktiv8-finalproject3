package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	pingHandlerPkg "github.com/maheswaradevo/hacktiv8-finalproject3/internal/ping/handler"
	pingService "github.com/maheswaradevo/hacktiv8-finalproject3/internal/ping/service"

	authHandler "github.com/maheswaradevo/hacktiv8-finalproject3/internal/auth/handler"
	authRepository "github.com/maheswaradevo/hacktiv8-finalproject3/internal/auth/repository"
	authService "github.com/maheswaradevo/hacktiv8-finalproject3/internal/auth/service"
)

func Init(router *gin.Engine, db *sql.DB) {
	api := router.Group("/api/v1")
	{
		InitPingModule(api)

		InitAuthModule(api, db)
	}
}

func InitPingModule(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	pingService := pingService.NewPingService()
	return pingHandlerPkg.NewPingHandler(routerGroup, pingService)
}

func InitAuthModule(routerGroup *gin.RouterGroup, db *sql.DB) *gin.RouterGroup {
	authReposiory := authRepository.NewUserRepository(db)
	authService := authService.NewAuthService(authReposiory)
	return authHandler.NewUserHandler(routerGroup, authService)
}
