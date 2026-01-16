# Go Daily - Instagram Automation for Go Libraries

Automated Instagram posting system that shares Go libraries daily using the Meta Graph API.

## Features

- 🤖 **Automated Daily Posts**: Scheduled via GitHub Actions
- 📚 **Library Management**: JSON-based library database
- 🎨 **Image Generation**: Dynamic image creation with library details
- 📝 **Template System**: Customizable caption templates
- 🏷️ **Smart Hashtags**: Automatic hashtag generation
- 📊 **History Tracking**: Prevents duplicate posts within 30 days
- 🔄 **Dual Storage**: JSON or SQLite backend options

## Project Structure

```

├── cmd/publisher/          # Application entry point
├── internal/
│   ├── config/            # Configuration management
│   ├── selector/          # Library selection logic
│   ├── template/          # Caption templates
│   ├── hashtag/           # Hashtag generation
│   ├── image/             # Image generation
│   ├── instagram/         # Meta Graph API client
│   ├── store/             # Data persistence
│   ├── model/             # Data models
│   └── logger/            # Structured logging
├── data/                  # JSON data files
├── scripts/               # Utility scripts
└── .github/workflows/     # GitHub Actions
```

## Setup

### Prerequisites

- Go 1.21 or higher
- Instagram Business Account
- Meta Graph API Access Token

### Installation

1. Clone the repository:

```bash
git clone https://github.com/nitin737/GoAutoPosts.git
cd GoAutoPosts
```

2. Install dependencies:

```bash
make install
```

3. Configure environment variables:

```bash
cp .env.example .env
# Edit .env with your credentials
```

4. Validate your library data:

```bash
make validate
```

### Configuration

Set the following environment variables:

- `INSTAGRAM_ACCESS_TOKEN`: Your Meta Graph API access token
- `INSTAGRAM_ACCOUNT_ID`: Your Instagram Business Account ID
- `LIBRARIES_PATH`: Path to libraries.json (default: `data/libraries.json`)
- `POSTED_PATH`: Path to posted.json (default: `data/posted.json`)

## Usage

### Local Development

Run the publisher locally:

```bash
make run
```

Run in development mode with debug logging:

```bash
make dev
```

### GitHub Actions

The workflow runs automatically every day at 9:00 AM UTC. You can also trigger it manually:

1. Go to Actions tab in your repository
2. Select "Daily Instagram Post"
3. Click "Run workflow"

### Managing Libraries

Add new libraries to `data/libraries.json`:

```json
{
  "name": "library-name",
  "description": "Brief description",
  "url": "https://github.com/author/library",
  "category": "Category Name",
  "tags": ["tag1", "tag2"],
  "stars": 10000,
  "author": "author-name"
}
```

Validate your changes:

```bash
make validate
```

## Development

### Build

```bash
make build
```

### Test

```bash
make test
```

### Format Code

```bash
make fmt
```

## Storage Options

### JSON Store (Default)

Simple file-based storage using `data/posted.json`.

### SQLite Store (Optional)

For better performance and querying:

```go
store, err := store.NewSQLiteStore("data/posted.db")
```

## License

MIT License - see LICENSE file for details

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
