//go:generate sh -c "test metadata_store_mock.go -nt $GOFILE && exit 0; mockgen -destination=metadata_store_mock.go -package=datastore -source=$GOFILE"

package datastore

import (
	"apiserver/messages"
	"apiserver/utils"
	dsync "github.com/sasha-s/go-deadlock"
)

// In memory datastore.
type Store struct {
	lock         dsync.RWMutex
	appToDetails map[string]map[string]utils.Metadata
}

type StoreIf interface {
	// POST methods.
	AddApplication(utils.Metadata) error

	// GET methods.
	GetAllMetadata() ([]utils.Metadata, error)
	GetApplication(string) ([]utils.Metadata, error)
	GetApplicationWithVersion(string, string) (utils.Metadata, error)

	// PUT methods.
	UpdateApplicationForVersion(utils.Metadata) error

	// DELETE methods.
	DeleteApplication(string) error
	DeleteApplicationWithVersion(string, string) error

	// helper methods.
	isApplicationPresent(string, string) (bool, error)
	isTitlePresent(string) bool
	isVersionPresent(string, string) bool
}

var metadataStore StoreIf = &Store{
	// Each title can have multiple versions.
	// Store info pertaining to each version.
	appToDetails: make(map[string]map[string]utils.Metadata),
}

func GetStore() StoreIf {
	return metadataStore
}

func (s Store) AddApplication(metadata utils.Metadata) error {
	title := metadata.Title
	version := metadata.Version
	s.lock.Lock()
	defer s.lock.Unlock()

	if exists, _ := s.isApplicationPresent(title, version); exists {
		return messages.DuplicateMetadata
	}

	if _, ok := s.appToDetails[title]; !ok {
		s.appToDetails[title] = make(map[string]utils.Metadata)
	}
	s.appToDetails[title][version] = metadata

	return nil
}

func (s Store) UpdateApplicationForVersion(metadata utils.Metadata) error {
	title := metadata.Title
	version := metadata.Version
	s.lock.Lock()
	defer s.lock.Unlock()

	if exists, err := s.isApplicationPresent(title, version); !exists {
		return err
	}
	s.appToDetails[title][version] = metadata

	return nil
}

// title will be give, remove info pertaining to all the versions.
func (s Store) DeleteApplication(title string) error {
	// remove all the versions pertaining to this title
	s.lock.Lock()
	defer s.lock.Unlock()
	// check if the application is present.
	if !s.isTitlePresent(title) {
		return messages.MetadataTitleAbsent
	}

	delete(s.appToDetails, title)
	return nil
}

func (s Store) DeleteApplicationWithVersion(title, version string) error {
	// only remove this version from the title.
	s.lock.Lock()
	defer s.lock.Unlock()

	if exists, err := s.isApplicationPresent(title, version); !exists {
		return err
	}

	delete(s.appToDetails[title], version)
	return nil
}

func (s Store) GetAllMetadata() ([]utils.Metadata, error) {
	var listMetadata []utils.Metadata

	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.appToDetails) == 0 {
		return listMetadata, messages.EmptyStore
	}

	for _, application := range s.appToDetails {
		for _, metadata := range application {
			listMetadata = append(listMetadata, metadata)
		}
	}

	return listMetadata, nil
}

func (s Store) GetApplication(title string) ([]utils.Metadata, error) {
	var listMetadata []utils.Metadata

	s.lock.Lock()
	defer s.lock.Unlock()

	// check if the application is present.
	if !s.isTitlePresent(title) {
		return listMetadata, messages.MetadataTitleAbsent
	}

	for _, metadata := range s.appToDetails[title] {
		listMetadata = append(listMetadata, metadata)
	}

	return listMetadata, nil
}

func (s Store) GetApplicationWithVersion(title, version string) (utils.Metadata, error) {
	var metadata utils.Metadata
	s.lock.Lock()
	defer s.lock.Unlock()

	if exists, err := s.isApplicationPresent(title, version); !exists {
		return metadata, err
	}

	return s.appToDetails[title][version], nil

}

func (s Store) isTitlePresent(title string) bool {
	if _, ok := s.appToDetails[title]; !ok {
		return false
	}
	return true
}

func (s Store) isVersionPresent(title, version string) bool {
	if _, ok := s.appToDetails[title][version]; !ok {
		return false
	}
	return true
}

func (s Store) isApplicationPresent(title, version string) (bool, error) {
	if !s.isTitlePresent(title) {
		return false, messages.MetadataTitleAbsent
	}
	if !s.isVersionPresent(title, version) {
		return false, messages.MetadataVersionAbsent
	}
	return true, nil
}
