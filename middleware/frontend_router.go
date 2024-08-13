package middleware

import (
	"embed"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func NewFrontedRouter(fsEmbed embed.FS) *gin.Engine {
	router := gin.Default()

	router.Use(static.Serve("/", static.EmbedFolder(fsEmbed, "web/build")))

	return router
}
