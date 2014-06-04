package term

import (
	"testing"
	"os"
	"io"
	"bytes"
)

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

func TestBlack(t *testing.T) {
	text := Black("color test")
	if (text.content != "color test") { t.Error() }
	if (text.color != 30) { t.Error() }
}

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
func TestDefault(t *testing.T) {
	text := Default("color test")
	if (text.content != "color test") { t.Error() }
	if (text.color != 39) { t.Error() }
}

// Background Colors

func TestOnBlack(t *testing.T) {
	if (Red("Hello").OnBlack().background != 40) { t.Error() }
}
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
func TestDim(t *testing.T) {
	if (Red("Hello").Dim().style != 2) { t.Error() }
}
func TestItalics(t *testing.T) {
	if (Red("Hello").Italics().style != 3) { t.Error() }
}
func TestUnderlined(t *testing.T) {
	if (Red("Hello").Underlined().style != 4) { t.Error() }
}
func TestBlinking(t *testing.T) {
	if (Red("Hello").Blinking().style != 5) { t.Error() }
}
func TestFastBlinking(t *testing.T) {
	if (Red("Hello").FastBlinking().style != 6) { t.Error() }
}
func TestInverted(t *testing.T) {
	if (Red("Hello").Inverted().style != 7) { t.Error() }
}
func TestHidden(t *testing.T) {
	if (Red("Hello").Hidden().style != 8) { t.Error() }
}

// Raw

func TestRawAllSet(t *testing.T) {
	text := Red("Hello").OnRed().Bold().Raw()
	if (text != "\033[31;41;1mHello\033[0m") { t.Error() }
}

func TestRawColorAndBackgroundSet(t *testing.T) {
	text := Red("Hello").OnRed().Raw()
	if (text != "\033[31;41mHello\033[0m") { t.Error() }
}

// Print

func TestPrint(t *testing.T) {
    old := os.Stdout // keep backup of the real stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    White("White on Black").OnBlack().Print()

    outC := make(chan string)
    // copy the output in a separate goroutine so printing can't block indefinitely
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, r)
        outC <- buf.String()
    }()

    // back to normal state
    w.Close()
    os.Stdout = old // restoring the real stdout
    out := <-outC

    if out != "\033[37;40mWhite on Black\033[0m" { t.Error(out) }
}

func TestPrintln(t *testing.T) {
    old := os.Stdout // keep backup of the real stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    White("White on Black").OnBlack().Println()

    outC := make(chan string)
    // copy the output in a separate goroutine so printing can't block indefinitely
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, r)
        outC <- buf.String()
    }()

    // back to normal state
    w.Close()
    os.Stdout = old // restoring the real stdout
    out := <-outC

    if out != "\033[37;40mWhite on Black\033[0m\n" { t.Error(out) }
}
