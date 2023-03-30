package main

import (
	"Challenge7/configs"
	"Challenge7/routes"
)

func main() {
	configs.StartDBConnection()

	var PORT = ":3000"

	err := routes.ApiInit().Run(PORT)
	if err != nil {
		panic(err)
	}
}
