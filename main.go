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
	port := 0.0.0.0:8080
	r.Run(port)
}
