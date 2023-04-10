package main

import (
	"log"
	"venn-auth-api/internal/api"
)

func main() {
	err := api.New().Run()
	if err != nil {
		log.Fatal(err)
	}
}
