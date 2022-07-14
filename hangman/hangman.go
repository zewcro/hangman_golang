package hangman

import "strings"

type Game struct {
	State        string   // Game state
	Letters      []string // Letters in the word to find
	FoundLetters []string // Good guesses
	UsedLetters  []string
	TurnsLeft    int // Remaining attemps
}

func New(turns int, word string) *Game {
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters)) // create a slice of same size as letters var
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}

	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}
	return g
}
