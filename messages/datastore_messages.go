package messages

import "fmt"

var (
	DuplicateMetadata     = fmt.Errorf("cannot add given metadata to the store because it is already present")
	MetadataTitleAbsent   = fmt.Errorf("metadata with given title is not present in the datastore")
	MetadataVersionAbsent = fmt.Errorf("metadata with given version is not present in the datastore")
	EmptyStore            = fmt.Errorf("there are no metadatas in the store")
)
