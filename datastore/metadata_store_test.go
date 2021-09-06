package datastore

import (
	"apiserver/messages"
	"apiserver/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleValidStruct = utils.Metadata{
	Title:   "dummy-app",
	Version: "v1.1",
	Maintainers: []utils.Maintainer{
		{
			Name:  "name1",
			Email: "email1",
		},
		{
			Name:  "name2",
			Email: "email2",
		},
	},
	Company:     "dummyCompany",
	Website:     "dummyWebsite",
	Source:      "dummySource",
	License:     "dummyLicense",
	Description: "dummyDescription",
}

func TestStore_AddMetadataStore(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockMetadataStore := NewMockStoreIf(mockCtrl)

	t.Run("test AddApplicationMetadata should succeed", func(t *testing.T) {
		mockMetadataStore.
			EXPECT().
			AddApplication(sampleValidStruct).
			Return(nil)

		err := mockMetadataStore.AddApplication(sampleValidStruct)
		assert.NoError(t, err)
	})

	t.Run("test AddApplicationMetadata should fail with duplicate data", func(t *testing.T) {
		expectedError := messages.DuplicateMetadata
		mockMetadataStore.
			EXPECT().
			AddApplication(sampleValidStruct).
			Return(expectedError)

		err := mockMetadataStore.AddApplication(sampleValidStruct)
		assert.Equal(t, err, expectedError)
	})
}

func TestStore_UpdateMetadataStore(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockMetadataStore := NewMockStoreIf(mockCtrl)

	t.Run("test UpdateApplicationMetadata should succeed", func(t *testing.T) {
		mockMetadataStore.
			EXPECT().
			UpdateApplicationForVersion(sampleValidStruct).
			Return(nil)

		err := mockMetadataStore.UpdateApplicationForVersion(sampleValidStruct)
		assert.NoError(t, err)
	})

	t.Run("test UpdateApplicationMetadata should fail because given application metadata version is not present", func(t *testing.T) {
		expectedError := messages.MetadataVersionAbsent
		mockMetadataStore.
			EXPECT().
			UpdateApplicationForVersion(sampleValidStruct).
			Return(expectedError)

		err := mockMetadataStore.UpdateApplicationForVersion(sampleValidStruct)
		assert.Equal(t, err, expectedError)
	})
}

func TestStore_GetMetadataStore(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockMetadataStore := NewMockStoreIf(mockCtrl)
	title := sampleValidStruct.Title
	version := sampleValidStruct.Version

	t.Run("test GetAllMetadata should succeed", func(t *testing.T) {
		expectedResult := []utils.Metadata{sampleValidStruct, sampleValidStruct}
		mockMetadataStore.
			EXPECT().
			GetAllMetadata().
			Return(expectedResult, nil)

		actualResult, err := mockMetadataStore.GetAllMetadata()
		assert.NoError(t, err)
		assert.Equal(t, actualResult, expectedResult)
	})

	t.Run("test GetAllMetadata should fail because store is empty", func(t *testing.T) {
		var expectedResult []utils.Metadata
		expectedError := messages.EmptyStore
		mockMetadataStore.
			EXPECT().
			GetAllMetadata().
			Return(expectedResult, expectedError)

		actualResult, err := mockMetadataStore.GetAllMetadata()
		assert.Equal(t, err, expectedError)
		assert.Equal(t, actualResult, expectedResult)
	})

	t.Run("test GetApplicationMetadata for given title should succeed", func(t *testing.T) {
		expectedResult := []utils.Metadata{sampleValidStruct}
		mockMetadataStore.
			EXPECT().
			GetApplication(title).
			Return(expectedResult, nil)

		actualResult, err := mockMetadataStore.GetApplication(title)
		assert.NoError(t, err)
		assert.Equal(t, actualResult, expectedResult)
	})

	t.Run("test GetApplicationMetadata should fail because metadata with given title is not present", func(t *testing.T) {
		var expectedResult []utils.Metadata
		expectedError := messages.MetadataTitleAbsent
		mockMetadataStore.
			EXPECT().
			GetApplication(title).
			Return(expectedResult, expectedError)

		actualResult, err := mockMetadataStore.GetApplication(title)
		assert.Equal(t, err, expectedError)
		assert.Equal(t, actualResult, expectedResult)
	})

	t.Run("test GetApplicationMetadataWithVersion should succeed for given title and version", func(t *testing.T) {
		expectedResult := sampleValidStruct
		mockMetadataStore.
			EXPECT().
			GetApplicationWithVersion(title, version).
			Return(expectedResult, nil)

		actualResult, err := mockMetadataStore.GetApplicationWithVersion(title, version)
		assert.NoError(t, err)
		assert.Equal(t, actualResult, expectedResult)
	})

	t.Run("test GetApplicationMetadataWithVersion should fail for version", func(t *testing.T) {
		var expectedResult utils.Metadata
		expectedError := messages.MetadataVersionAbsent
		mockMetadataStore.
			EXPECT().
			GetApplicationWithVersion(title, version).
			Return(expectedResult, expectedError)

		actualResult, err := mockMetadataStore.GetApplicationWithVersion(title, version)
		assert.Equal(t, err, expectedError)
		assert.Equal(t, actualResult, expectedResult)
	})
}

func TestStore_DeleteMetadataStore(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockMetadataStore := NewMockStoreIf(mockCtrl)
	title := sampleValidStruct.Title
	version := sampleValidStruct.Version

	t.Run("test DeleteApplicationMetadata for given title should succeed", func(t *testing.T) {
		mockMetadataStore.
			EXPECT().
			DeleteApplication(title).
			Return(nil)

		err := mockMetadataStore.DeleteApplication(title)
		assert.NoError(t, err)
	})

	t.Run("test DeleteApplicationMetadata should fail because metadata with given title is not present", func(t *testing.T) {
		expectedError := messages.MetadataTitleAbsent
		mockMetadataStore.
			EXPECT().
			DeleteApplication(title).
			Return(expectedError)

		err := mockMetadataStore.DeleteApplication(title)
		assert.Equal(t, err, expectedError)
	})

	t.Run("test DeleteApplicationMetadataWithVersion should succeed for given title and version", func(t *testing.T) {
		mockMetadataStore.
			EXPECT().
			DeleteApplicationWithVersion(title, version).
			Return(nil)

		err := mockMetadataStore.DeleteApplicationWithVersion(title, version)
		assert.NoError(t, err)
	})

	t.Run("test DeleteApplicationMetadataWithVersion should fail for version", func(t *testing.T) {
		expectedError := messages.MetadataVersionAbsent
		mockMetadataStore.
			EXPECT().
			DeleteApplicationWithVersion(title, version).
			Return(expectedError)

		err := mockMetadataStore.DeleteApplicationWithVersion(title, version)
		assert.Equal(t, err, expectedError)
	})
}
