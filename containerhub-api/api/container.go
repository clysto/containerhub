package api

import (
	"containerhub-api/common"
	"containerhub-api/containerhub"
	"containerhub-api/global"
	"containerhub-api/models"
	"errors"

	"github.com/gin-gonic/gin"
)

func ListContainers(ctx *gin.Context) {
	user := ctx.Keys["user"].(models.Claims)
	containers, err := global.Hub.ListContainers(user.Username)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, containers)
}

func CreateContainer(ctx *gin.Context) {
	user := ctx.Keys["user"].(models.Claims)
	param := models.ContainerCreateParam{}
	if err := ctx.BindJSON(&param); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := global.Hub.CreateContainer(param.Image, user.Username, global.Config.SSH.CAPubkeyPEM)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"id": id})
}

func StartContainer(ctx *gin.Context) {
	user := ctx.Keys["user"].(models.Claims)
	id := ctx.Query("id")
	container, err := global.Hub.GetContainer(id)
	if err != nil {
		var code int
		if errors.Is(err, containerhub.ContainerNotFoundError{}) {
			code = 404
		} else {
			code = 500
		}
		ctx.JSON(code, gin.H{"error": err.Error()})
	}
	if container.Labels["containerhub-user"] != user.Username {
		ctx.JSON(403, gin.H{"error": "Permission denied"})
		return
	}
	err = global.Hub.StartContainer(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Container started"})
}

func StopContainer(ctx *gin.Context) {
	user := ctx.Keys["user"].(models.Claims)
	id := ctx.Query("id")
	container, err := global.Hub.GetContainer(id)
	if err != nil {
		var code int
		if errors.Is(err, containerhub.ContainerNotFoundError{}) {
			code = 404
		} else {
			code = 500
		}
		ctx.JSON(code, gin.H{"error": err.Error()})
	}
	if container.Labels["containerhub-user"] != user.Username {
		ctx.JSON(403, gin.H{"error": "Permission denied"})
		return
	}
	err = global.Hub.StopContainer(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Container stopped"})
}

func DestroyContainer(ctx *gin.Context) {
	user := ctx.Keys["user"].(models.Claims)
	id := ctx.Query("id")
	container, err := global.Hub.GetContainer(id)
	if err != nil {
		var code int
		if errors.Is(err, containerhub.ContainerNotFoundError{}) {
			code = 404
		} else {
			code = 500
		}
		ctx.JSON(code, gin.H{"error": err.Error()})
	}
	if container.Labels["containerhub-user"] != user.Username {
		ctx.JSON(403, gin.H{"error": "Permission denied"})
		return
	}
	err = global.Hub.DestroyContainer(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Container destroyed"})
}

func GenerateSSHKey(ctx *gin.Context) {
	containerID := ctx.Query("id")
	user := ctx.Keys["user"].(models.Claims)
	container, err := global.Hub.GetContainer(containerID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if container.Labels["containerhub-user"] != user.Username {
		ctx.JSON(403, gin.H{"error": "Permission denied"})
		return
	}

	var key models.SSHKey

	privateKey, publicKey, cert, hash, err := common.GenerateRSAKeyPairWithCert(user.Username)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	key.PrivateKey = string(privateKey)
	key.PublicKey = string(publicKey)
	key.Certificate = string(cert)
	key.Hash = hash
	key.ConatainerID = containerID
	key.UserID = user.UserID
	key.ContainerHost = container.Labels["containerhub-name"]
	if err := global.DB.Create(&key).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, key)
}

func ListSSHKeys(ctx *gin.Context) {
	user := ctx.Keys["user"].(models.Claims)
	keys := []models.SSHKey{}
	if err := global.DB.Where("user_id = ?", user.UserID).Find(&keys).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, keys)
}

func DeleteSSHKey(ctx *gin.Context) {
	user := ctx.Keys["user"].(models.Claims)
	id := ctx.Query("id")
	key := models.SSHKey{}
	if err := global.DB.Where("id = ? AND user_id = ?", id, user.UserID).First(&key).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if err := global.DB.Delete(&key).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "SSH key deleted"})
}
