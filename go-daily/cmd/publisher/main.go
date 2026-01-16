package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nitin737/GoAutoPosts/go-daily/internal/config"
	"github.com/nitin737/GoAutoPosts/go-daily/internal/hashtag"
	"github.com/nitin737/GoAutoPosts/go-daily/internal/image"
	"github.com/nitin737/GoAutoPosts/go-daily/internal/instagram"
	"github.com/nitin737/GoAutoPosts/go-daily/internal/logger"
	"github.com/nitin737/GoAutoPosts/go-daily/internal/model"
	"github.com/nitin737/GoAutoPosts/go-daily/internal/selector"
	"github.com/nitin737/GoAutoPosts/go-daily/internal/store"
	"github.com/nitin737/GoAutoPosts/go-daily/internal/template"
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

	instagramClient := instagram.NewClient(cfg.InstagramAccessToken, cfg.InstagramAccountID)
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

	// Step 4: Generate image
	logger.Info("Generating image...")
	imagePath := fmt.Sprintf("/tmp/go-daily-%s.png", library.Name)
	if err := imageGen.Generate(library, imagePath); err != nil {
		logger.Error("Failed to generate image", "error", err)
		os.Exit(1)
	}
	logger.Info("Image generated", "path", imagePath)

	// Step 5: Publish to Instagram
	logger.Info("Publishing to Instagram...")
	postID, err := publisher.PublishPost(imagePath, caption)
	if err != nil {
		logger.Error("Failed to publish to Instagram", "error", err)
		os.Exit(1)
	}
	logger.Info("Successfully published to Instagram", "postID", postID)

	// Step 6: Update posted history
	logger.Info("Updating posted history...")
	postedLibrary := &model.PostedLibrary{
		Library:   *library,
		PostedAt:  time.Now(),
		PostID:    postID,
		ImagePath: imagePath,
	}

	if err := store.Save(postedLibrary); err != nil {
		logger.Error("Failed to save posted history", "error", err)
		// Don't exit here - the post was successful
	}

	logger.Info("Daily publisher completed successfully", "library", library.Name, "postID", postID)
}
