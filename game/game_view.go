package game

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// TODO: validate all the inputs
func DisplayAllGames() {
	OutputFormatter(Games)
}

func OutputFormatter(games map[int]Game) {
	var Writer = new(tabwriter.Writer)
	Writer.Init(os.Stdout, 0, 8, 1, '\t', 0)
	fmt.Fprintln(Writer, "ID\tTitle\tDescription\tGenre\tPrice\tStocks")

	for id, game := range games {
		fmt.Fprintf(Writer, "%d\t%s\t%s\t%s\t%.2f\t%d\n", id, game.Title, game.Description, game.Genre, game.Price, game.Stocks)
	}
	Writer.Flush()
	fmt.Println(Line)
}
