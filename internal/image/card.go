package image

import (
	"fmt"
	"strings"

	"github.com/nitin737/GoAutoPosts/internal/model"
)

// CardType represents the type of card in a carousel
type CardType string

const (
	CardTypeCover   CardType = "cover"
	CardTypeIntro   CardType = "intro"
	CardTypeContent CardType = "content"
	CardTypeCode    CardType = "code"
	CardTypeCTA     CardType = "cta"
)

// Card represents a single slide in the carousel
type Card struct {
	Type        CardType
	Title       string
	Subtitle    string
	Body        string
	Code        string
	Index       int // 1-based index
	TotalSlides int
}

// GenerateStoryboard creates a sequence of cards for a library
func GenerateStoryboard(lib *model.Library) []Card {
	var cards []Card

	// 1. Cover Card
	cards = append(cards, Card{
		Type:     CardTypeCover,
		Title:    lib.Name,
		Subtitle: "Go Library Spotlight",
	})

	// 2. Intro Card - What is it?
	cards = append(cards, Card{
		Type:  CardTypeIntro,
		Title: "What is it?",
		Body:  lib.Description,
	})

	// 3. Installation Card (Code)
	installCmd := fmt.Sprintf("go get %s", lib.URL)
	cards = append(cards, Card{
		Type:  CardTypeCode,
		Title: "Installation",
		Code:  installCmd,
	})

	// 4. Category & Tags Card
	if lib.Category != "" || len(lib.Tags) > 0 {
		body := ""
		if lib.Category != "" {
			body += fmt.Sprintf("Category: %s\n\n", lib.Category)
		}
		if len(lib.Tags) > 0 {
			body += fmt.Sprintf("Tags: %s", strings.Join(lib.Tags, ", "))
		}
		cards = append(cards, Card{
			Type:  CardTypeContent,
			Title: "Details",
			Body:  body,
		})
	}

	// 5. Stats Card (if available)
	if lib.Stars > 0 || lib.Author != "" {
		body := ""
		if lib.Stars > 0 {
			body += fmt.Sprintf("⭐ %d stars\n\n", lib.Stars)
		}
		if lib.Author != "" {
			body += fmt.Sprintf("By: %s", lib.Author)
		}
		cards = append(cards, Card{
			Type:  CardTypeContent,
			Title: "Community",
			Body:  body,
		})
	}

	// 6. CTA Card
	cards = append(cards, Card{
		Type: CardTypeCTA,
		Body: "Start using " + lib.Name + " today!",
	})

	return cards
}
