package app

import "fmt"

type Tetromino int

func (r Tetromino) IsI() bool {
	return fmt.Sprintf("%s", r) == "I"
}

func (r Tetromino) String() string {
	res := ""
	switch r {
	case 0:
		res = "I"
	case 1:
		res = "S"
	case 2:
		res = "Z"
	case 3:
		res = "T"
	case 4:
		res = "L"
	case 5:
		res = "J"
	case 6:
		res = "O"
	}
	return res
}
