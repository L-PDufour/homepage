package markdown

import (
	"bytes"
	"log"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
)

// MarkdownConverter interface defines the method for converting Markdown to HTML
type MarkdownConverter interface {
	ConvertAndSanitize(markdown string) (string, error)
}

// MarkdownService handles Markdown conversion and sanitization.
type MarkdownService struct {
	logger *log.Logger
}

// NewMarkdownService creates a new instance of MarkdownService.
func NewMarkdownService(logger *log.Logger) MarkdownConverter {
	return &MarkdownService{
		logger: logger,
	}
}

// ConvertAndSanitize converts Markdown to HTML and sanitizes it.
func (s *MarkdownService) ConvertAndSanitize(markdown string) (string, error) {
	var htmlContent bytes.Buffer
	if err := goldmark.Convert([]byte(markdown), &htmlContent); err != nil {
		s.logger.Printf("Error converting Markdown to HTML: %v", err)
		return "", err
	}
	s.logger.Println("Markdown converted to HTML successfully")

	// Sanitize HTML
	p := bluemonday.UGCPolicy()
	sanitized := p.Sanitize(htmlContent.String())
	s.logger.Println("HTML sanitized successfully")

	return sanitized, nil
}
