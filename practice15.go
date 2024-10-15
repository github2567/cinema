package main

import (
	"fmt"
	
	// "github.com/PongthornP/cinema/movie"
	"github.com/PongthornP/cinema/ticket"
	"github.com/PongthornP/cinema/movie"
)

func init() {
	fmt.Println("init: main")
}

func main() {
	movieName := movie.FindMovieName("tt4154796")
	fmt.Println(movieName)
	movie.ReviewMovie(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
