package main

import (
	"github.com/DianaBurca/app1/utils"
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
	driver.GET("/hello", utils.HelloHandler)
	driver.GET("/test", utils.TestHandler)
	driver.GET("/.well-known/live", utils.Health)
	driver.GET("/.well-known/ready", utils.Health)
	driver.POST("/store-data", utils.StoreData)

	driver.Run()
}
