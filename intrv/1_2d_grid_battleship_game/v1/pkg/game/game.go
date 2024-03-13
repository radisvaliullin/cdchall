package game

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	GRID_ROW_SIZE = 10
	GRID_COL_SIZE = 10

	GRID_OCEAN = 0
	GRID_SHIP  = 1

	// ship cell value number base for store ship HIT status, TYPE, NUMBER
	SHIP_HIT_BASE    = 10
	SHIP_TYPE_BASE   = 100
	SHIP_NUMBER_BASE = 1000

	defaultInBuffSize = 1024
)

type Config struct {
	// Input file path
	InPath string
}

type Game struct {
	config Config

	// game grid 10x10 size
	// value
	// Base 0: 0 - ocean, 1 ship
	// Base 10: status of hit; 0x - no hit, 1x - hit
	// Base 100: ship type: 0xx - carrier for example, 1xx - battleship
	// Base 1000: ship number: 0xxx - ship #0, 1xxx - ship #1
	grid [][]int
	//
	shipCount int
}

func New(config Config) *Game {
	g := &Game{
		config: config,
		grid:   make([][]int, GRID_ROW_SIZE),
	}
	for i, _ := range g.grid {
		g.grid[i] = make([]int, GRID_COL_SIZE)
	}
	return g
}

func (g *Game) Run() error {

	// inFile can be any io.Reader for example os.Stdin
	inFile, err := os.Open(g.config.InPath)
	if err != nil {
		log.Printf("game: open input file error: %v", err)
		return err
	}
	defer inFile.Close()

	if err := g.Play(inFile); err != nil {
		log.Printf("game: play error: %v", err)
	}

	return nil
}

func (g *Game) Play(inSrc io.Reader) error {

	buffInSrc := bufio.NewReaderSize(inSrc, defaultInBuffSize)

	isPlaceRange := false
	isFireRange := false

	// read input source
	var ioErr error
	for {
		if ioErr != nil {
			log.Printf("game: play: read next command error: %v", ioErr)
			return ioErr
		}
		var cmd command
		var cmdErr error
		cmd, cmdErr, ioErr = readNextCommandFromInSource(buffInSrc)
		// if io error we need return
		// but if io.EOF we need first handle command and after return, see ioErr check in beginning of for loop
		if ioErr != nil && !errors.Is(ioErr, io.EOF) {
			log.Printf("game: play: read next command error: %v", ioErr)
			return ioErr
		} else if cmdErr != nil {
			log.Printf("game: play: read next command error: %v", cmdErr)
			continue
		}

		// validate command execution order
		if !isPlaceRange && !isFireRange && cmd.cmdType == CMD_TYPE_FIRE {
			err := errors.New("game: play: fire command before done with place")
			return err
		}
		if isFireRange && cmd.cmdType == CMD_TYPE_PLACE_SHIP {
			err := errors.New("game: play: place command after start fire commands")
			return err
		}

		// handle command
		switch cmd.cmdType {
		case CMD_TYPE_PLACE_SHIP:
			isPlaceRange = true

			if err := g.placeShip(
				cmd.placeShip.shipType, cmd.placeShip.direction,
				cmd.placeShip.r, cmd.placeShip.c,
			); err != nil {
				log.Printf("play: place ship command error: %v", err)
			} else {
				printPlaceStatus(shipType(cmd.placeShip.shipType))
			}
		case CMD_TYPE_FIRE:
			isFireRange = true

			r, c := cmd.fire.r, cmd.fire.c
			shipStatus, shipType := g.makeFire(r, c)
			printFireStatus(shipStatus, shipType)
		}
	}
}

func (g *Game) placeShip(shipType, direct, r, c int) error {
	size := getShipSize(shipType)
	if ok := isShipSpace(size, direct, r, c); !ok {
		err := errors.New("place ship: ship len out of space")
		return err
	}
	if ok := isShipCollision(g.grid, size, direct, r, c); ok {
		err := errors.New("place ship: ships collisions")
		return err
	}
	// place ship
	// ship value
	// value includes ship flag, ship type, ship number
	cellValue := GRID_SHIP + shipType*SHIP_TYPE_BASE + g.shipCount*SHIP_NUMBER_BASE
	// set values
	switch direct {
	case SHIP_DIRECT_TYPE_RIGHT:
		for i := c; i < (c + size); i++ {
			g.grid[r][i] = cellValue
		}
	case SHIP_DIRECT_TYPE_DOWN:
		for i := r; i < (r + size); i++ {
			g.grid[i][c] = cellValue
		}
	}
	g.shipCount++
	return nil
}

func (g *Game) makeFire(r, c int) (shipStatus, shipType) {
	// not ocean
	if g.grid[r][c] > 0 {
		// check if cell is ship and not hit
		// we need only last two base
		if g.grid[r][c]%SHIP_TYPE_BASE == 1 {
			// if not hit mark as hit
			g.grid[r][c] = g.grid[r][c] + SHIP_HIT_BASE
		}
		st := checkShipStatus(g.grid, r, c)
		switch st {
		case FIRE_STATUS_SUNK:
			g.shipCount--
			if g.shipCount == 0 {
				st = FIRE_STATUS_OVER
			}
		}
		sType := g.grid[r][c] % SHIP_NUMBER_BASE / SHIP_TYPE_BASE
		return st, shipType(sType)
	}
	return FIRE_STATUS_MISS, 0
}

func isShipSpace(size, direct, r, c int) bool {
	switch direct {
	case SHIP_DIRECT_TYPE_RIGHT:
		if (c + size) > GRID_COL_SIZE {
			return false
		}
	case SHIP_DIRECT_TYPE_DOWN:
		if (r + size) > GRID_ROW_SIZE {
			return false
		}
	}
	return true
}

func isShipCollision(grid [][]int, size, direct, r, c int) bool {
	switch direct {
	case SHIP_DIRECT_TYPE_RIGHT:
		for i := c; i < (c + size); i++ {
			if grid[r][i] > 0 {
				return true
			}
		}
	case SHIP_DIRECT_TYPE_DOWN:
		for i := r; i < (r + size); i++ {
			if grid[i][c] > 0 {
				return true
			}
		}
	}
	return false
}

// check ship status
// check all cell exclude input cell (r, c)
// if at least one cell of ship not hit ship still sail
// then return ship status as HIT because input cell hit
// if all cells hit return SUNK
func checkShipStatus(grid [][]int, r, c int) shipStatus {

	// interate all cell of ship
	bcells := getShipBorderCells(grid, r, c)
	for _, c := range bcells {
		// next row
		nr := c[0]
		// next col
		nc := c[1]
		// next direction
		nd := c[2]
		for {
			// if not hit return
			if grid[nr][nc]%SHIP_TYPE_BASE == 1 {
				return FIRE_STATUS_HIT
			}
			// next
			var ok bool
			nr, nc, ok = getShipBorderCell(grid, nr, nc, nd)
			// if not next cell break loop
			if !ok {
				break
			}
		}
	}
	return FIRE_STATUS_SUNK
}

// returns list of [(row, cell, direction), ...]
func getShipBorderCells(grid [][]int, r, c int) [][]int {
	// ship can have only 2 border cell maximum
	bcells := make([][]int, 0, 2)

	// up
	ur, uc, ok := getShipBorderCell(grid, r, c, SHIP_CELL_UP)
	if ok {
		bcells = append(bcells, []int{ur, uc, SHIP_CELL_UP})
	}
	// down
	dr, dc, ok := getShipBorderCell(grid, r, c, SHIP_CELL_DOWN)
	if ok {
		bcells = append(bcells, []int{dr, dc, SHIP_CELL_DOWN})
	}
	// left
	lr, lc, ok := getShipBorderCell(grid, r, c, SHIP_CELL_LEFT)
	if ok {
		bcells = append(bcells, []int{lr, lc, SHIP_CELL_LEFT})
	}
	// right
	rr, rc, ok := getShipBorderCell(grid, r, c, SHIP_CELL_RIGHT)
	if ok {
		bcells = append(bcells, []int{rr, rc, SHIP_CELL_RIGHT})
	}
	return bcells
}

// returns row, col of next cell for direction or false ok flag if cell not find
func getShipBorderCell(grid [][]int, r, c, d int) (int, int, bool) {
	br := r
	bc := c

	switch d {
	case SHIP_CELL_UP:
		br--
	case SHIP_CELL_DOWN:
		br++
	case SHIP_CELL_LEFT:
		bc--
	case SHIP_CELL_RIGHT:
		bc++
	}
	if br < 0 || br >= GRID_ROW_SIZE {
		return 0, 0, false
	}
	if bc < 0 || bc >= GRID_COL_SIZE {
		return 0, 0, false
	}

	cv := grid[r][c]
	bv := grid[br][bc]
	// check that border cell is ship and same ship
	if bv > 0 && (cv/SHIP_NUMBER_BASE) == (bv/SHIP_NUMBER_BASE) {
		return br, bc, true
	}
	return 0, 0, false
}

func printPlaceStatus(sType shipType) {
	out := fmt.Sprintf(PLACED_STATUS_TEMPLATE, getShipNameByType(sType))
	fmt.Printf("%s\n", out)
}

func printFireStatus(st shipStatus, sType shipType) {
	switch st {
	case FIRE_STATUS_MISS:
		fmt.Printf("%s\n", FIRE_STATUS_NAME_MISS)
	case FIRE_STATUS_HIT:
		fmt.Printf("%s\n", FIRE_STATUS_NAME_HIT)
	case FIRE_STATUS_OVER:
		fmt.Printf("%s\n", FIRE_STATUS_NAME_HIT)
		out := fmt.Sprintf(FIRE_STATUS_NAME_SUNK, getShipNameByType(sType))
		fmt.Printf("%s\n", out)
		fmt.Printf("%s\n", FIRE_STATUS_NAME_OVER)
	case FIRE_STATUS_SUNK:
		fmt.Printf("%s\n", FIRE_STATUS_NAME_HIT)
		out := fmt.Sprintf(FIRE_STATUS_NAME_SUNK, getShipNameByType(sType))
		fmt.Printf("%s\n", out)
	}
}
