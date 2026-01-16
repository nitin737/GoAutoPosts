package instagram

import (
	"fmt"
	"net/url"
)

// Publisher handles the complete publishing workflow
type Publisher struct {
	client *Client
}

// NewPublisher creates a new publisher
func NewPublisher(client *Client) *Publisher {
	return &Publisher{
		client: client,
	}
}

// PublishPost publishes a complete single post to Instagram
func (p *Publisher) PublishPost(imagePath, caption string) (string, error) {
	// Step 1: Upload image
	imageID, err := p.client.UploadImage(imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to upload image: %w", err)
	}

	// Step 2: Create media container with caption
	encodedCaption := url.QueryEscape(caption)
	creationID, err := p.client.CreateMedia(imageID, encodedCaption)
	if err != nil {
		return "", fmt.Errorf("failed to create media: %w", err)
	}

	// Step 3: Publish media
	postID, err := p.client.PublishMedia(creationID)
	if err != nil {
		return "", fmt.Errorf("failed to publish media: %w", err)
	}

	return postID, nil
}

// PublishCarousel publishes a carousel post to Instagram
func (p *Publisher) PublishCarousel(imagePaths []string, caption string) (string, error) {
	// Step 1: Upload all images as carousel items
	var childrenIDs []string
	for _, path := range imagePaths {
		id, err := p.client.UploadCarouselImage(path)
		if err != nil {
			return "", fmt.Errorf("failed to upload carousel item %s: %w", path, err)
		}
		childrenIDs = append(childrenIDs, id)
	}

	// Step 2: Create Carousel container
	encodedCaption := url.QueryEscape(caption)
	creationID, err := p.client.CreateCarouselContainer(childrenIDs, encodedCaption)
	if err != nil {
		return "", fmt.Errorf("failed to create carousel container: %w", err)
	}

	// Step 3: Publish
	postID, err := p.client.PublishMedia(creationID)
	if err != nil {
		return "", fmt.Errorf("failed to publish carousel: %w", err)
	}

	return postID, nil
}
