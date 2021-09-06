package apis

import (
	"apiserver/datastore"
	"apiserver/messages"
	"apiserver/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteMetadata deletes all the metadata for a given title from the datastore.
func DeleteMetadata(c *gin.Context) {
	title := utils.FormatString(c.Param("title"))

	metadataStore := datastore.GetStore()
	if err := metadataStore.DeleteApplication(title); err != nil {
		c.YAML(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": messages.DeleteAppMetadataErr, "error": err})
	} else {
		c.YAML(http.StatusOK, gin.H{"status": http.StatusOK, "message": messages.DeleteAppMetadataSuccess})
	}
}

// DeleteMetadataWithVersion deletes a particular metadata for a given title and version from the datastore.
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
