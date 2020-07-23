package main

import (
	_ "test/db"
	"test/router"
)
func main() {
	router:= router.Init()
	router.Run(":8000")
}
