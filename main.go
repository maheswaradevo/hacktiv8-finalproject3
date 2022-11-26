package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/global/config"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/global/router"
	"github.com/maheswaradevo/hacktiv8-finalproject3/pkg/database"
)

func main() {
	config.Init()
	cfg := config.GetConfig()
	r := gin.Default()
	db := database.GetDatabase()

	router.Init(r, db)
	PORT := fmt.Sprintf("%s:%s", "localhost", cfg.PORT)
	r.Run(PORT)
}
