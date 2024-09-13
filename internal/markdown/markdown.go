package markdown

import (
	"bytes"
	"sync"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
)

type CachedHTML struct {
	HTML      string
	Timestamp time.Time
}

var (
	htmlCache       = make(map[string]CachedHTML)
	cacheMutex      sync.RWMutex
	cacheExpiration = 1 * time.Hour
)

// convertAndSanitize converts markdown text to HTML and sanitizes it.
func convertAndSanitize(markdown string) (string, error) {
	var htmlContent bytes.Buffer
	if err := goldmark.Convert([]byte(markdown), &htmlContent); err != nil {
		return "", err
	}

	p := bluemonday.UGCPolicy()
	sanitized := p.Sanitize(htmlContent.String())

	return sanitized, nil
}

// GetHTMLContent retrieves the HTML content from cache or generates it if not cached.
func GetHTMLContent(markdownContent string) (string, error) {
	cacheMutex.RLock()
	cached, exists := htmlCache[markdownContent]
	cacheMutex.RUnlock()

	if exists && time.Since(cached.Timestamp) < cacheExpiration {
		return cached.HTML, nil
	}

	htmlContent, err := convertAndSanitize(markdownContent)
	if err != nil {
		return "", err
	}

	cacheMutex.Lock()
	htmlCache[markdownContent] = CachedHTML{
		HTML:      htmlContent,
		Timestamp: time.Now(),
	}
	cacheMutex.Unlock()

	return htmlContent, nil
}
