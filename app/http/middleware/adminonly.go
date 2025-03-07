package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jadevelopmentgrp/Tickets-Import-API/config"
	"github.com/jadevelopmentgrp/Tickets-Import-API/utils"
)

func AdminOnly(ctx *gin.Context) {
	userId := ctx.Keys["userid"].(uint64)

	if !utils.Contains(config.Conf.Admins, userId) {
		ctx.JSON(401, utils.ErrorStr("Unauthorized"))
		ctx.Abort()
		return
	}
}
