package main

import (
	"github.com/gin-gonic/gin"
	"github.com/cacahootie/echgo/lib"
)

func main() {
    r := gin.Default()
    r.GET("/", echgo.Echgo)
    r.GET("/echo", echgo.Echgo)
    r.POST("/", echgo.Echgo)
    r.POST("/echo", echgo.Echgo)
    r.PUT("/", echgo.Echgo)
    r.PUT("/echo", echgo.Echgo)
    r.PATCH("/", echgo.Echgo)
    r.PATCH("/echo", echgo.Echgo)
    r.DELETE("/", echgo.Echgo)
    r.DELETE("/echo", echgo.Echgo)
    r.Run() // listen and server on 0.0.0.0:8080
}