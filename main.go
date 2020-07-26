package main

import (
	"test/db"
	_ "test/db"
	"test/router"
)
func main() {
	db.Init()
	router:= router.Init()
	router.Run(":8000")
}
