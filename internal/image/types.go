package image

// CardType represents the role of a slide in the carousel.
type CardType int

const (
	CardTypeCover CardType = iota
	CardTypeIntro
	CardTypeContent
	CardTypeCode
	CardTypeCTA
)

// Card represents a single slide in the carousel story.
type Card struct {
	Type     CardType
	Title    string
	Subtitle string // Used in Cover
	Body     string // Used in Intro, Content
	Code     string // Used in Code snippets
	Language string // For syntax highlighting

	// Layout hints
	Highlight bool // Should this card have distinct styling?
}

// Storyboard is the sequence of cards to generation.
type Storyboard []Card
