package image

import (
	"fmt"
	"strings"

	"github.com/nitin737/GoAutoPosts/internal/model"
)

// GenerateStoryboard creates a sequence of cards from a library model
func GenerateStoryboard(lib *model.Library) Storyboard {
	var cards []Card

	// 1. Cover Card
	cards = append(cards, Card{
		Type:     CardTypeCover,
		Title:    lib.Name,
		Subtitle: formatCategory(lib.Category),
	})

	// 2. Intro/Content Cards
	// Simple logic: if description is very long (> 200 chars), split it?
	// For now, let's just make one content card or multiple if we had sections.
	// We'll trust the description is concise enough or fit it in one for V1.
	cards = append(cards, Card{
		Type:  CardTypeContent,
		Title: "What is it?",
		Body:  lib.Description,
	})

	// 3. Installation / Usage Pattern (Placeholder)
	// If we knew the import path, we could show `go get ...`
	// Assuming URL can give us a hint, e.g. "github.com/..."
	if strings.Contains(lib.URL, "github.com") {
		importPath := strings.TrimPrefix(lib.URL, "https://")
		cards = append(cards, Card{
			Type:     CardTypeCode,
			Title:    "Installation",
			Code:     fmt.Sprintf("go get %s", importPath),
			Language: "bash",
		})
	}

	// 4. CTA
	cards = append(cards, Card{
		Type: CardTypeCTA,
		Body: "Did you learn something?",
	})

	return cards
}

func formatCategory(c string) string {
	if c == "" {
		return "Go Library"
	}
	return fmt.Sprintf("%s Library", strings.Title(c))
}
