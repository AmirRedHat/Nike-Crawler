package main 

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"io"
	"encoding/json"
	"log"

	"localPackage/lib"
	
)

func returnData(dataBytes io.Reader) map[string]interface{} {
	data := make(map[string]interface{})
	posted_data, err := ioutil.ReadAll(dataBytes)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(posted_data, &data)
	return data
}

func Server() {
	fmt.Println("server is running")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// use middlewares
	router.Use(CheckLatency())

	router.POST("/nike", NikeProductHandler)

	router.Run("0.0.0.0:8080")

}

func NikeProductHandler(context *gin.Context) {
	data := returnData(context.Request.Body)

	nikeURL := data["url"].(string)
	fmt.Println("Crawling ", nikeURL)
	product := lib.Crawl(nikeURL)

	context.JSON(200, product)
}