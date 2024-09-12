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

func ConvertAndSanitize(markdown string) (string, error) {
	var htmlContent bytes.Buffer
	if err := goldmark.Convert([]byte(markdown), &htmlContent); err != nil {
		return "", err
	}

	p := bluemonday.UGCPolicy()
	sanitized := p.Sanitize(htmlContent.String())

	return sanitized, nil
}
