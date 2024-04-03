package api

import (
	"containerhub-api/global"

	"github.com/gin-gonic/gin"
)

func ListImages(c *gin.Context) {
	c.JSON(200, global.Config.Images)
}
