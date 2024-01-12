package game

type Game struct {
	Title       string
	Description string
	Genre       string
	Price       float64
	Stocks      uint64
}

var Games = map[int]Game{
	1: {Title: "Mario", Description: "Good", Genre: "Jump", Price: 498, Stocks: 3},
	2: {Title: "Dora", Description: "Good", Genre: "Adventure", Price: 10.25, Stocks: 10},
}
