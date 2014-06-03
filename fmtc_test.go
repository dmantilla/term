package fmtc

import "testing"

func TestContructor(t *testing.T) {
	text := construct()
	if (text != ColoredText{color: 0, background: 49, style: -1}) { t.Error() }
}

func TestSetContent(t *testing.T) {
	text1 := construct()
	text2 := text1.setContent("Hi")
	if (text1 != *text2) { t.Error() }
	if (text1.content != "Hi") { t.Errorf("text.content is '%s'", text1.content) }
}

func TestSetColor(t *testing.T) {
	text1 := construct()
	text2 := text1.setColor('r')
	if (text1 != *text2) { t.Error() }
	if (text1.color != 31) { t.Error() }
}

func TestSetBackground(t *testing.T) {
	text1 := construct()
	text2 := text1.setBackground('r')
	if (text1 != *text2) { t.Error() }
	if (text1.background != 41) { t.Error() }
}

func TestSetStyle(t *testing.T) {
	text1 := construct()
	text2 := text1.setStyle('!')
	if (text1 != *text2) { t.Error() }
	if (text1.style != 1) { t.Error() }
}

// Foreground Colors

func TestRed(t *testing.T) {
	text := Red("color test")
	if (text.content != "color test") { t.Error() }
	if (text.color != 31) { t.Error() }
}

func TestGreen(t *testing.T) {
	text := Green("color test")
	if (text.content != "color test") { t.Error() }
	if (text.color != 32) { t.Error() }
}
func TestYellow(t *testing.T) {
	text := Yellow("color test")
	if (text.content != "color test") { t.Error() }
	if (text.color != 33) { t.Error() }
}
func TestBlue(t *testing.T) {
	text := Blue("color test")
	if (text.content != "color test") { t.Error() }
	if (text.color != 34) { t.Error() }
}
func TestMagenta(t *testing.T) {
	text := Magenta("color test")
	if (text.content != "color test") { t.Error() }
	if (text.color != 35) { t.Error() }
}
func TestCyan(t *testing.T) {
	text := Cyan("color test")
	if (text.content != "color test") { t.Error() }
	if (text.color != 36) { t.Error() }
}
func TestWhite(t *testing.T) {
	text := White("color test")
	if (text.content != "color test") { t.Error() }
	if (text.color != 37) { t.Error() }
}

// Background Colors

func TestOnRed(t *testing.T) {
	if (Red("Hello").OnRed().background != 41) { t.Error() }
}
func TestOnGreen(t *testing.T) {
	if (Red("Hello").OnGreen().background != 42) { t.Error() }
}
func TestOnYellow(t *testing.T) {
	if (Red("Hello").OnYellow().background != 43) { t.Error() }
}
func TestOnBlue(t *testing.T) {
	if (Red("Hello").OnBlue().background != 44) { t.Error() }
}
func TestOnMagenta(t *testing.T) {
	if (Red("Hello").OnMagenta().background != 45) { t.Error() }
}
func TestOnCyan(t *testing.T) {
	if (Red("Hello").OnCyan().background != 46) { t.Error() }
}
func TestOnWhite(t *testing.T) {
	if (Red("Hello").OnWhite().background != 47) { t.Error() }
}

// Styles

func TestBold(t *testing.T) {
	if (Red("Hello").Bold().style != 1) { t.Error() }
}

// Raw

func TestRawAllSet(t *testing.T) {
	text := Red("Hello").OnRed().Bold().Raw()
	if (text != "\033[31;41;1mHello") { t.Error() }
}

func TestRawColorAndBackgroundSet(t *testing.T) {
	text := Red("Hello").OnRed().Raw()
	if (text != "\033[31;41mHello") { t.Error() }
}
