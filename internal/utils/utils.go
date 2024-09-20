package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"homepage/internal/middleware"
	"homepage/internal/models"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"golang.org/x/sync/singleflight"
)

const (
	cacheExpiration = 1 * time.Hour
	maxCacheSize    = 1000
	cleanupInterval = 10 * time.Minute
)

var (
	htmlCache         = make(map[string]models.CachedHTML)
	cacheMutex        sync.RWMutex
	singleflightGroup singleflight.Group
)

func init() {
	go periodicCleanup()
}

func hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func convertAndSanitize(markdown string) (string, error) {
	var htmlContent bytes.Buffer
	if err := goldmark.Convert([]byte(markdown), &htmlContent); err != nil {
		return "", err
	}

	p := bluemonday.UGCPolicy()
	sanitized := p.Sanitize(htmlContent.String())

	return sanitized, nil
}

func GetHTMLContent(markdownContent string) (string, error) {
	cacheKey := hash(markdownContent)

	cacheMutex.RLock()
	cached, exists := htmlCache[cacheKey]
	if exists {
		if time.Since(cached.Timestamp) < cacheExpiration {
			htmlCache[cacheKey] = cached
			cacheMutex.RUnlock()
			return cached.HTML, nil
		}
	}
	cacheMutex.RUnlock()

	htmlContent, err, _ := singleflightGroup.Do(cacheKey, func() (interface{}, error) {
		return convertAndSanitize(markdownContent)
	})

	if err != nil {
		return "", err
	}

	cacheMutex.Lock()
	htmlCache[cacheKey] = models.CachedHTML{
		HTML:       htmlContent.(string),
		Timestamp:  time.Now(),
		LastAccess: time.Now(),
	}
	cacheMutex.Unlock()

	return htmlContent.(string), nil
}

func periodicCleanup() {
	for {
		time.Sleep(cleanupInterval)
		removeOldEntries()
	}
}

func TruncateMarkdown(content string, length int) string {
	if len(content) <= length {
		return content
	}
	return content[:length] + "..."
}

func removeOldEntries() {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if len(htmlCache) <= maxCacheSize {
		return
	}

	var entries []struct {
		key        string
		lastAccess time.Time
	}

	now := time.Now()
	for key, entry := range htmlCache {
		if now.Sub(entry.Timestamp) > cacheExpiration {
			delete(htmlCache, key)
		} else {
			entries = append(entries, struct {
				key        string
				lastAccess time.Time
			}{key, entry.LastAccess})
		}
	}

	if len(htmlCache) > maxCacheSize {
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].lastAccess.Before(entries[j].lastAccess)
		})

		for i := 0; i < len(entries)-maxCacheSize; i++ {
			delete(htmlCache, entries[i].key)
		}
	}

}

func IsUserAdmin(r *http.Request) bool {
	user, _ := middleware.GetUserFromContext(r.Context())
	return user != nil && user.IsAdmin
}
