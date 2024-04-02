package common

import (
	"containerhub-api/global"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"math"

	"golang.org/x/crypto/ssh"
)

func GeneratePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func EncodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}
	privatePEM := pem.EncodeToMemory(&privBlock)
	return privatePEM
}

func GeneratePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}
	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)
	return pubKeyBytes, nil
}

func SignSSHKey(publicKey ssh.PublicKey, user string) (*ssh.Certificate, error) {
	permissions := ssh.Permissions{
		CriticalOptions: map[string]string{},
		Extensions: map[string]string{
			"permit-X11-forwarding":   "",
			"permit-agent-forwarding": "",
			"permit-port-forwarding":  "",
			"permit-pty":              "",
		},
	}
	cert := ssh.Certificate{
		Key:             publicKey,
		KeyId:           "user@containerhub",
		ValidPrincipals: []string{user},
		Serial:          0,
		CertType:        ssh.UserCert,
		Permissions:     permissions,
		ValidAfter:      0,
		ValidBefore:     math.MaxUint64,
	}
	signer, err := ssh.ParsePrivateKey(global.Config.SSH.CAPrivkeyPEM)
	if err != nil {
		return nil, err
	}
	err = cert.SignCert(rand.Reader, signer)
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

func GenerateRSAKeyPairWithCert(user string) ([]byte, []byte, []byte, string, error) {
	privateKey, err := GeneratePrivateKey(2048)
	if err != nil {
		return nil, nil, nil, "", err
	}
	privateKeyPEM := EncodePrivateKeyToPEM(privateKey)
	publicKeyPEM, err := GeneratePublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, nil, "", err
	}
	sshPublicKey, err := ssh.NewPublicKey(privateKey.Public())
	if err != nil {
		return nil, nil, nil, "", err
	}
	cert, err := SignSSHKey(sshPublicKey, user)
	if err != nil {
		return nil, nil, nil, "", err
	}
	certBytes := ssh.MarshalAuthorizedKey(cert)
	sha256sum := sha256.Sum256(sshPublicKey.Marshal())
	return privateKeyPEM, publicKeyPEM, certBytes, hex.EncodeToString(sha256sum[:]), nil
}
