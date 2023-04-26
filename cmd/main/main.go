package main

import (
	"l/internal/web"
	"log"
)

func main() {
	log.Fatal(web.StartServer())
}
