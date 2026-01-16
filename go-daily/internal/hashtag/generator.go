package hashtag

import (
	"strings"

	"github.com/nitin737/GoAutoPosts/go-daily/internal/model"
)

// Generator handles hashtag generation
type Generator struct {
	baseHashtags []string
}

// NewGenerator creates a new hashtag generator
func NewGenerator() *Generator {
	return &Generator{
		baseHashtags: []string{
			"golang",
			"go",
			"programming",
			"coding",
			"developer",
			"software",
			"opensource",
			"tech",
		},
	}
}

// Generate creates hashtags for a library
func (g *Generator) Generate(lib *model.Library) []string {
	hashtags := make([]string, 0)

	// Add base hashtags
	hashtags = append(hashtags, g.baseHashtags...)

	// Add category-specific hashtag
	if lib.Category != "" {
		hashtags = append(hashtags, normalizeHashtag(lib.Category))
	}

	// Add library tags
	for _, tag := range lib.Tags {
		hashtags = append(hashtags, normalizeHashtag(tag))
	}

	// Limit to 30 hashtags (Instagram limit)
	if len(hashtags) > 30 {
		hashtags = hashtags[:30]
	}

	return hashtags
}

// normalizeHashtag converts a string to a valid hashtag format
func normalizeHashtag(s string) string {
	// Remove spaces and special characters
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, "_", "")
	s = strings.ToLower(s)
	return s
}
