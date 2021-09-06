package apis

import (
	"apiserver/datastore"
	"apiserver/messages"
	"apiserver/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteMetadata(c *gin.Context) {
	title := utils.FormatString(c.Param("title"))

	metadataStore := datastore.GetStore()
	if err := metadataStore.DeleteApplication(title); err != nil {
		c.YAML(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": messages.DeleteAppMetadataErr, "error": err})
	} else {
		c.YAML(http.StatusOK, gin.H{"status": http.StatusOK, "message": messages.DeleteAppMetadataSuccess})
	}
}

func DeleteMetadataWithVersion(c *gin.Context) {
	title := utils.FormatString(c.Param("title"))
	version := utils.FormatString(c.Param("version"))

	metadataStore := datastore.GetStore()
	if err := metadataStore.DeleteApplicationWithVersion(title, version); err != nil {
		c.YAML(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": messages.DeleteAppMetadataErr, "error": err})
	} else {
		c.YAML(http.StatusOK, gin.H{"status": http.StatusOK, "message": messages.DeleteAppMetadataSuccess})
	}
}
