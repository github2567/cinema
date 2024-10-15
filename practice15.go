package main

import (
	"fmt"
	
	// "github.com/PongthornP/cinema/movie"
	"github.com/github2567/cinema/ticket"
	"github.com/github2567/cinema/movie"
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
