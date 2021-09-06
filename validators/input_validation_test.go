package validators

import (
	"apiserver/messages"
	"apiserver/utils"
	"github.com/mohae/deepcopy"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var (
	validInput        = true
	invalidInput      = false
	sampleValidStruct = utils.Metadata{
		Title:   "dummy-app",
		Version: "v1.1",
		Maintainers: []utils.Maintainer{
			{
				Name:  "name1",
				Email: "email1@gmail.com",
			},
			{
				Name:  "name2",
				Email: "email2@yahoo.com",
			},
		},
		Company:     "dummyCompany",
		Website:     "dummyWebsite",
		Source:      "dummySource",
		License:     "dummyLicense",
		Description: "dummyDescription",
	}
)

func TestValidInputs(t *testing.T) {
	t.Run("input validation should succeed", func(t *testing.T) {
		valid, expectedErrorMap := ValidateInput(sampleValidStruct)
		assert.Equal(t, valid, validInput)
		assert.Equal(t, len(expectedErrorMap), 0)
	})
}

func TestInvalidInputs(t *testing.T) {

	emptyCompany := deepcopy.Copy(sampleValidStruct).(utils.Metadata)
	emptyCompany.Company = ""
	log.Println(emptyCompany)

	emptyName := deepcopy.Copy(sampleValidStruct).(utils.Metadata)
	emptyName.Maintainers[0].Name = ""

	invalidEmail := deepcopy.Copy(sampleValidStruct).(utils.Metadata)
	invalidEmail.Maintainers[1].Email = "invalidEmail"

	invalidTitleVersion := deepcopy.Copy(sampleValidStruct).(utils.Metadata)
	invalidTitleVersion.Title = "%invalid"
	invalidTitleVersion.Version = "invalid,"

	fields := []string{"company", "description", "license", "maintainers", "source", "title", "version", "website"}
	errorMap := make(map[string][]string)
	for _, field := range fields {
		errorMap[field] = append(errorMap[field], messages.RequiredError)
	}

	testCases := []struct {
		name             string
		input            utils.Metadata
		expectedErrorMap map[string][]string
	}{
		{
			name:             " input validation should fail with empty data struct",
			input:            utils.Metadata{},
			expectedErrorMap: errorMap,
		},
		{
			name:  "input validation should fail with empty company name",
			input: emptyCompany,
			expectedErrorMap: map[string][]string{
				"company": {messages.RequiredError},
			},
		},
		{
			name:  "input validation should fail with empty name",
			input: emptyName,
			expectedErrorMap: map[string][]string{
				"name": {messages.RequiredError},
			},
		},
		{
			name:  "input validation should fail with invalid email",
			input: invalidEmail,
			expectedErrorMap: map[string][]string{
				"email": {messages.EmailError},
			},
		},
		{
			name:  "input validation should fail with invalid title and version",
			input: invalidTitleVersion,
			expectedErrorMap: map[string][]string{
				"title":   {messages.TitleError},
				"version": {messages.VersionError},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			valid, actualErrorMap := ValidateInput(testCase.input)
			assert.Equal(t, invalidInput, valid)
			assert.Equal(t, testCase.expectedErrorMap, actualErrorMap)
		})
	}
}
