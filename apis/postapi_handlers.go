package apis

import (
	"apiserver/datastore"
	"apiserver/messages"
	"apiserver/utils"
	"apiserver/validators"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateMetadata stores metadata specified in --data-binary in the datastore if not present.
func CreateMetadata(c *gin.Context) {
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
	err = metadataStore.AddApplication(metadataConfig)
	if err != nil {
		c.YAML(http.StatusAlreadyReported, gin.H{"status": http.StatusAlreadyReported, "message": messages.AddAppMetadataErr, "error": err})
	} else {
		c.YAML(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": messages.AddAppMetadataSuccess})
	}
}

// CreateMetadataFromFileUrl stores metadata specified in the filepath in the datastore if not present.
func CreateMetadataFromFileUrl(c *gin.Context) {
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
	err = metadataStore.AddApplication(metadataConfig)
	if err != nil {
		c.YAML(http.StatusAlreadyReported, gin.H{"status": http.StatusAlreadyReported, "message": messages.AddAppMetadataErr, "error": err})
	} else {
		c.YAML(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": messages.AddAppMetadataSuccess})
	}
}
