package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)


func CheckLatency() gin.HandlerFunc {
	fmt.Println("checking user latency")

	return func(context *gin.Context) {
		
		start := time.Now()

		// before request
		context.Next()
		// after request

		latency := time.Since(start)
		latencySeconds := latency.Seconds()

		if latencySeconds > 10 {
			fmt.Println("request is heavy")
		}

	}

}