package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/global/utils"
	"github.com/maheswaradevo/hacktiv8-finalproject3/internal/ping"
)

type PingHandler struct {
	r  *gin.RouterGroup
	ps ping.Ping
}

func NewPingHandler(r *gin.RouterGroup, ps ping.Ping) *gin.RouterGroup {
	delivery := PingHandler{
		r:  r,
		ps: ps,
	}
	pingRoute := delivery.r.Group("/ping")
	{
		pingRoute.Handle(http.MethodGet, "", delivery.Ping)
	}
	return pingRoute
}

func (p *PingHandler) Ping(c *gin.Context) {
	res := p.ps.Ping()
	response := utils.NewSuccessResponseWriter(c.Writer, "SUCCESS", http.StatusOK, res)
	c.JSON(http.StatusOK, response)
}
