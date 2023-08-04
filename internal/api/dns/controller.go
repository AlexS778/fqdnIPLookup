package dnsapi

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type handler struct {
	DB *sql.DB
}

func RegiserV1Routes(router *gin.Engine, db *sql.DB) {
	h := &handler{DB: db}
	v1 := router.Group("/v1")
	{
		v1.POST("/FQDNToIP", h.FQDNToIPHanlder)
		v1.POST("/IPToFQDN", h.IPToFQDN)
		v1.POST("/whoishere", h.GetWhoIsData)
	}
}
