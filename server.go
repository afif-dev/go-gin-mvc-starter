package main

import (
	"gin-mvc/console"
	"gin-mvc/router"
)

func main() {

	// run console
	console.InitConsole()
	
	// gin app
	r := router.InitRouter()
	r.Run()
}
