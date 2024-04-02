package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"github.com/tg123/sshpiper/libplugin"
	"github.com/urfave/cli/v2"
)

var db *sql.DB

func queryConnectInfo(hash string) ([]byte, []byte, string, error) {
	query := "SELECT container_host, private_key, certificate FROM ssh_keys WHERE hash = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, nil, "", err
	}
	defer stmt.Close()
	row := stmt.QueryRow(hash)
	var privateKey []byte
	var certificate []byte
	var host string
	err = row.Scan(&host, &privateKey, &certificate)
	if err != nil {
		return nil, nil, "", err
	}
	return privateKey, certificate, host, nil
}

func main() {
	libplugin.CreateAndRunPluginTemplate(&libplugin.PluginTemplate{
		Name:  "sshmux",
		Usage: "sshpiperd sshmux plugin, redirect to a docker container",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "db",
				Usage:    "database file",
				EnvVars:  []string{"SSHPIPERD_SSHMUX_DB"},
				Required: true,
			},
		},
		CreateConfig: func(c *cli.Context) (*libplugin.SshPiperPluginConfig, error) {
			dbFile := c.String("db")

			var err error
			db, err = sql.Open("sqlite3", dbFile)
			if err != nil {
				return nil, err
			}

			return &libplugin.SshPiperPluginConfig{
				PublicKeyCallback: func(conn libplugin.ConnMetadata, key []byte) (*libplugin.Upstream, error) {
					// user := conn.User()
					hash := sha256.Sum256(key)
					hashStr := hex.EncodeToString(hash[:])
					log.Info("pubkey hash ", hashStr)
					privkey, capublickey, host, err := queryConnectInfo(hashStr)
					if err != nil {
						return nil, err
					}
					log.Info("routing to ", host+":22")
					return &libplugin.Upstream{
						Host:          host,
						Port:          22,
						IgnoreHostKey: true,
						UserName:      "user",
						Auth:          libplugin.CreatePrivateKeyAuth(privkey, capublickey),
					}, nil
				},
			}, nil
		},
	})
}
