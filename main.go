package main

import (
	"github.com/ariyn/Lcd/Controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Controllers.InitController(r)

	r.Run()
}
