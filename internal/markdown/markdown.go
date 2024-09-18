package markdown

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"golang.org/x/sync/singleflight"
)

type CachedHTML struct {
	HTML       string
	Timestamp  time.Time
	LastAccess time.Time
}

const (
	cacheExpiration = 1 * time.Hour
	maxCacheSize    = 1000 // Adjust based on your needs
	cleanupInterval = 10 * time.Minute
)

var (
	htmlCache         = make(map[string]CachedHTML)
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
			// Update last access time
			cached.LastAccess = time.Now()
			htmlCache[cacheKey] = cached
			cacheMutex.RUnlock()
			log.Println("Cache hit: Returning cached HTML content")
			return cached.HTML, nil
		}
	}
	cacheMutex.RUnlock()

	// Use singleflight to prevent multiple goroutines from generating the same content
	htmlContent, err, _ := singleflightGroup.Do(cacheKey, func() (interface{}, error) {
		log.Println("Cache miss or expired: Generating new HTML content")
		return convertAndSanitize(markdownContent)
	})

	if err != nil {
		log.Printf("Error converting markdown to HTML: %v", err)
		return "", err
	}

	// Update cache
	cacheMutex.Lock()
	htmlCache[cacheKey] = CachedHTML{
		HTML:       htmlContent.(string),
		Timestamp:  time.Now(),
		LastAccess: time.Now(),
	}
	cacheMutex.Unlock()

	log.Println("New HTML content cached")
	return htmlContent.(string), nil
}
func periodicCleanup() {
	for {
		time.Sleep(cleanupInterval)
		removeOldEntries()
	}
}

// removeOldEntries removes expired and least recently used entries
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

	// If still over maxCacheSize, remove least recently used
	if len(htmlCache) > maxCacheSize {
		// Sort entries by last access time
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].lastAccess.Before(entries[j].lastAccess)
		})

		// Remove oldest entries until we're at maxCacheSize
		for i := 0; i < len(entries)-maxCacheSize; i++ {
			delete(htmlCache, entries[i].key)
		}
	}

	log.Printf("Cache cleanup complete. Current cache size: %d", len(htmlCache))
}
