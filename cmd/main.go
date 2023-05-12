package main

import (
	"go-mssql/cmd/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Start(); err != nil {
		log.Fatalf("Error start server : " + err.Error())
	}
}
