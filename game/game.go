package game

//handling input errors
//ayusin ang view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Line string = "-------------------------------------"
var newLine = "\n"

func InitializeGame() {
	fmt.Print(newLine, Line, newLine)
	fmt.Println("Game Inventory")
	fmt.Println(Line, newLine)
	fmt.Println(
		"Choices:",
		newLine,
		Line, newLine,
		"[1] Display Games", newLine,
		"[2] Update Game", newLine,
		"[3] Add Game", newLine,
		"[4] Delete Game", newLine,
		"[5] Filter Game", newLine,
		"[6] Sort Games", newLine,
		"[7] Exit", newLine,
		Line,
	)

	choice, err := inputChoice("menu")
	if err == nil {
		evaluateChoice(choice)
	}
}

func evaluateChoice(choice int) {
	switch choice {
	case 1:
		DisplayAllGames()
		InitializeGame()
	case 2:
		updateGame()
	case 3:
		addGame()
	case 4:
		deleteGame()
	case 5:
		filterGame()
	case 6:
		sortGames()
	case 7:
		ExitGame()
	default:
		invalidInput()
	}
}

func inputChoice(inputUsage string) (int, error) {
	var choice string
	address := &choice

	for {
		fmt.Print("Enter choice: ")
		fmt.Scanf("%v \n", address)
		choiceInt, err := strconv.ParseInt(choice, 10, 64)
		if inputUsage == "menu" {
			if err != nil || choiceInt < 1 || choiceInt > 7 {
				invalidInput()
			} else {
				return int(choiceInt), nil
			}
		} else {
			if err != nil || choiceInt < 1 || choiceInt > 3 {
				invalidInput()
			} else {
				return int(choiceInt), nil
			}
		}
	}
}

func inputFilter(choice int) {
	for {
		fmt.Print("Input Search:")
		reader := bufio.NewReader(os.Stdin)
		word, err := reader.ReadString('\n')
		word = strings.TrimSpace(word)

		if err != nil {
			fmt.Println("Invalid Word inputted")
		} else {
			switch choice {
			case 1:
				game, err := FilterByTitle(word)
				if err != nil {
					fmt.Println(err)
				} else {
					OutputFormatter(game)
				}
				filterGame()
			case 2:
				games, err := FilterByGenre(word)
				if err != nil {
					fmt.Println(err)
				} else {
					OutputFormatter(games)
				}
				filterGame()
			case 3:
				InitializeGame()
			default:
				fmt.Println("Invalid Input")
			}
			break
		}
	}
}
func inputGame() {

	//bufio to read the full line
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Title: ")
		title, errTitle := reader.ReadString('\n')
		fmt.Print("Description: ")
		description, errDesc := reader.ReadString('\n')
		fmt.Print("Genre: ")
		genre, errGenre := reader.ReadString('\n')
		fmt.Print("Price: ")
		price, errPrice := reader.ReadString('\n')
		price = strings.TrimSpace(price)
		fmt.Print("Stocks: ")
		stocks, errStocks := reader.ReadString('\n')
		stocks = strings.TrimSpace(stocks)

		if errTitle != nil {
			fmt.Println("Invalid Title")
		} else if errDesc != nil {
			fmt.Println("Invalid Description")
		} else if errGenre != nil {
			fmt.Println("Invalid Genre")
		} else if errPrice != nil {
			fmt.Println("Invalid Price")
		} else if errStocks != nil {
			fmt.Println("Invalid Stocks")
		} else {
			priceFloat, errPriceFloat := strconv.ParseFloat(price, 64)
			stocksInt, errStocksInt := strconv.ParseUint(stocks, 10, 64)

			if errPriceFloat != nil {
				fmt.Printf("%v", errPriceFloat)
			} else if errStocksInt != nil {
				fmt.Printf("%v", errStocksInt)
			} else if priceFloat < 1 {
				fmt.Printf("Invalid Price")
			} else if stocksInt < 1 {
				fmt.Println("Invalid Stocks")
			} else {
				AddGame(title, description, genre, priceFloat, stocksInt)
				InitializeGame()
				break
			}
		}

	}
}
func invalidInput() {
	fmt.Println("Invalid Input. Please try again", newLine, newLine)
	fmt.Println(Line)
}

func sortGames() {
	fmt.Print(Line,
		"\nSort By:", newLine,
		"[1] Price", newLine,
		"[2] Stocks", newLine,
		"[3] Exit", newLine,
	)
	fmt.Println(Line)
	sortChoice, err := inputChoice("sortChoice")
	fmt.Println(Line)
	if err == nil {
		switch sortChoice {
		case 1:
			OutputFormatter(SortGameByPriceOrStocks("price"))
			sortGames()
		case 2:
			OutputFormatter(SortGameByPriceOrStocks("stocks"))
			sortGames()
		case 3:
			InitializeGame()
		default:
			fmt.Println("Invalid Input")
		}

	}

}

func filterGame() {
	var choice string
	for {
		fmt.Print(Line,
			"\nSort By:", newLine,
			"[1] Title", newLine,
			"[2] Genre", newLine,
			"[3] Exit", newLine,
		)
		fmt.Println(Line)
		fmt.Print("Enter Choice: ")
		fmt.Scanf("%v \n", &choice)
		choiceInt, err := strconv.ParseInt(choice, 10, 64)
		if err != nil {
			fmt.Println("Invalid Input")
			filterGame()
			break
		} else if choiceInt == 3 {
			InitializeGame()
		} else {
			inputFilter(int(choiceInt))
			break
		}
	}

}

func deleteGame() {
	//input id then call to the delete method
}

func addGame() {
	inputGame()
}

func updateGame() {
	inputGame()
}
