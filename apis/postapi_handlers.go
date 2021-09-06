package apis

import (
	"apiserver/datastore"
	"apiserver/messages"
	"apiserver/utils"
	"apiserver/validators"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateMetadata(c *gin.Context) {
	ctx := utils.GinContext{C: c}
	metadataConfig, err := ctx.GetMetadataConfigFromFileBinary()
	if err != nil {
		c.YAML(http.StatusBadRequest, gin.H{"status ": http.StatusBadRequest, "errorMessage": messages.ErrGeneratingMetadata, "error": err})
		c.Abort()
		return
	}
	// Validate metadataConfig
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

func CreateMetadataFromFileUrl(c *gin.Context) {
	ctx := utils.GinContext{C: c}

	metadataConfig, err := ctx.GetMetadataConfigFromFileURL()
	if err != nil {
		c.YAML(http.StatusBadRequest, gin.H{"status ": http.StatusBadRequest, "errorMessage": messages.ErrGeneratingMetadata, "error": err})
		c.Abort()
		return
	}
	// Validate metadataConfig
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
