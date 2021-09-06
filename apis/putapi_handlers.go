package apis

import (
	"apiserver/datastore"
	"apiserver/messages"
	"apiserver/utils"
	"apiserver/validators"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateMetadata updates metadata in the datastore for a given title and version specified using --data-binary.
func UpdateMetadata(c *gin.Context) {
	metadataConfig, err := utils.GetMetadataConfigFromFileBinary(c)
	if err != nil {
		c.YAML(http.StatusBadRequest, gin.H{"status ": http.StatusBadRequest, "errorMessage": messages.ErrGeneratingMetadata, "error": err})
		c.Abort()
		return
	}

	// Validate metadataConfig.
	if valid, err := validators.ValidateInput(metadataConfig); !valid {
		c.YAML(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": messages.InvalidInput, "error": err})
		c.Abort()
		return
	}

	metadataStore := datastore.GetStore()
	if err = metadataStore.UpdateApplicationForVersion(metadataConfig); err != nil {
		c.YAML(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": messages.UpdateAppMetadataErr, "error": err})
	} else {
		c.YAML(http.StatusOK, gin.H{"status": http.StatusOK, "message": messages.UpdateAppMetadataSuccess})
	}
}

// UpdateMetadataFromFileUrl updates metadata in the datastore for a given title and version specified using filepath.
func UpdateMetadataFromFileUrl(c *gin.Context) {
	metadataConfig, err := utils.GetMetadataConfigFromFileURL(c)
	if err != nil {
		c.YAML(http.StatusBadRequest, gin.H{"status ": http.StatusBadRequest, "errorMessage": messages.ErrGeneratingMetadata, "error": err})
		c.Abort()
		return
	}

	// Validate metadataConfig.
	if valid, err := validators.ValidateInput(metadataConfig); !valid {
		c.YAML(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": messages.InvalidInput, "error": err})
		c.Abort()
		return
	}

	metadataStore := datastore.GetStore()
	if err = metadataStore.UpdateApplicationForVersion(metadataConfig); err != nil {
		c.YAML(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": messages.UpdateAppMetadataErr, "error": err})
	} else {
		c.YAML(http.StatusOK, gin.H{"status": http.StatusOK, "message": messages.UpdateAppMetadataSuccess})
	}
}
