package main

import (
	"flag"
	"l/internal/web"
	"log"
)

var (
	addr = flag.String("addr", ":5000", "addres server")
)

func main() {
	log.Fatal(web.StartServer(addr))
}
