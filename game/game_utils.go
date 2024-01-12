package game

import "strings"

//Validate A single Game before adding to Games
func ValidateGame(game Game) {

}

func ValidateString(text string) {
}

func TrimSpaces(text string) string {
	text = strings.Trim(text, " ")
	return text
}
