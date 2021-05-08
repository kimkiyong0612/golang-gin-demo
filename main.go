package main

import (
	_ "fmt"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Static("/static", "./static")
	engine.Run(":3000")
}
