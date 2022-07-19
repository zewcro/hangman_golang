package hangman

import "fmt"

func DrawWelcome() {
	fmt.Println(`
$$\   $$\                                                                 
$$ |  $$ |                                                                
$$ |  $$ | $$$$$$\  $$$$$$$\   $$$$$$\  $$$$$$\$$$$\   $$$$$$\  $$$$$$$\  
$$$$$$$$ | \____$$\ $$  __$$\ $$  __$$\ $$  _$$  _$$\  \____$$\ $$  __$$\ 
$$  __$$ | $$$$$$$ |$$ |  $$ |$$ /  $$ |$$ / $$ / $$ | $$$$$$$ |$$ |  $$ |
$$ |  $$ |$$  __$$ |$$ |  $$ |$$ |  $$ |$$ | $$ | $$ |$$  __$$ |$$ |  $$ |
$$ |  $$ |\$$$$$$$ |$$ |  $$ |\$$$$$$$ |$$ | $$ | $$ |\$$$$$$$ |$$ |  $$ |
\__|  \__| \_______|\__|  \__| \____$$ |\__| \__| \__| \_______|\__|  \__|
                              $$\   $$ |                                  
                              \$$$$$$  |                                  
                               \______/                                   `)
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
		__________
		|	  	  | 
		|	  	  o
		|		 /|\
		|		  |
		|		 / \ 
		|		 
		| 
		======
		|====|
		|____|
`

	case 1:
		draw = `
	__________
	|	  	  | 
	|	  	  o
	|		 /|\
	|		  |
	|		 / 
	|		 
	| 
	======
	|====|
	|____|
`
	case 2:
		draw = `
	__________
	|	  	  | 
	|	  	  o
	|		 /|\
	|		  |
	|		 
	|		 
	| 
	======
	|====|
	|____|

`
	case 3:
		draw = `
	__________
	|	  	  | 
	|	  	  o
	|		 /|\
	|		  
	|		 
	|		 
	| 
	======
	|====|
	|____|

`
	case 4:
		draw = `
	__________
	|	  	  | 
	|	  	  o
	|		 /|
	|		  
	|		 
	|		 
	| 
	======
	|====|
	|____|`

	case 5:
		draw = `
	__________
	|	  	  | 
	|	  	  o
	|		  |
	|		  
	|		 
	|		 
	| 
	======
	|====|
	|____|`

	case 6:
		draw = `
	__________
	|	  	  | 
	|	  	  o
	|		  
	|		  
	|		 
	|		 
	| 
	======
	|====|
	|____|`

	case 7:
		draw = `
	__________
	|	  	  | 
	|	  	  
	|		  
	|		  
	|		 
	|		 
	| 
	======
	|====|
	|____|`

	case 8:
		draw = ` `

	}

	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed :")
	drawLetters(g.FoundLetters)

	fmt.Print("Used :")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Print("Good guess!")
	case "alreadyGuessed":
		fmt.Printf("Letter ' %s ' was already used\n", guess)
	case "badGuess":
		fmt.Printf("Bad Guess, '%s' is not in the word", guess)
	case "lost":
		fmt.Print("You lost :( ! The word was :")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("YOU WON !! The word was :")
		drawLetters(g.Letters)

	}

}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v", c)
	}
	fmt.Println()
}
