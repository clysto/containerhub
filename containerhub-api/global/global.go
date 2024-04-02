package global

import (
	"containerhub-api/config"
	"containerhub-api/containerhub"

	"gorm.io/gorm"
)

var (
	Config config.Config
	DB     *gorm.DB
	Hub    *containerhub.Hub
)
