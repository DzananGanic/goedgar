package http

const (
	// FolderType represents folder item
	FolderType = "folder.gif"

	// TextType represents text item
	TextType = "text.gif"

	// ImageType represents image item
	ImageType = "image2.gif"
)

// Response represents the edgar response
// from edgar API
type Response struct {
	Directory Directory `json:"directory"`
}

// Directory represents the edgar directory
// that contains items
type Directory struct {
	Items []Item `json:"item"`
}

// Item represents sec edgar item
type Item struct {
	LastModified string `json:"last-modified"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Size         string `json:"size"`
}
