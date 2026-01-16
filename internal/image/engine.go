package image

import (
	"image"
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

	// Add watermark/branding
	e.drawBranding(dc)

	switch card.Type {
	case CardTypeCover:
		e.renderCover(dc, card)
	case CardTypeIntro:
		e.renderContent(dc, card) // Intro is similar to content for now
	case CardTypeContent:
		e.renderContent(dc, card)
	case CardTypeCode:
		e.renderCode(dc, card)
	case CardTypeCTA:
		e.renderCTA(dc, card)
	}

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
}

func (e *Engine) drawBranding(dc *gg.Context) {
	dc.SetColor(ColorTextSecondary)
	dc.SetFontFace(truetype.NewFace(e.fontBold, &truetype.Options{Size: 24}))
	dc.DrawStringAnchored("GO DAILY", Width-40, 40, 1, 0.5)
}

func (e *Engine) renderCover(dc *gg.Context, card Card) {
	// Title
	dc.SetColor(ColorTextPrimary)
	// Larger title for cover
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
	dc.DrawStringAnchored(card.Title, Padding, Padding*1.5, 0, 0.5)

	// Body
	dc.SetColor(ColorTextPrimary)
	dc.SetFontFace(truetype.NewFace(e.fontRegular, &truetype.Options{Size: FontSizeBody}))

	// Calculate text position - simple flow
	// TODO: Better text flow
	dc.DrawStringWrapped(card.Body, Padding, Padding*3, 0, 0, Width-Padding*2, 1.5, gg.AlignLeft)
}

func (e *Engine) renderCode(dc *gg.Context, card Card) {
	// Header
	dc.SetColor(ColorAccent)
	dc.SetFontFace(truetype.NewFace(e.fontBold, &truetype.Options{Size: FontSizeSubtitle}))
	dc.DrawStringAnchored(card.Title, Padding, Padding*1.5, 0, 0.5)

	// Code Block Background
	margin := 60.0
	codeY := Padding * 3.0
	codeH := float64(Height) - codeY - Padding*2

	dc.SetColor(ColorCodeBackground)
	dc.DrawRoundedRectangle(margin, codeY, float64(Width)-margin*2, codeH, 20)
	dc.Fill()

	// Code Text
	dc.SetColor(ColorTextPrimary)
	dc.SetFontFace(truetype.NewFace(e.fontMono, &truetype.Options{Size: FontSizeCode}))

	// Very simple code positioning
	// TODO: Syntax highlighting
	dc.DrawStringWrapped(card.Code, margin+40, codeY+40, 0, 0, Width-margin*2-80, 1.4, gg.AlignLeft)
}

func (e *Engine) renderCTA(dc *gg.Context, card Card) {
	dc.SetColor(ColorTextPrimary)
	dc.SetFontFace(truetype.NewFace(e.fontBold, &truetype.Options{Size: FontSizeTitle}))
	dc.DrawStringAnchored(card.Body, Width/2, Height/2, 0.5, 0.5)

	dc.SetColor(ColorAccent)
	dc.SetFontFace(truetype.NewFace(e.fontRegular, &truetype.Options{Size: FontSizeSubtitle}))
	dc.DrawStringAnchored("Follow @go.daily for more!", Width/2, Height/2+100, 0.5, 0.5)
}
