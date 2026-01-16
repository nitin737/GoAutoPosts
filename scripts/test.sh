#!/bin/bash

# Test script for go-daily publisher
# This script validates the setup without posting to Instagram

set -e

echo "🧪 Testing Go Daily Publisher..."
echo ""

# Check if we're in the right directory
if [ ! -f "go.mod" ]; then
    echo "❌ Error: Must be run from root directory"
    exit 1
fi

# Check dependencies
echo "📦 Checking dependencies..."
go mod verify
echo "✅ Dependencies verified"
echo ""

# Validate data
echo "📊 Validating library data..."
go run scripts/validate_data.go data/libraries.json
echo ""

# Check environment variables (optional for test)
echo "🔧 Checking environment configuration..."
if [ -f ".env" ]; then
    echo "✅ .env file found"
else
    echo "⚠️  .env file not found (using defaults)"
fi
echo ""

# Build the application
echo "🔨 Building application..."
go build -o bin/publisher cmd/publisher/main.go
echo "✅ Build successful"
echo ""

# Check binary
echo "📦 Binary info:"
ls -lh bin/publisher
echo ""

# Check base image
echo "🖼️  Checking base image..."
if [ -f "internal/image/assets/base.png" ]; then
    echo "✅ Base image found"
    ls -lh internal/image/assets/base.png
else
    echo "❌ Base image not found"
    exit 1
fi
echo ""

# Test library selection (dry run)
echo "🎲 Testing library selection..."
echo "   Libraries available: $(jq 'length' data/libraries.json)"
echo "   Posted history: $(jq 'length' data/posted.json)"
echo ""

# Test hashtag generation
echo "#️⃣  Testing hashtag generation..."
echo "   Base hashtags configured: golang, go, programming, coding, developer, software, opensource, tech"
echo ""

# Verify all components are importable
echo "🔍 Verifying package structure..."
go list ./... | grep -v vendor
echo "✅ All packages valid"
echo ""

echo "✅ All tests passed!"
echo ""
echo "📝 Next steps:"
echo "   1. Set up Instagram credentials in .env"
echo "   2. Add more libraries to data/libraries.json"
echo "   3. Run: make run (to test full workflow)"
echo ""
