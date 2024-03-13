package game

// Type of commands
const (
	CMD_TYPE_PLACE_SHIP = iota
	CMD_TYPE_FIRE
)

// Type of command name
const (
	CMD_NAME_PLACE_SHIP = "PLACE_SHIP"
	CMD_NAME_FIRE       = "FIRE"
)

// Ship types
const (
	SHIP_TYPE_CARRIER = iota
	SHIP_TYPE_BATTLESHIP
	SHIP_TYPE_CRUISER
	SHIP_TYPE_SUBMARINE
	SHIP_TYPE_DESTROYER
)

// Ship size
const (
	SHIP_TYPE_SIZE_CARRIER    = 5
	SHIP_TYPE_SIZE_BATTLESHIP = 4
	SHIP_TYPE_SIZE_CRUISER    = 3
	SHIP_TYPE_SIZE_SUBMARINE  = 3
	SHIP_TYPE_SIZE_DESTROYER  = 2
)

// Ship type names
const (
	SHIP_NAME_CARRIER    = "Carrier"
	SHIP_NAME_BATTLESHIP = "Battleship"
	SHIP_NAME_CRUISER    = "Cruiser"
	SHIP_NAME_SUBMARINE  = "Submarine"
	SHIP_NAME_DESTROYER  = "Destroyer"
)

// Type of ships direction
const (
	SHIP_DIRECT_TYPE_RIGHT = iota
	SHIP_DIRECT_TYPE_DOWN
)

// Ship type names
const (
	SHIP_DIRECT_NAME_RIGHT = "right"
	SHIP_DIRECT_NAME_DOWN  = "down"
)

const (
	FIRE_STATUS_HIT = iota
	FIRE_STATUS_MISS
	FIRE_STATUS_SUNK
	FIRE_STATUS_OVER
)

const (
	PLACED_STATUS_TEMPLATE = "Placed %s"
)

const (
	FIRE_STATUS_NAME_HIT  = "Hit"
	FIRE_STATUS_NAME_MISS = "Miss"
	FIRE_STATUS_NAME_SUNK = "You sunk my %s!"
	FIRE_STATUS_NAME_OVER = "Game Over"
)

const (
	SHIP_CELL_UP = iota
	SHIP_CELL_DOWN
	SHIP_CELL_LEFT
	SHIP_CELL_RIGHT
)

type command struct {
	// type of command
	cmdType int
	//
	placeShip placeShipCmd
	//
	fire fireCmd
}

type placeShipCmd struct {
	shipType  int
	direction int
	// start row index
	r int
	// start column index
	c int
}

type fireCmd struct {
	// row index
	r int
	// column index
	c int
}

type shipStatus int

type shipType int

func getShipNameByType(sType shipType) string {
	switch sType {
	case SHIP_TYPE_CARRIER:
		return SHIP_NAME_CARRIER
	case SHIP_TYPE_BATTLESHIP:
		return SHIP_NAME_BATTLESHIP
	case SHIP_TYPE_CRUISER:
		return SHIP_NAME_CRUISER
	case SHIP_TYPE_SUBMARINE:
		return SHIP_NAME_SUBMARINE
	case SHIP_TYPE_DESTROYER:
		return SHIP_NAME_DESTROYER
	}
	return ""
}

func getShipSize(shipType int) int {
	switch shipType {
	case SHIP_TYPE_CARRIER:
		return SHIP_TYPE_SIZE_CARRIER
	case SHIP_TYPE_BATTLESHIP:
		return SHIP_TYPE_SIZE_BATTLESHIP
	case SHIP_TYPE_CRUISER:
		return SHIP_TYPE_SIZE_CRUISER
	case SHIP_TYPE_SUBMARINE:
		return SHIP_TYPE_SIZE_SUBMARINE
	case SHIP_TYPE_DESTROYER:
		return SHIP_TYPE_SIZE_DESTROYER
	}
	return 0
}
