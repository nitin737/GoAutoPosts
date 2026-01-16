# Go Daily - Quick Start Guide

## 📁 Repository Structure Created

✅ **Complete project structure** has been created with all necessary files and directories.

### Key Components

#### 1. **Entry Point**

- `cmd/publisher/main.go` - Main application entry point

#### 2. **Core Modules**

- `internal/config/` - Configuration management with environment variables
- `internal/selector/` - Random library selection with history filtering
- `internal/template/` - Caption template rendering
- `internal/hashtag/` - Automatic hashtag generation
- `internal/image/` - Dynamic image generation with text overlay
- `internal/instagram/` - Meta Graph API client and publisher
- `internal/store/` - Data persistence (JSON & SQLite options)
- `internal/model/` - Data models
- `internal/logger/` - Structured logging

#### 3. **Data Files**

- `data/libraries.json` - Library database (3 sample libraries included)
- `data/posted.json` - Posting history tracker
- `internal/image/assets/base.png` - Instagram post template (447KB)

#### 4. **Automation**

- `.github/workflows/daily-publish.yml` - GitHub Actions workflow for daily posts

#### 5. **Development Tools**

- `Makefile` - Common development tasks
- `scripts/validate_data.go` - Data validation utility
- `.env.example` - Environment variable template

## 🚀 Next Steps

### 1. Set Up Environment Variables

```bash

cp .env.example .env
```

Edit `.env` and add your Instagram credentials:

- `INSTAGRAM_ACCESS_TOKEN` - Get from Meta Developer Console
- `INSTAGRAM_ACCOUNT_ID` - Your Instagram Business Account ID

### 2. Install Dependencies

```bash
make install
```

### 3. Add More Libraries

Edit `data/libraries.json` to add more Go libraries:

```json
{
  "name": "your-library",
  "description": "What it does",
  "url": "https://github.com/author/library",
  "category": "Category",
  "tags": ["tag1", "tag2"],
  "stars": 1000,
  "author": "author-name"
}
```

### 4. Validate Your Data

```bash
make validate
```

### 5. Test Locally

```bash
make dev
```

### 6. Set Up GitHub Secrets

In your GitHub repository settings, add these secrets:

- `INSTAGRAM_ACCESS_TOKEN`
- `INSTAGRAM_ACCOUNT_ID`

### 7. Enable GitHub Actions

The workflow will run automatically every day at 9:00 AM UTC.
You can also trigger it manually from the Actions tab.

## 📋 Available Make Commands

- `make help` - Show all available commands
- `make install` - Install dependencies
- `make build` - Build the application
- `make run` - Run the application
- `make dev` - Run in development mode
- `make test` - Run tests
- `make validate` - Validate libraries.json
- `make clean` - Clean build artifacts
- `make fmt` - Format code
- `make lint` - Run linter

## 🎨 Customization

### Modify Caption Template

Edit `internal/template/caption.tmpl` to customize the Instagram caption format.

### Customize Image Design

Replace `internal/image/assets/base.png` with your own design, or modify `internal/image/generator.go` to change text positioning and styling.

### Adjust Hashtags

Edit `internal/hashtag/generator.go` to modify the base hashtags or hashtag generation logic.

### Change Posting Frequency

Modify the cron schedule in `.github/workflows/daily-publish.yml`:

```yaml
schedule:
  - cron: "0 9 * * *" # Daily at 9 AM UTC
```

## 🔧 Storage Options

### JSON Store (Default)

Simple file-based storage using `data/posted.json`.

### SQLite Store (Optional)

For better performance, modify `cmd/publisher/main.go` to use SQLite:

```go
store, err := store.NewSQLiteStore("data/posted.db")
```

## 📊 Project Statistics

- **Total Files**: 24 source files
- **Go Packages**: 9 internal packages
- **Sample Libraries**: 3 included
- **Base Image**: 447KB professional template
- **Dependencies**: 4 external packages

## 🎯 Features Implemented

✅ Automated daily posting via GitHub Actions
✅ Random library selection with 30-day cooldown
✅ Dynamic image generation with library details
✅ Template-based caption rendering
✅ Smart hashtag generation
✅ Posting history tracking
✅ Meta Graph API integration
✅ Dual storage options (JSON/SQLite)
✅ Structured logging
✅ Data validation tools
✅ Comprehensive documentation

## 📝 Notes

- The base image template has been generated with a modern, tech-focused design
- All Go modules have been downloaded and are ready to use
- Data validation passed successfully for the sample libraries
- The project follows Go best practices and clean architecture principles

## 🆘 Troubleshooting

### Missing Dependencies

```bash
go mod tidy
```

### Validation Errors

```bash
make validate
```

### Build Issues

```bash
make clean
make build
```

---

**Ready to start posting Go libraries to Instagram! 🎉**
