package apis

import (
	"apiserver/datastore"
	"apiserver/messages"
	"apiserver/utils"
	"apiserver/validators"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateMetadata(c *gin.Context) {
	ctx := utils.GinContext{C: c}
	metadataConfig, err := ctx.GetMetadataConfigFromFileBinary()
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

func UpdateMetadataFromFileUrl(c *gin.Context) {
	ctx := utils.GinContext{C: c}
	metadataConfig, err := ctx.GetMetadataConfigFromFileURL()
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
