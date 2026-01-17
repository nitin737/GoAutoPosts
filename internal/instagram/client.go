package instagram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Client handles Instagram Graph API interactions
type Client struct {
	accessToken string
	accountID   string
	graphAPIURL string
	httpClient  *http.Client
}

// NewClient creates a new Instagram API client
func NewClient(accessToken, accountID, graphAPIURL string) *Client {
	return &Client{
		accessToken: accessToken,
		accountID:   accountID,
		graphAPIURL: graphAPIURL,
		httpClient:  &http.Client{},
	}
}

// UploadImageResponse represents the response from image upload
type UploadImageResponse struct {
	ID string `json:"id"`
}

// CreateMediaResponse represents the response from media creation
type CreateMediaResponse struct {
	ID string `json:"id"`
}

// PublishResponse represents the response from publishing
type PublishResponse struct {
	ID string `json:"id"`
}

// UploadImage uploads an image to Instagram (Legacy single post)
func (c *Client) UploadImage(imagePath string) (string, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("image", imagePath)
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(part, file); err != nil {
		return "", err
	}

	if err := writer.Close(); err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("access_token", c.accessToken)

	u := fmt.Sprintf("%s/%s/media?%s", c.graphAPIURL, c.accountID, params.Encode())
	req, err := http.NewRequest("POST", u, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("upload failed (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var uploadResp UploadImageResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadResp); err != nil {
		return "", err
	}

	return uploadResp.ID, nil
}

// CreateCarouselItem creates a carousel item container from a public image URL
func (c *Client) CreateCarouselItem(imageURL string) (string, error) {
	params := url.Values{}
	params.Set("is_carousel_item", "true")
	params.Set("image_url", imageURL)
	params.Set("access_token", c.accessToken)

	u := fmt.Sprintf("%s/%s/media?%s", c.graphAPIURL, c.accountID, params.Encode())

	resp, err := c.httpClient.Post(u, "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("create carousel item failed (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var uploadResp UploadImageResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadResp); err != nil {
		return "", err
	}

	return uploadResp.ID, nil
}

// CreateMedia creates a single media container
func (c *Client) CreateMedia(imageURLOrID, caption string) (string, error) {
	// If the "CreateMedia" was used with IDs (from upload), we use image_url?
	// That's confusing. But if existing worked...
	// Standard: image_url=URL
	params := url.Values{}
	params.Set("image_url", imageURLOrID)
	params.Set("caption", caption)
	params.Set("access_token", c.accessToken)

	u := fmt.Sprintf("%s/%s/media?%s", c.graphAPIURL, c.accountID, params.Encode())

	resp, err := c.httpClient.Post(u, "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("create media failed (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var mediaResp CreateMediaResponse
	if err := json.NewDecoder(resp.Body).Decode(&mediaResp); err != nil {
		return "", err
	}

	return mediaResp.ID, nil
}

// CreateCarouselContainer creates the carousel container with children
func (c *Client) CreateCarouselContainer(children []string, caption string) (string, error) {
	childrenStr := strings.Join(children, ",")
	// media_type=CAROUSEL
	params := url.Values{}
	params.Set("media_type", "CAROUSEL")
	params.Set("children", childrenStr)
	params.Set("caption", caption)
	params.Set("access_token", c.accessToken)

	u := fmt.Sprintf("%s/%s/media?%s", c.graphAPIURL, c.accountID, params.Encode())

	resp, err := c.httpClient.Post(u, "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("create carousel container failed (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var mediaResp CreateMediaResponse
	if err := json.NewDecoder(resp.Body).Decode(&mediaResp); err != nil {
		return "", err
	}

	return mediaResp.ID, nil
}

// PublishMedia publishes a media container (works for both single and carousel)
func (c *Client) PublishMedia(creationID string) (string, error) {
	params := url.Values{}
	params.Set("creation_id", creationID)
	params.Set("access_token", c.accessToken)

	u := fmt.Sprintf("%s/%s/media_publish?%s", c.graphAPIURL, c.accountID, params.Encode())

	resp, err := c.httpClient.Post(u, "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("publish failed (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var publishResp PublishResponse
	if err := json.NewDecoder(resp.Body).Decode(&publishResp); err != nil {
		return "", err
	}

	return publishResp.ID, nil
}
