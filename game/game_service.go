package game

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// Update Game
// need to update price, stocks, title and genre
func UpdateGame(id int, newGame *Game) (Game, error) {
	game, err := FilterById(id)
	fmt.Println(err)

	if err == nil {
		game.Title = newGame.Title
		game.Description = newGame.Description
		game.Genre = newGame.Genre
		game.Price = newGame.Price
		game.Stocks = newGame.Stocks

		updatedGame := &Game{
			Title:       game.Title,
			Description: game.Description,
			Genre:       game.Genre,
			Price:       game.Price,
			Stocks:      game.Stocks,
		}

		Games[id] = *updatedGame
	} else {
		return Game{}, fmt.Errorf("no game found with id %d", id)
	}

	return *newGame, nil
}

// VALIDATION NALANGS
func AddGame(title string, description string, genre string, price float64, stocks uint64) {
	newGame := &Game{
		Title:       title,
		Description: description,
		Genre:       genre,
		Price:       price,
		Stocks:      stocks,
	}
	id := len(Games) + 1
	Games[id] = *newGame
}

// Delete the Game From Input title or genre
func DeleteGame(id int) error {
	_, err := FilterById(id)

	if err != nil {
		return fmt.Errorf("no game with id %d found", id)
	}
	delete(Games, id)
	return nil
}

func FilterById(id int) (Game, error) {
	game := Games[id]

	if game.Title != "" {
		return game, nil
	} else {
		return Game{}, fmt.Errorf("no game found with id %d", id)
	}
}

// Return list of Games if Genre is Found else not found
func FilterByGenre(genre string) (map[int]Game, error) {
	genre = TrimSpaces(genre)
	listOfGamesByGenre := make(map[int]Game, len(Games))

	for id, game := range Games {
		game.Genre = TrimSpaces(game.Genre)
		if strings.EqualFold(game.Genre, genre) {
			listOfGamesByGenre[id] = game
		}
	}

	if len(listOfGamesByGenre) > 0 {
		return listOfGamesByGenre, nil
	} else {
		return nil, fmt.Errorf("no games found with genre %s", genre)
	}
}

// NAGANA KANA HANGGANG DITO
// Return A single Game when Found
func FilterByTitle(title string) (map[int]Game, error) {
	title = TrimSpaces(title)
	gameByTitle := make(map[int]Game, 1)

	for id, game := range Games {
		game.Title = TrimSpaces(game.Title)
		if strings.EqualFold(game.Title, title) {
			gameByTitle[id] = game
			return gameByTitle, nil
		}
	}

	return nil, fmt.Errorf("no game found with title %s", title)
}

func SortGameByPriceOrStocks(priceOrStocks string) map[int]Game {
	type mapRepl struct {
		key int
		val Game
	}

	var games []mapRepl

	for id, game := range Games {
		games = append(games, mapRepl{id, game})
	}

	if priceOrStocks == "price" {
		sort.Slice(games, func(i, j int) bool {
			return games[i].val.Price < games[j].val.Price
		})
	} else {
		sort.Slice(games, func(i, j int) bool {
			return float64(games[i].val.Stocks) < float64(games[j].val.Stocks)
		})
	}

	sortedMap := make(map[int]Game)
	for _, sortedGame := range games {
		sortedMap[sortedGame.key] = sortedGame.val
	}
	return sortedMap
}

func ExitGame() {
	os.Exit(0)
}
