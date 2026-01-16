package image

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/nitin737/GoAutoPosts/go-daily/internal/model"
	"golang.org/x/image/font/gofont/goregular"
)

// Generator handles image generation
type Generator struct {
	basePath string
	font     *truetype.Font
}

// NewGenerator creates a new image generator
func NewGenerator(basePath string) (*Generator, error) {
	// Parse embedded font
	font, err := freetype.ParseFont(goregular.TTF)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font: %w", err)
	}

	return &Generator{
		basePath: basePath,
		font:     font,
	}, nil
}

// Generate creates an image for a library
func (g *Generator) Generate(lib *model.Library, outputPath string) error {
	// Load base image
	baseImg, err := g.loadBaseImage()
	if err != nil {
		return fmt.Errorf("failed to load base image: %w", err)
	}

	// Create a new image with the same dimensions
	bounds := baseImg.Bounds()
	img := image.NewRGBA(bounds)
	draw.Draw(img, bounds, baseImg, bounds.Min, draw.Src)

	// Add text overlay
	if err := g.addText(img, lib); err != nil {
		return fmt.Errorf("failed to add text: %w", err)
	}

	// Save the image
	if err := g.saveImage(img, outputPath); err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	return nil
}

func (g *Generator) loadBaseImage() (image.Image, error) {
	file, err := os.Open(g.basePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (g *Generator) addText(img *image.RGBA, lib *model.Library) error {
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(g.font)
	c.SetFontSize(48)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.NewUniform(color.RGBA{255, 255, 255, 255}))

	// Draw library name
	pt := freetype.Pt(50, 100)
	if _, err := c.DrawString(lib.Name, pt); err != nil {
		return err
	}

	// Draw description (smaller font)
	c.SetFontSize(24)
	pt = freetype.Pt(50, 150)
	if _, err := c.DrawString(lib.Description, pt); err != nil {
		return err
	}

	return nil
}

func (g *Generator) saveImage(img *image.RGBA, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
