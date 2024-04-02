package main

import (
	"containerhub-api/api"
	"containerhub-api/config"
	"containerhub-api/containerhub"
	"containerhub-api/global"
	"containerhub-api/middleware"
	"containerhub-api/models"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configFile := flag.String("c", "config.toml", "config file")
	flag.Parse()
	config.LoadConfig(*configFile, &global.Config)
	var err error
	global.Hub, err = containerhub.ConnectHub()
	if err != nil {
		panic(err)
	}
	global.DB, err = gorm.Open(sqlite.Open(global.Config.DatabasePath), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	global.DB.AutoMigrate(models.User{}, models.SSHKey{})
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/login", api.Login)
		v1.POST("/signup", api.Signup)
		v1.GET("/containers", middleware.Auth, api.ListContainers)
		v1.POST("/containers", middleware.Auth, api.CreateContainer)
		v1.POST("/containers/start", middleware.Auth, api.StartContainer)
		v1.POST("/containers/stop", middleware.Auth, api.StopContainer)
		v1.POST("/containers/destroy", middleware.Auth, api.DestroyContainer)
		v1.POST("/containers/keys", middleware.Auth, api.GenerateSSHKey)
		v1.GET("/containers/keys", middleware.Auth, api.ListSSHKeys)
		v1.DELETE("/containers/keys", middleware.Auth, api.DeleteSSHKey)
	}
	router.Run(fmt.Sprintf("%s:%d", global.Config.Host, global.Config.Port))
}
