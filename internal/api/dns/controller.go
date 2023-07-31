package dnsapi

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegiserV1Routes(router *gin.Engine, db *gorm.DB) {
	h := &handler{DB: db}
	v1 := router.Group("/v1")
	{
		v1.POST("/FQDNToIP", h.FQDNToIPHanlder)
		v1.POST("/IPToFQDN", h.IPToFQDN)
		v1.POST("/whoishere", h.GetWhoIsData)
	}
}
