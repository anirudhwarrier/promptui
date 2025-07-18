package promptui

import "github.com/chzyer/readline"

// These runes are used to identify the commands entered by the user in the command prompt. They map
// to specific actions of promptui in prompt mode and can be remapped if necessary.
var (
	// KeyEnter is the default key for submission/selection.
	KeyEnter rune = readline.CharEnter

	// KeyCtrlH is the key for deleting input text.
	KeyCtrlH rune = readline.CharCtrlH

	// KeyPrev is the default key to go up during selection.
	KeyPrevDisplay      = "↑"

	// KeyNext is the default key to go down during selection.
	KeyNextDisplay      = "↓"

	// KeyBackward is the default key to page up during selection.
	KeyBackwardDisplay      = "←"

	// KeyForward is the default key to page down during selection.
	KeyForwardDisplay      = "→"
)
