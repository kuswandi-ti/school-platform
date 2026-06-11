package files

type Classification string

const (
	ClassificationPublic       Classification = "public"
	ClassificationInternal     Classification = "internal"
	ClassificationRestricted   Classification = "restricted"
	ClassificationConfidential Classification = "confidential"
)

type Metadata struct {
	FileID         string
	FoundationID   string
	SchoolID       *string
	OwnerService   string
	EntityType     string
	EntityID       string
	StorageKey     string
	OriginalName   string
	MimeType       string
	SizeBytes      int64
	Classification Classification
}
