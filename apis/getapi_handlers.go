package apis

import (
	"apiserver/datastore"
	"apiserver/messages"
	"apiserver/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllMetadata(c *gin.Context) {
	metadataStore := datastore.GetStore()

	data, err := metadataStore.GetAllMetadata()
	if err != nil {
		c.YAML(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": messages.MetadataNotFound, "error": err})
	} else {
		c.YAML(http.StatusOK, gin.H{"status": http.StatusOK, "message": messages.MetadataFound, "data": data})
	}
}

func GetMetadata(c *gin.Context) {
	title := utils.FormatString(c.Param("title"))
	metadataStore := datastore.GetStore()

	data, err := metadataStore.GetApplication(title)
	if err != nil {
		c.YAML(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": messages.MetadataNotFound, "error": err})
	} else {
		c.YAML(http.StatusOK, gin.H{"status": http.StatusOK, "message": messages.MetadataFound, "data": data})
	}
}

func GetMetadataForVersion(c *gin.Context) {
	title := utils.FormatString(c.Param("title"))
	version := utils.FormatString(c.Param("version"))

	metadataStore := datastore.GetStore()
	data, err := metadataStore.GetApplicationWithVersion(title, version)
	if err != nil {
		c.YAML(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": messages.MetadataNotFound, "error": err})
	} else {
		c.YAML(http.StatusOK, gin.H{"status": http.StatusOK, "message": messages.MetadataFound, "data": data})
	}
}
