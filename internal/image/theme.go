package image

import (
	"image/color"
)

// Design System Constants

// Dimensions
const (
	Width   = 1080
	Height  = 1080
	Padding = 80
)

// Colors
var (
	ColorBackgroundStart = HexToColor("#1a1a1a") // Example Dark Theme
	ColorBackgroundEnd   = HexToColor("#2d2d2d")
	ColorTextPrimary     = HexToColor("#ffffff")
	ColorTextSecondary   = HexToColor("#a0a0a0")
	ColorAccent          = HexToColor("#00ADD8") // Go Blue
	ColorCodeBackground  = HexToColor("#000000") // 80% opacity usually
)

// Typography
const (
	FontSizeTitle    = 72.0
	FontSizeSubtitle = 48.0
	FontSizeBody     = 42.0
	FontSizeCode     = 32.0
	FontSizeFooter   = 24.0
)

// HexToColor parses a hex string (e.g., "#1a1a1a") to color.RGBA
func HexToColor(s string) color.RGBA {
	c := color.RGBA{A: 0xff}
	if s[0] != '#' {
		return c
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	}
	return c
}
