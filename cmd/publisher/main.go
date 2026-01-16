package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nitin737/GoAutoPosts/internal/config"
	"github.com/nitin737/GoAutoPosts/internal/hashtag"
	"github.com/nitin737/GoAutoPosts/internal/image"
	"github.com/nitin737/GoAutoPosts/internal/instagram"
	"github.com/nitin737/GoAutoPosts/internal/logger"
	"github.com/nitin737/GoAutoPosts/internal/model"
	"github.com/nitin737/GoAutoPosts/internal/selector"
	"github.com/nitin737/GoAutoPosts/internal/store"
	"github.com/nitin737/GoAutoPosts/internal/template"
)

func main() {
	// Initialize logger
	logger := logger.NewLogger()
	logger.Info("Starting daily publisher...")

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	// Initialize components
	selector := selector.NewLibrarySelector(cfg.LibrariesPath, cfg.PostedPath)
	hashtagGen := hashtag.NewGenerator()
	renderer, err := template.NewRenderer()
	if err != nil {
		logger.Error("Failed to initialize template renderer", "error", err)
		os.Exit(1)
	}

	imageGen, err := image.NewGenerator(cfg.ImageBasePath)
	if err != nil {
		logger.Error("Failed to initialize image generator", "error", err)
		os.Exit(1)
	}

	instagramClient := instagram.NewClient(cfg.InstagramAccessToken, cfg.InstagramAccountID, cfg.GraphAPIURL)
	publisher := instagram.NewPublisher(instagramClient)
	store := store.NewJSONStore(cfg.PostedPath)

	// Step 1: Select a random library
	logger.Info("Selecting random library...")
	library, err := selector.SelectRandom()
	if err != nil {
		logger.Error("Failed to select library", "error", err)
		os.Exit(1)
	}
	logger.Info("Selected library", "name", library.Name, "category", library.Category)

	// Step 2: Generate hashtags
	logger.Info("Generating hashtags...")
	hashtags := hashtagGen.Generate(library)
	logger.Info("Generated hashtags", "count", len(hashtags))

	// Step 3: Render caption
	logger.Info("Rendering caption...")
	caption, err := renderer.RenderCaption(library, hashtags)
	if err != nil {
		logger.Error("Failed to render caption", "error", err)
		os.Exit(1)
	}

	// Step 4: Generate images (Carousel)
	logger.Info("Generating carousel images...")
	// Use a clean directory
	outputDir := fmt.Sprintf("/tmp/go-daily-%s-%d", library.Name, time.Now().Unix())
	imagePaths, err := imageGen.GenerateCarousel(library, outputDir)
	if err != nil {
		logger.Error("Failed to generate carousel", "error", err)
		os.Exit(1)
	}
	logger.Info("Carousel generated", "count", len(imagePaths), "dir", outputDir)

	// Step 5: Publish to Instagram
	logger.Info("Publishing to Instagram...")
	postID, err := publisher.PublishCarousel(imagePaths, caption)
	if err != nil {
		logger.Error("Failed to publish to Instagram", "error", err)
		os.Exit(1)
	}
	logger.Info("Successfully published to Instagram", "postID", postID)

	// Step 6: Update posted history
	logger.Info("Updating posted history...")
	postedLibrary := &model.PostedLibrary{
		Library:  *library,
		PostedAt: time.Now(),
		PostID:   postID,
		// Store the first image path as reference or comma separated
		ImagePath: imagePaths[0],
	}

	if err := store.Save(postedLibrary); err != nil {
		logger.Error("Failed to save posted history", "error", err)
		// Don't exit here - the post was successful
	}

	logger.Info("Daily publisher completed successfully", "library", library.Name, "postID", postID)
}
