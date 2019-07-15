package edgar

import (
	ghttp "net/http"

	"github.com/goedgar/pkg/platform/http"
)

// New creates new EDGAR service
func New() *Service {
	return &Service{
		http: *http.New(),
	}
}

// Service represents EDGAR service
type Service struct {
	http http.Client
}

// FilingsForCIK returns the filings for the company
// with provided CIK number
func (s *Service) FilingsForCIK(cik string) ([]http.Item, error) {
	return s.http.RequestDataItems([]string{cik, "index"})
}

// DocsForFiling accepts company cik number,
// filing name and returns filing documents
func (s *Service) DocsForFiling(cik, name string) ([]http.Item, error) {
	return s.http.RequestDataItems([]string{cik, name, "index"})
}

// DailyIndex accepts the year and returns the daily index
// for provided year
func (s *Service) DailyIndex(year string) ([]http.Item, error) {
	return s.http.RequestDailyIndexItems([]string{year, "index"})
}

// File type contains the document content
// and document type
type File struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

// FileForDoc returns the filing document file
func (s *Service) FileForDoc(cik, filingName, docName string) (*File, error) {
	c, err := s.http.RequestDataContent([]string{cik, filingName, docName})
	if err != nil {
		return nil, err
	}

	return &File{
		Content: string(c),
		Type:    ghttp.DetectContentType(c),
	}, nil
}
