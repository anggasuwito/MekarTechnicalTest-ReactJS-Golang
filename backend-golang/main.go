package main

import (
	"mekarTechnicalTest/app/domains"
	"mekarTechnicalTest/configs"
	"mekarTechnicalTest/configs/router"
)

func main() {
	database, route := configs.InitConfig()
	domains.InitDomain(database, route)
	router.RunServer(route)
}
