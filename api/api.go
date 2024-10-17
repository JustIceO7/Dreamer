package api

import (
	"context"

	"github.com/Strum355/log"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

func InitAPI(s *discordgo.Session) {
	// Creates the Gin Router Object
	router := gin.New()

	// MiddleWare
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		ctx := context.WithValue(context.Background(), log.Key, log.Fields{
			"time":    param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			"method":  param.Method,
			"path":    param.Path,
			"status":  param.StatusCode,
			"latency": param.Latency,
			"agent":   param.Request.UserAgent(),
		})
		log.WithContext(ctx).Info("invoked request")
		return ""
	}))
	router.Use(gin.Recovery())
}
