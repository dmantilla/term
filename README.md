# term

Simple package to print text on the console with ANSI colors and styles.

## Installation

$ go get github.com/dmantilla/term

## Usage

	term.Blue("Blue on White Blinking").OnWhite().Blinking().Println()
	
	term.Green("Green on Black Underlined").OnBlack().Underlined().Print()
	term.Yellow(" and Yellow text").Println()
	
	fmt.Println(term.White("White on blue").OnBlue().Raw())
	
	fmt.Println(term.Default("Default Text").Raw())
	
	fmt.Println(term.Default("Default text on Red").OnRed().Raw())
	
	fmt.Println(term.Yellow("Yellow on Cyan Underlined").OnCyan().Underlined().Raw())
	
