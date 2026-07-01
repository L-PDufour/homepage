package utils

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"homepage/internal/middleware"
	"homepage/internal/models"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/niklasfasching/go-org/org"
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

func convertAndSanitize(content string) (string, error) {
	doc := org.New().Parse(strings.NewReader(content), "")
	html, err := doc.Write(org.NewHTMLWriter())
	if err != nil {
		return "", err
	}

	p := bluemonday.UGCPolicy()
	return p.Sanitize(html), nil
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
	//TODO Find a way to look better
	return content[:length] + " ..."
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

func IsUserAdmin(ctx context.Context) bool {
	user, _ := middleware.GetUserFromContext(ctx)
	return user != nil && user.IsAdmin
}

func StripTags(s string) string {
	var result strings.Builder
	inTag := false
	for _, r := range s {
		switch {
		case r == '<':
			inTag = true
		case r == '>':
			inTag = false
		case !inTag:
			result.WriteRune(r)
		}
	}
	return result.String()
}
