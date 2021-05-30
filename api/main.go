package main

import (
	"fmt"
	"giraffe/api/controller"
)

func main() {
	startServer()
}

func startServer() {
	e, Giraffe := controller.SetEchoEnv()
	e.Start(fmt.Sprintf(":%s", Giraffe.GetString("port")))
}
