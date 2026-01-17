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
	// Modern Developer Dark Theme
	ColorBackgroundStart = HexToColor("#0F172A") // Slate 900
	ColorBackgroundEnd   = HexToColor("#1E293B") // Slate 800
	ColorTextPrimary     = HexToColor("#F8FAFC") // Slate 50
	ColorTextSecondary   = HexToColor("#94A3B8") // Slate 400
	ColorAccent          = HexToColor("#38BDF8") // Sky 400

	// Code Window Colors
	ColorCodeBackground      = HexToColor("#1E1E1E") // VS Code Dark
	ColorWindowControlRed    = HexToColor("#FF5F56")
	ColorWindowControlYellow = HexToColor("#FFBD2E")
	ColorWindowControlGreen  = HexToColor("#27C93F")

	// Syntax Highlighting
	ColorKeyword  = HexToColor("#C586C0") // Purple
	ColorString   = HexToColor("#CE9178") // Orange/Brown
	ColorComment  = HexToColor("#6A9955") // Green
	ColorFunction = HexToColor("#DCDCAA") // Yellow
	ColorNormal   = HexToColor("#D4D4D4") // Light Gray
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
