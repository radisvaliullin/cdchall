package game

import (
	"bufio"
	"errors"
	"io"
	"log"
	"strconv"
	"strings"
)

// Parse input line by line and return parsed command
// For each call return one command
// last error for io error for catch EOF
func readNextCommandFromInSource(inSrc *bufio.Reader) (command, error, error) {

	var line string
	var ioErr error
	// read command line by line
	// if empty line skip
	// if command break and return command
	for {
		line, ioErr = inSrc.ReadString('\n')
		if ioErr != nil && !errors.Is(ioErr, io.EOF) {
			log.Printf("input parser: read line error: %v", ioErr)
			return command{}, nil, ioErr
		}

		// trim empty chars
		line = strings.TrimSpace(line)
		// empty line skip
		if len(line) == 0 {
			if ioErr == nil {
				continue
			}
			// if EOF return
			return command{}, nil, ioErr
		}

		// parse single command
		cmd, cmdErr := commandParse(line)
		if cmdErr != nil {
			log.Printf("input parser: parse next command error: %v", cmdErr)
			return command{}, cmdErr, ioErr
		}
		return cmd, nil, ioErr
	}
}

// parse command
func commandParse(cmdStr string) (command, error) {
	cmd := command{}

	cmdParams := strings.Fields(cmdStr)
	// if len of command args not 2 (for fire) or not 4 (for place ship) return error
	if !(len(cmdParams) == 2 || len(cmdParams) == 4) {
		err := errors.New("command parse: args len should be 2 or 4")
		return cmd, err
	}

	// command type parsing
	switch cmdParams[0] {
	case CMD_NAME_PLACE_SHIP:
		cmd.cmdType = CMD_TYPE_PLACE_SHIP
		if len(cmdParams) != 4 {
			err := errors.New("command parse: place ship command args len should be 4")
			return cmd, err
		}
		shipType, err := parseShipType(cmdParams[1])
		if err != nil {
			return cmd, err
		}
		cmd.placeShip.shipType = shipType
		direct, err := parseShipDirection(cmdParams[2])
		if err != nil {
			return cmd, err
		}
		cmd.placeShip.direction = direct
		r, c, err := parsePosition(cmdParams[3])
		if err != nil {
			return cmd, err
		}
		cmd.placeShip.r = r
		cmd.placeShip.c = c
	case CMD_NAME_FIRE:
		cmd.cmdType = CMD_TYPE_FIRE
		r, c, err := parsePosition(cmdParams[1])
		if err != nil {
			return cmd, err
		}
		cmd.fire.r = r
		cmd.fire.c = c
	default:
		err := errors.New("command parse: unknown command type")
		return cmd, err
	}
	return cmd, nil
}

// returns row index, column index and error
func parsePosition(posStr string) (int, int, error) {
	r := 0
	c := 0
	if !(len(posStr) == 2 || len(posStr) == 3) {
		err := errors.New("parse position: wrong len, should be 2 or 3")
		return 0, 0, err
	}
	switch posStr[0] {
	case 'A':
		r = 0
	case 'B':
		r = 1
	case 'C':
		r = 2
	case 'D':
		r = 3
	case 'E':
		r = 4
	case 'F':
		r = 5
	case 'G':
		r = 6
	case 'H':
		r = 7
	case 'I':
		r = 8
	case 'J':
		r = 9
	default:
		err := errors.New("parse position: wrong row value")
		return 0, 0, err
	}
	colPosStr := posStr[1:]
	c, err := strconv.Atoi(colPosStr)
	if err != nil {
		return 0, 0, err
	}
	// decrement to 1 for numerate from 0 to 9
	c--
	if c < 0 || c > 9 {
		err := errors.New("parse position: wrong column value")
		return 0, 0, err
	}

	return r, c, nil
}

func parseShipType(shipType string) (int, error) {
	t := 0
	switch shipType {
	case SHIP_NAME_CARRIER:
		t = SHIP_TYPE_CARRIER
	case SHIP_NAME_BATTLESHIP:
		t = SHIP_TYPE_BATTLESHIP
	case SHIP_NAME_CRUISER:
		t = SHIP_TYPE_CRUISER
	case SHIP_NAME_SUBMARINE:
		t = SHIP_TYPE_SUBMARINE
	case SHIP_NAME_DESTROYER:
		t = SHIP_TYPE_DESTROYER
	default:
		err := errors.New("parse ship type: unknown ship type")
		return 0, err
	}
	return t, nil
}

func parseShipDirection(shipDirect string) (int, error) {
	d := 0
	switch shipDirect {
	case SHIP_DIRECT_NAME_RIGHT:
		d = SHIP_DIRECT_TYPE_RIGHT
	case SHIP_DIRECT_NAME_DOWN:
		d = SHIP_DIRECT_TYPE_DOWN
	default:
		err := errors.New("parse ship direction: unknown ship direction")
		return 0, err
	}
	return d, nil
}
