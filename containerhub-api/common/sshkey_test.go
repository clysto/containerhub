package common

import (
	"containerhub-api/config"
	"containerhub-api/global"
	"testing"
)

func TestRSAKeyGeneration(t *testing.T) {
	config.LoadConfig("../config.toml", &global.Config)
	privateKey, publicKey, cert, hash, err := GenerateRSAKeyPairWithCert("test")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Private key:\n%s\n", string(privateKey))
	t.Logf("Public key:\n%s\n", string(publicKey))
	t.Logf("Certificate:\n%s\n", string(cert))
	t.Logf("Hash: %s\n", hash)
}
