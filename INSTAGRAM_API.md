# Instagram Graph API Endpoints

This document outlines the Instagram Graph API endpoints used by the `GoAutoPosts` publisher.

## Base URL

The default Graph API URL is: `https://graph.facebook.com/v18.0`

---

## 1. Create Media Container (Single Image)

Used to upload a single image to Instagram's servers before publishing.

- **Endpoint:** `POST /{ig-user-id}/media`
- **Authentication:** `access_token` (User Access Token)
- **Request Parameters:**
  | Parameter | Type | Description |
  | :--- | :--- | :--- |
  | `image_url` | string | **Required.** The public URL of the image. |
  | `caption` | string | Optional. The caption for the post. |
  | `access_token` | string | **Required.** Valid User Access Token. |

- **Go Method:** `CreateMedia(imageURLOrID, caption string)`
- **Success Response:**
  ```json
  {
    "id": "1234567890" // This is the creation_id
  }
  ```

---

## 2. Create Media Container (Carousel Item)

Used to upload individual items (images/videos) that will be part of a carousel post.

- **Endpoint:** `POST /{ig-user-id}/media`
- **Authentication:** `access_token`
- **Request Parameters:**
  | Parameter | Type | Description |
  | :--- | :--- | :--- |
  | `image_url` | string | **Required.** The public URL of the image. |
  | `is_carousel_item` | boolean | **Required.** Must be set to `true`. |
  | `access_token` | string | **Required.** |

- **Go Method:** `UploadCarouselImage(imagePath string)` (Note: Current impl uses multipart, but IG usually requires `image_url`)
- **Success Response:**
  ```json
  {
    "id": "9876543210"
  }
  ```

---

## 3. Create Carousel Container

Used to group multiple `creation_id`s from Step 2 into a single carousel post.

- **Endpoint:** `POST /{ig-user-id}/media`
- **Authentication:** `access_token`
- **Request Parameters:**
  | Parameter | Type | Description |
  | :--- | :--- | :--- |
  | `media_type` | string | **Required.** Must be `CAROUSEL`. |
  | `children` | string | **Required.** Comma-separated list of children IDs. |
  | `caption` | string | Optional. The caption for the carousel. |
  | `access_token` | string | **Required.** |

- **Go Method:** `CreateCarouselContainer(children []string, caption string)`
- **Success Response:**
  ```json
  {
    "id": "1122334455" // The carousel creation_id
  }
  ```

---

## 4. Publish Media

The final step to make the post (single or carousel) live on the profile.

- **Endpoint:** `POST /{ig-user-id}/media_publish`
- **Authentication:** `access_token`
- **Request Parameters:**
  | Parameter | Type | Description |
  | :--- | :--- | :--- |
  | `creation_id` | string | **Required.** The ID from Step 1 or Step 3. |
  | `access_token` | string | **Required.** |

- **Go Method:** `PublishMedia(creationID string)`
- **Success Response:**
  ```json
  {
    "id": "17841405303063522" // The final Post ID
  }
  ```

---

## Error Response Format

When an error occurs, the API returns a standard error object:

```json
{
  "error": {
    "message": "Error message description",
    "type": "OAuthException",
    "code": 100,
    "error_subcode": 2207026,
    "fbtrace_id": "A1B2C3D4E5"
  }
}
```

## Important Limitations

1. **Public URLs:** Instagram's servers must be able to reach your `image_url`. Local paths will not work.
2. **Rate Limits:** Keep an eye on your app's dashboard for rate limit usage.
3. **Carousel Limit:** A maximum of 10 items can be included in a carousel.
