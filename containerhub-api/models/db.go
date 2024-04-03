package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"-"`
	Admin    bool   `json:"admin"`
}

type SSHKey struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Hash          string `json:"hash"`
	Certificate   string `json:"certificate"`
	PublicKey     string `json:"publicKey"`
	PrivateKey    string `json:"privateKey"`
	ContainerID   string `json:"containerID" gorm:"uniqueIndex"`
	ContainerName string `json:"containerName"`
	ContainerHost string `json:"containerHost"`
	UserID        uint   `json:"userID"`
	User          User   `json:"user" gorm:"foreignKey:UserID"`
}
