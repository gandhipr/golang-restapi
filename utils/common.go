package utils

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

func GetMetadataConfigFromFileURL(c *gin.Context) (Metadata, error) {
	var metadataConfig Metadata

	filePath := c.Param("filepath")
	filePath = FormatString(filePath)

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return metadataConfig, err
	}

	if err = yaml.Unmarshal(bytes, &metadataConfig); err != nil {
		return metadataConfig, err
	}

	return metadataConfig, nil
}

func GetMetadataConfigFromFileBinary(c *gin.Context) (Metadata, error) {
	var metadataConfig Metadata

	f := c.Request.Body
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return metadataConfig, err
	}

	if err = yaml.Unmarshal(bytes, &metadataConfig); err != nil {
		return metadataConfig, err
	}
	return metadataConfig, nil
}

// Replaces "%20" with " " and "+" with "/".
// Currently used to format following fields - title, version, filePath.
func FormatString(strData string) string {
	res := strings.ReplaceAll(strData, "%20", " ")
	res = strings.ReplaceAll(res, "+", "/")
	return res
}
