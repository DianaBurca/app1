package main

import (
	"github.com/DianaBurca/app1/internal"
	"github.com/gin-gonic/gin"
)

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func main() {
	driver := gin.Default()
	driver.GET("/hello", internal.HelloHandler)
	driver.GET("/.well-known/live", internal.Health)
	driver.GET("/.well-known/ready", internal.Health)
	driver.POST("/store-data", internal.StoreData)

	driver.Run()
}
