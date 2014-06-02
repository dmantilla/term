package fmtc

import "fmt"

const (
	resetCode = "\033[0m"
	baseCode = "\033[%d;%d;%dm%s"
)

var foregroundColors = map[int]int{
	'k': 30, // Black
	'r': 31, // Red
	'g': 32, // Green
	'y': 33, // Yellow
	'b': 34, // Blue
	'm': 35, // Magenta
	'c': 36, // Cyan
	'w': 37, // White
	0  : 39, // Default	
}

var backgroundColors = map[int]int{
	'k': 40, // Black
	'r': 41, // Red
	'g': 42, // Green
	'y': 43, // Yellow
	'b': 44, // Blue
	'm': 45, // Magenta
	'c': 46, // Cyan
	'w': 47, // White
	0  : 49, // Default
}

var styles = map[int]int{
	0  : 0, // Default
	'!': 1, // Bold
	'.': 2, // Dim
	'/': 3, // Italic
	'_': 4, // Underline
	'^': 5, // Blink
	'&': 6, // Fast blink
	'?': 7, // Reverse foreground and background colors
	'-': 8, // Hide
}

type ColoredText struct {
	content string
	color int
	background int
	style int
}

func construct() ColoredText {
	// return ColoredText{color: foregroundColors[0], background: backgroundColors[0], style: styles[0] }
	return ColoredText{color: foregroundColors[0], background: backgroundColors[0], style: styles[0] }
}

func (text *ColoredText) setContent(content string) *ColoredText {
	text.content = content
	return text
}

func (text *ColoredText) setColor(colorCode int) *ColoredText {
	text.color = foregroundColors[colorCode]
	return text
}

func (text *ColoredText) setBackground(colorCode int) *ColoredText {
	text.background = backgroundColors[colorCode]
	return text
}

func (text *ColoredText) setStyle(styleCode int) *ColoredText {
	text.style = styles[styleCode]
	return text
}

// Foreground Colors
func Red(content string) *ColoredText {
	text := construct()
	return text.setContent(content).setColor('r')
}

// Background Colors
func (text *ColoredText) OnRed() *ColoredText {
	return text.setBackground('r')
}

// Styles
func (text *ColoredText) Bold() *ColoredText {
	return text.setStyle('!')
}

// Output
func (text *ColoredText) Raw() string {
	return fmt.Sprintf("\033[%d;%d;%dm%s", text.color, text.background, text.style, text.content)
}
