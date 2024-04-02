package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Host         string `toml:"host"`
	Port         int    `toml:"port"`
	JWTSecret    string `toml:"jwt_secret"`
	DatabasePath string `toml:"db_path"`
	SSH          struct {
		CAPrivkeyFile string `toml:"ca_privkey_file"`
		CAPubkeyFile  string `toml:"ca_pubkey_file"`
		CAPrivkeyPEM  []byte
		CAPubkeyPEM   []byte
	} `toml:"ssh"`
}

func LoadConfig(file string, config *Config) {
	config.Host = "127.0.0.1"
	config.Port = 8080
	_, err := toml.DecodeFile(file, config)
	if err != nil {
		panic(err)
	}

	// Load CA private key
	caPrivKeyPEM, err := os.ReadFile(config.SSH.CAPrivkeyFile)
	if err != nil {
		panic(err)
	}
	config.SSH.CAPrivkeyPEM = caPrivKeyPEM

	// Load CA public key
	caPubKeyPEM, err := os.ReadFile(config.SSH.CAPubkeyFile)
	if err != nil {
		panic(err)
	}
	config.SSH.CAPubkeyPEM = caPubKeyPEM
}
