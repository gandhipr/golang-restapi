package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

func (ctx GinContext) GetMetadataConfigFromFileURL() (Metadata, error) {
	c := ctx.C
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

func (ctx GinContext) GetMetadataConfigFromFileBinary() (Metadata, error) {
	c := ctx.C
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
// Currently used by get and delete rest apis - title and version, filePath - above.
func FormatString(strData string) string {
	res := strings.ReplaceAll(strData, "%20", " ")
	res = strings.ReplaceAll(res, "+", "/")
	return res
}
