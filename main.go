package main

import "golangSimpleCrud/server"

func main() {
	r := server.GetRouter()
	// TODO: Place config parser here so that the host port and the db can be set
	r.Run()
}
