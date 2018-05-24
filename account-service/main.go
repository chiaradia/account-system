package main

import (
	"fmt"
	"github.com/chiaradia/account-system/account-service/services"
	"github.com/chiaradia/account-system/account-service/dbclient"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	services.StartWebServer("6767") // NEW
}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient() {
	services.DBClient = &dbclient.BoltClient{}
	services.DBClient.OpenBoltDb()
	services.DBClient.Seed()
}
