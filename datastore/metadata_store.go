//go:generate sh -c "test metadata_store_mock.go -nt $GOFILE && exit 0; mockgen -destination=metadata_store_mock.go -package=datastore -source=$GOFILE"

package datastore

import (
	"apiserver/messages"
	"apiserver/utils"
	dsync "github.com/sasha-s/go-deadlock"
)

// Store represents datatype for in-memory datastore.
type Store struct {
	lock         dsync.RWMutex
	appToDetails map[string]map[string]utils.Metadata
}

// StoreIf represents store interface. Used for mocking.
type StoreIf interface {
	// POST methods.
	AddApplication(utils.Metadata) error

	// PUT methods.
	UpdateApplicationForVersion(utils.Metadata) error

	// GET methods.
	GetAllMetadata() ([]utils.Metadata, error)
	GetApplication(string) ([]utils.Metadata, error)
	GetApplicationWithVersion(string, string) (utils.Metadata, error)

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
	// Store metadata information pertaining to each version.
	appToDetails: make(map[string]map[string]utils.Metadata),
}

// GetStore returns the singleton instance of the metadata store.
func GetStore() StoreIf {
	return metadataStore
}

// AddApplication stores metadata in the store.
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

// UpdateApplicationForVersion updates title,version - metadata mapping in the store.
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

// GetAllMetadata returns all the application metadata present in the store.
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

// GetApplication returns all the metadata for a given title.
func (s Store) GetApplication(title string) ([]utils.Metadata, error) {
	var listMetadata []utils.Metadata

	s.lock.Lock()
	defer s.lock.Unlock()

	if !s.isTitlePresent(title) {
		return listMetadata, messages.MetadataTitleAbsent
	}

	for _, metadata := range s.appToDetails[title] {
		listMetadata = append(listMetadata, metadata)
	}

	return listMetadata, nil
}

// GetApplicationWithVersion return metadata pertaining to given title and version.
func (s Store) GetApplicationWithVersion(title, version string) (utils.Metadata, error) {
	var metadata utils.Metadata
	s.lock.Lock()
	defer s.lock.Unlock()

	if exists, err := s.isApplicationPresent(title, version); !exists {
		return metadata, err
	}

	return s.appToDetails[title][version], nil

}

// DeleteApplication deletes all the metadata from the store for a given title.
func (s Store) DeleteApplication(title string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if !s.isTitlePresent(title) {
		return messages.MetadataTitleAbsent
	}

	delete(s.appToDetails, title)
	return nil
}

// DeleteApplicationWithVersion delete a particular metadata for a given title and version.
func (s Store) DeleteApplicationWithVersion(title, version string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if exists, err := s.isApplicationPresent(title, version); !exists {
		return err
	}

	delete(s.appToDetails[title], version)
	return nil
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
