package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	pingHandlerPkg "github.com/maheswaradevo/hacktiv8-finalproject3/internal/ping/handler"
	pingService "github.com/maheswaradevo/hacktiv8-finalproject3/internal/ping/service"

	authHandler "github.com/maheswaradevo/hacktiv8-finalproject3/internal/auth/handler"
	authRepository "github.com/maheswaradevo/hacktiv8-finalproject3/internal/auth/repository"
	authService "github.com/maheswaradevo/hacktiv8-finalproject3/internal/auth/service"

	taskHandler "github.com/maheswaradevo/hacktiv8-finalproject3/internal/task/handler"
	taskRepository "github.com/maheswaradevo/hacktiv8-finalproject3/internal/task/repository"
	taskService "github.com/maheswaradevo/hacktiv8-finalproject3/internal/task/service"

	categoriesHandler "github.com/maheswaradevo/hacktiv8-finalproject3/internal/categories/handler"
	categoriesRepository "github.com/maheswaradevo/hacktiv8-finalproject3/internal/categories/repository"
	categoriesService "github.com/maheswaradevo/hacktiv8-finalproject3/internal/categories/service"
)

func Init(router *gin.Engine, db *sql.DB) {
	api := router.Group("/api/v1")
	{
		InitPingModule(api)

		InitAuthModule(api, db)

		InitTaskModule(api, db)

		InitCategoriesModule(api, db)
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

func InitTaskModule(routerGroup *gin.RouterGroup, db *sql.DB) *gin.RouterGroup {
	taskRepository := taskRepository.ProvideTaskRepository(db)
	taskService := taskService.ProvideTaskService(taskRepository)
	return taskHandler.NewTaskHandler(routerGroup, taskService)
}

func InitCategoriesModule(routerGroup *gin.RouterGroup, db *sql.DB) *gin.RouterGroup {
	categoriesRepository := categoriesRepository.ProvideCategoriesRepository(db)
	categoriesService := categoriesService.ProvideCategoriesService(categoriesRepository)
	return categoriesHandler.NewCategoriesHandler(routerGroup, categoriesService)
}
