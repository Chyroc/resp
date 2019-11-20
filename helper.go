package resp

import (
	"errors"
	"strconv"
)

var ErrKeyNotExist = errors.New("key not exist")

func stringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
