package main

import (
	"flag"
	"log"

	"github.com/radisvaliullin/cdchall/intrv/1_2d_grid_battleship_game/v1/pkg/game"
)

func main() {

	// app config flags
	inPathFlag := flag.String("in", "./inputs/basic_inputs", "input file path")
	flag.Parse()

	// app config
	conf := game.Config{
		InPath: *inPathFlag,
	}

	// init and run game
	g := game.New(conf)
	if err := g.Run(); err != nil {
		log.Fatalf("game: run error: %v", err)
	}

}
