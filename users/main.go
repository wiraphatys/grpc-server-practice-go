package main

import (
	"user-services/config"
	"user-services/databases"
	"user-services/servers"
)

func main() {
	cfg := config.GetConfig()
	db := databases.NewPostgresDatabase(cfg)

	servers.NewFiberServer(cfg, db.GetDb()).Start()
}
