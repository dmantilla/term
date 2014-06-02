package fmtc

import "testing"

func TestContructor(t *testing.T) {
	text := construct()
	if (text != ColoredText{color: 39, background: 49, style: 0}) { t.Error() }
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

func TestRed(t *testing.T) {
	text := Red("Hello")
	if (text.content != "Hello") { t.Errorf("text.content is '%s'", text.content) }
	if (text.color != 31) { t.Errorf("text.color is '%d'", text.color) }
}

func TestOnRed(t *testing.T) {
	text := Red("Hello").OnRed()
	if (text.background != 41) { t.Errorf("text.background is '%d'", text.background) }
}

func TestBold(t *testing.T) {
	text := Red("Hello").Bold()
	if (text.style != 1) { t.Error() }
}

func TestRaw(t *testing.T) {
	text := Red("Hello").OnRed().Bold().Raw()
	if (text != "\033[31;41;1mHello") { t.Error() }
}
