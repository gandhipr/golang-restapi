package messages

var (
	AddAppMetadataErr        = "error adding metadata to the store"
	AddAppMetadataSuccess    = "metadata was successfully added to the store"
	ErrGeneratingMetadata    = "error generating metadata config from file binary or filepath"
	InvalidInput             = "incompatible input format"
	UpdateAppMetadataSuccess = "metadata was successfully updated in the store"
	UpdateAppMetadataErr     = "error updating metadata in the store"
	MetadataFound            = "successfully retrieved metadata from teh store"
	MetadataNotFound         = "didn't find relevant metadata in the store"
	DeleteAppMetadataErr     = "error deleting metadata from the store"
	DeleteAppMetadataSuccess = "successfully deleted metadata from the store"
)
