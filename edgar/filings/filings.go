package filings

// Resp represents the filings response
// from edgar API
type Resp struct {
	Dir Dir `json:"directory"`
}

// Dir represents the filings directory
// that contains filing items
type Dir struct {
	Items []Item `json:"item"`
}

// Item represents sec filing item
type Item struct {
	LastModified string `json:"last-modified"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Size         string `json:"size"`
}
