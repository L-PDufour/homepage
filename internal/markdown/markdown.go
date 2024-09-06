package markdown

import (
	"bytes"
	"log"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
)

type MarkdownConverter interface {
	ConvertAndSanitize(markdown string) (string, error)
}

type MarkdownService struct {
	logger *log.Logger
}

func NewMarkdownService(logger *log.Logger) *MarkdownService {
	return &MarkdownService{
		logger: logger,
	}
}

func (s *MarkdownService) ConvertAndSanitize(markdown string) (string, error) {
	var htmlContent bytes.Buffer
	if err := goldmark.Convert([]byte(markdown), &htmlContent); err != nil {
		s.logger.Printf("Error converting Markdown to HTML: %v", err)
		return "", err
	}
	s.logger.Println("Markdown converted to HTML successfully")

	p := bluemonday.UGCPolicy()
	sanitized := p.Sanitize(htmlContent.String())
	s.logger.Println("HTML sanitized successfully")

	return sanitized, nil
}
