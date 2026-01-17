package image

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"

	"github.com/nitin737/GoAutoPosts/internal/model"
)

// Generator handles image generation
type Generator struct {
	engine *Engine
}

// NewGenerator creates a new image generator
func NewGenerator(basePath string) (*Generator, error) {
	engine, err := NewEngine()
	if err != nil {
		return nil, fmt.Errorf("failed to init graphics engine: %w", err)
	}

	return &Generator{
		engine: engine,
	}, nil
}

// Generate creates a single image (Cover) for a library.
// Maintains backward compatibility.
func (g *Generator) Generate(lib *model.Library, outputPath string) error {
	// Generate just the cover card
	cards := GenerateStoryboard(lib)
	if len(cards) == 0 {
		return fmt.Errorf("no cards generated")
	}

	img, err := g.engine.RenderCard(cards[0])
	if err != nil {
		return fmt.Errorf("failed to render cover: %w", err)
	}

	return g.saveImage(img, outputPath)
}

// GenerateCarousel creates a set of images for a library and returns their paths.
func (g *Generator) GenerateCarousel(lib *model.Library, outputDir string) ([]string, error) {
	cards := GenerateStoryboard(lib)
	var paths []string

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to ensure output dir: %w", err)
	}

	for i := range cards {
		cards[i].Index = i + 1
		cards[i].TotalSlides = len(cards)

		img, err := g.engine.RenderCard(cards[i])
		if err != nil {
			return nil, fmt.Errorf("failed to render card %d: %w", i, err)
		}

		filename := fmt.Sprintf("%s_slide_%d.png", sanitizeFilename(lib.Name), i+1)
		fullPath := filepath.Join(outputDir, filename)

		if err := g.saveImage(img, fullPath); err != nil {
			return nil, fmt.Errorf("failed to save card %d: %w", i, err)
		}

		paths = append(paths, fullPath)
	}

	return paths, nil
}

func (g *Generator) saveImage(img image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}

func sanitizeFilename(name string) string {
	safe := ""
	for _, c := range name {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			safe += string(c)
		} else {
			safe += "_"
		}
	}
	return safe
}
