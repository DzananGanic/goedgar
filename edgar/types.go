package edgar

// Type represents item type in EDGAR
type Type string

const (
	// FolderType represents folder item
	FolderType = "folder.gif"

	// TextType represents text item
	TextType = "text.gif"

	// ImageType represents image item
	ImageType = "image2.gif"
)

// Resp represents the edgar response
// from edgar API
type Resp struct {
	Dir Dir `json:"directory"`
}

// Dir represents the edgar directory
// that contains filing items
type Dir struct {
	Items []Item `json:"item"`
}

// Item represents sec edgar item
type Item struct {
	LastModified string `json:"last-modified"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Size         string `json:"size"`
}
