package main

import (
	"apiserver/apis"
	"github.com/gin-gonic/gin"
)

// In memory object to store application metadata.
// A single application can have multiples version
func main() {
	router := gin.Default()

	api := router.Group("apiserver/metadata")
	{
		api.POST("/", apis.CreateMetadata)
		api.POST("/:filepath", apis.CreateMetadataFromFileUrl)

		api.PUT("/", apis.UpdateMetadata)
		api.PUT("/:filepath", apis.UpdateMetadataFromFileUrl)

		api.GET("/", apis.GetAllMetadata)
		api.GET("/:title", apis.GetMetadata)
		api.GET("/:title/:version", apis.GetMetadataForVersion)

		api.DELETE("/:title", apis.DeleteMetadata)
		api.DELETE("/:title/:version", apis.DeleteMetadataWithVersion)
	}

	listenPort := "8080"
	// Listen and Server on the LocalHost:Port
	router.Run(":" + listenPort)
}
