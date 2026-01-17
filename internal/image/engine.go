package image

import (
	"fmt"
	"image"
	"image/color"
	"strings"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/goregular"
)

// Engine handles the high-level drawing steps using gg
type Engine struct {
	fontRegular *truetype.Font
	fontBold    *truetype.Font
	fontMono    *truetype.Font
}

// NewEngine creates a new graphics engine with loaded fonts
func NewEngine() (*Engine, error) {
	reg, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}
	bold, err := truetype.Parse(gobold.TTF)
	if err != nil {
		return nil, err
	}
	mono, err := truetype.Parse(gomono.TTF)
	if err != nil {
		return nil, err
	}

	return &Engine{
		fontRegular: reg,
		fontBold:    bold,
		fontMono:    mono,
	}, nil
}

// RenderCard draws a single card based on its type
func (e *Engine) RenderCard(card Card) (image.Image, error) {
	dc := gg.NewContext(Width, Height)

	// Draw Background
	e.drawBackground(dc)

	switch card.Type {
	case CardTypeCover:
		e.renderCover(dc, card)
	case CardTypeIntro:
		e.renderContent(dc, card)
	case CardTypeContent:
		e.renderContent(dc, card)
	case CardTypeCode:
		e.renderCode(dc, card)
	case CardTypeCTA:
		e.renderCTA(dc, card)
	}

	// Draw Footer (Page Number & Branding)
	e.drawFooter(dc, card)

	return dc.Image(), nil
}

func (e *Engine) drawBackground(dc *gg.Context) {
	// Gradient Background
	grad := gg.NewLinearGradient(0, 0, 0, float64(Height))
	grad.AddColorStop(0, ColorBackgroundStart)
	grad.AddColorStop(1, ColorBackgroundEnd)
	dc.SetFillStyle(grad)
	dc.DrawRectangle(0, 0, float64(Width), float64(Height))
	dc.Fill()

	// Subtle header accent
	dc.SetColor(ColorAccent)
	dc.DrawRectangle(0, 0, float64(Width), 20)
	dc.Fill()
}

func (e *Engine) drawFooter(dc *gg.Context, card Card) {
	dc.SetColor(ColorTextSecondary)
	dc.SetFontFace(truetype.NewFace(e.fontBold, &truetype.Options{Size: FontSizeFooter}))

	// Branding Left
	dc.DrawStringAnchored("GO DAILY", 40, Height-40, 0, 0.5)

	// Page Number Right
	if card.Index > 0 && card.TotalSlides > 0 {
		pageStr := fmt.Sprintf("%d / %d", card.Index, card.TotalSlides)
		dc.DrawStringAnchored(pageStr, Width-40, Height-40, 1, 0.5)
	}
}

func (e *Engine) renderCover(dc *gg.Context, card Card) {
	// Title
	dc.SetColor(ColorTextPrimary)
	dc.SetFontFace(truetype.NewFace(e.fontBold, &truetype.Options{Size: FontSizeTitle * 1.2}))
	dc.DrawStringWrapped(strings.ToUpper(card.Title), Width/2, Height/2-100, 0.5, 0.5, Width-Padding*2, 1.2, gg.AlignCenter)

	// Subtitle
	dc.SetColor(ColorAccent)
	dc.SetFontFace(truetype.NewFace(e.fontRegular, &truetype.Options{Size: FontSizeSubtitle}))
	dc.DrawStringAnchored(card.Subtitle, Width/2, Height/2+50, 0.5, 0.5)
}

func (e *Engine) renderContent(dc *gg.Context, card Card) {
	// Header
	dc.SetColor(ColorAccent)
	dc.SetFontFace(truetype.NewFace(e.fontBold, &truetype.Options{Size: FontSizeSubtitle}))
	dc.DrawStringAnchored(card.Title, Padding, Padding*2, 0, 0.5)

	// Body
	dc.SetColor(ColorTextPrimary)
	dc.SetFontFace(truetype.NewFace(e.fontRegular, &truetype.Options{Size: FontSizeBody}))
	dc.DrawStringWrapped(card.Body, Padding, Padding*4, 0, 0, Width-Padding*2, 1.5, gg.AlignLeft)
}

func (e *Engine) renderCode(dc *gg.Context, card Card) {
	// Header
	dc.SetColor(ColorAccent)
	dc.SetFontFace(truetype.NewFace(e.fontBold, &truetype.Options{Size: FontSizeSubtitle}))
	dc.DrawStringAnchored(card.Title, Padding, Padding*2, 0, 0.5)

	// Code Window
	margin := 60.0
	codeY := Padding * 3.5
	codeH := float64(Height) - codeY - Padding*3

	// Window Shadow
	dc.SetColor(color.RGBA{0, 0, 0, 100})
	dc.DrawRoundedRectangle(margin+10, codeY+10, float64(Width)-margin*2, codeH, 20)
	dc.Fill()

	// Window Background
	dc.SetColor(ColorCodeBackground)
	dc.DrawRoundedRectangle(margin, codeY, float64(Width)-margin*2, codeH, 20)
	dc.Fill()

	// Window Controls (Mac Style)
	dc.SetColor(ColorWindowControlRed)
	dc.DrawCircle(margin+30, codeY+30, 8)
	dc.Fill()
	dc.SetColor(ColorWindowControlYellow)
	dc.DrawCircle(margin+60, codeY+30, 8)
	dc.Fill()
	dc.SetColor(ColorWindowControlGreen)
	dc.DrawCircle(margin+90, codeY+30, 8)
	dc.Fill()

	// Simple Syntax Highlighting (Regex-free approach for simplicity)
	dc.SetFontFace(truetype.NewFace(e.fontMono, &truetype.Options{Size: FontSizeCode}))
	e.drawHighlightedText(dc, card.Code, margin+40, codeY+80, Width-margin*2-80)
}

func (e *Engine) drawHighlightedText(dc *gg.Context, text string, x, y, maxWidth float64) {
	lines := strings.Split(text, "\n")
	lineHeight := FontSizeCode * 1.5
	spaceWidth, _ := dc.MeasureString(" ")

	for i, line := range lines {
		words := strings.Split(line, " ")
		curX := x
		isLineComment := false

		for _, word := range words {
			color := ColorTextPrimary

			if isLineComment || strings.HasPrefix(word, "//") {
				isLineComment = true
				color = ColorComment
			} else {
				switch word {
				case "package", "import", "func", "return", "var", "const", "type", "struct", "interface", "map", "if", "else", "for", "range", "go", "defer":
					color = ColorKeyword
				case "string", "int", "bool", "error", "nil", "true", "false", "byte":
					color = ColorFunction
				default:
					if strings.HasPrefix(word, "\"") || strings.HasPrefix(word, "`") {
						color = ColorString
					} else if strings.Contains(word, "(") {
						color = ColorFunction
					}
				}
			}

			dc.SetColor(color)
			dc.DrawString(word, curX, y+float64(i)*lineHeight)
			w, _ := dc.MeasureString(word)
			curX += w + spaceWidth
		}
	}
}

func (e *Engine) renderCTA(dc *gg.Context, card Card) {
	// Centered CTA
	dc.SetColor(ColorTextPrimary)
	dc.SetFontFace(truetype.NewFace(e.fontBold, &truetype.Options{Size: FontSizeTitle}))
	dc.DrawStringWrapped(card.Body, Width/2, Height/2-50, 0.5, 0.5, Width-Padding*2, 1.3, gg.AlignCenter)

	dc.SetColor(ColorAccent)
	dc.SetFontFace(truetype.NewFace(e.fontRegular, &truetype.Options{Size: FontSizeSubtitle}))
	dc.DrawStringAnchored("Follow @go.daily for more!", Width/2, Height/2+100, 0.5, 0.5)
}
