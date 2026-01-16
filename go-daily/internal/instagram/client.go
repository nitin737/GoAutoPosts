package instagram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	graphAPIURL = "https://graph.facebook.com/v18.0"
)

// Client handles Instagram Graph API interactions
type Client struct {
	accessToken string
	accountID   string
	httpClient  *http.Client
}

// NewClient creates a new Instagram API client
func NewClient(accessToken, accountID string) *Client {
	return &Client{
		accessToken: accessToken,
		accountID:   accountID,
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

// UploadImage uploads an image to Instagram
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

	url := fmt.Sprintf("%s/%s/media", graphAPIURL, c.accountID)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+c.accessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("upload failed: %s", string(bodyBytes))
	}

	var uploadResp UploadImageResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadResp); err != nil {
		return "", err
	}

	return uploadResp.ID, nil
}

// CreateMedia creates a media container
func (c *Client) CreateMedia(imageURL, caption string) (string, error) {
	url := fmt.Sprintf("%s/%s/media?image_url=%s&caption=%s&access_token=%s",
		graphAPIURL, c.accountID, imageURL, caption, c.accessToken)

	resp, err := c.httpClient.Post(url, "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("create media failed: %s", string(bodyBytes))
	}

	var mediaResp CreateMediaResponse
	if err := json.NewDecoder(resp.Body).Decode(&mediaResp); err != nil {
		return "", err
	}

	return mediaResp.ID, nil
}

// PublishMedia publishes a media container
func (c *Client) PublishMedia(creationID string) (string, error) {
	url := fmt.Sprintf("%s/%s/media_publish?creation_id=%s&access_token=%s",
		graphAPIURL, c.accountID, creationID, c.accessToken)

	resp, err := c.httpClient.Post(url, "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("publish failed: %s", string(bodyBytes))
	}

	var publishResp PublishResponse
	if err := json.NewDecoder(resp.Body).Decode(&publishResp); err != nil {
		return "", err
	}

	return publishResp.ID, nil
}
