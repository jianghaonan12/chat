package main

import (
	"GINCHAT/router"
	_ "GINCHAT/utils"
)

func main() {
	r := router.Router()
	r.Run()
}
