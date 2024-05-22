package main

import (
	"auth.services/config"
	"auth.services/servers"
)

func main() {
	cfg := config.GetConfig()

	servers.NewFiberServer(cfg).Start()
}
