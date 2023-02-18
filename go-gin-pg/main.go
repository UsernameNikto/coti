package main

import "C"
import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	Cat.Router(r)
	Dog.Router(r)

	r.Run(addr...:"0.0.0.0:8888")
}
