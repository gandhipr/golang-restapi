package test

import (
	"apiserver/apis"
	"apiserver/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type outputYaml1 struct {
	Status  int              `yaml:"status,omitempty"`
	Message string           `yaml:"message,omitempty"`
	Data    []utils.Metadata `yaml:"data,omitempty"`
}

// Test_POST_CreateMetadataFromFileUrl tests POST api handler.
func Test_POST_CreateMetadataFromFileUrl(t *testing.T) {
	t.Run("test POST restApi with CreateMetadataFromFileUrl as the handler function should succeed", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/apiserver/metadata/:filepath", apis.CreateMetadataFromFileUrl)
		req, err := http.NewRequest(http.MethodPost, "/apiserver/metadata/..+samples+validinput1-1.0.1.yaml", nil)
		if err != nil {
			t.Fatalf("couldn't create request: %v\n", err)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code == http.StatusCreated {
			t.Logf("Expected to get status %d is same ast %d\n", http.StatusCreated, w.Code)
		} else {
			t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusCreated, w.Code)
		}
	})
}

// Test_GET_GetAllMetadata tests GET api handler.
func Test_GET_GetAllMetadata(t *testing.T) {
	t.Run("test GET restApi with GetAllMetadata as the handler function should succeed", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/apiserver/metadata/", apis.GetAllMetadata)
		req, err := http.NewRequest(http.MethodGet, "/apiserver/metadata/", nil)
		if err != nil {
			t.Fatalf("couldn't create request, err: %v", err)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code == http.StatusOK {
			t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
		} else {
			t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
		}

		filePath := "../samples/validinput1-1.0.1.yaml"
		verifyOutput(t, filePath, w.Body.String())
	})
}

// Test_UPDATE_UpdateMetadataFromFileUrl tests PUT api handler.
func Test_UPDATE_UpdateMetadataFromFileUrl(t *testing.T) {
	t.Run("test PUT restApi with UpdateMetadataFromFileUrl as the handler function should succeed", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PUT("/apiserver/metadata/:filepath", apis.UpdateMetadataFromFileUrl)
		req, err := http.NewRequest(http.MethodPut, "/apiserver/metadata/..+samples+validinput1-1.0.1-update.yaml", nil)
		if err != nil {
			t.Fatalf("couldn't create request: %v\n", err)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code == http.StatusOK {
			t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
		} else {
			t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
		}
	})
}

// Test_DELETE_DeleteMetadataWithVersion tests DELETE api handler.
func Test_DELETE_DeleteMetadataWithVersion(t *testing.T) {
	t.Run("test DELETE restapi wuth DeleteMetadata as the handler function should succeed", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/apiserver/metadata/:title", apis.DeleteMetadata)
		req, err := http.NewRequest(http.MethodDelete, "/apiserver/metadata/App1", nil)
		if err != nil {
			t.Fatalf("couldn't create request, err: %v", err)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code == http.StatusOK {
			t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
		} else {
			t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
		}
	})
}

func verifyOutput(t *testing.T, expectedFilePath, responseBody string) {
	expectedData, err := getMetadata(expectedFilePath)
	if err != nil {
		t.Fatalf("error generating expected metadata, err: %v", err)
	}

	var actualData outputYaml1
	if err := yaml.Unmarshal([]byte(responseBody), &actualData); err != nil {
		t.Fatalf("error unmarshalling the output, err: %v", err)
	}

	assert.Equal(t, len(actualData.Data), 1)
	assert.Equal(t, actualData.Data[0], expectedData)
}

func getMetadata(filePath string) (utils.Metadata, error) {
	var expectedData utils.Metadata
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return expectedData, err
	}

	if err = yaml.Unmarshal(bytes, &expectedData); err != nil {
		return expectedData, err
	}
	return expectedData, nil
}
