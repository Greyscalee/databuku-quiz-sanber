package main

import (
	"data-buku/config"
	"data-buku/routers"
)

func main() {
	config.LoadEnv(".env")

	db := config.ConnectDB()
	defer db.Close()

	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}
