package main

import (
	"meryl/conf"
	"meryl/server"
)

func main() {
	conf.Init()
	r := server.NewRouter()
	r.Run(":3000")
}
