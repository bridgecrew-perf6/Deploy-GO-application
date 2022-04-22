package main

import (
	"fmt"
	"os"
	"time"

	"backend/routes"
	"backend/storage/database"
)

var err error

func init() {
}

func main() {
	bootstrap()
}

func bootstrap() {
	defer restart()
	database.ConnectDB()
	err = routes.App.Listen(fmt.Sprintf(":%s", os.Getenv("3101")))
	if err != nil {
		panic(err)
	}
}

func restart() {
	if err != nil {
		fmt.Println("waiting for 5 second to restart app", err)
		time.Sleep(time.Second * 5)
		bootstrap()
	}
}
