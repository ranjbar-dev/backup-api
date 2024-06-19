package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ranjbar-dev/backup-api/internal/config"
	"github.com/ranjbar-dev/backup-api/tools/loggertool"
)

func (a *Api) registerMiddlewares() {

	a.hs.GetRouter().Use(func(ctx *gin.Context) {

		// check if client ip is in whitelist
		whitelist := config.Single.WhitelistIp()
		clientIpAddress := ctx.ClientIP()
		for _, ip := range whitelist {

			if clientIpAddress == ip {
				ctx.Next()
				return
			}
		}

		// log
		loggertool.Debug("api", "client ip not in whitelist", clientIpAddress)

		// abort request
		ctx.AbortWithStatus(403)
	})
}
