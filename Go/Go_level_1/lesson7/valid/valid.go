package valid

import (
	"math"
	"net/url"
)

//IsPort -
func IsPort(port float64) bool {
	return uint64(port) < math.MaxUint16
}

//IsURL  -
func IsURL(urlStr string) (bool, error) {
	_, err := url.Parse(urlStr)
	if err == nil {
		return true, err
	}
	return false, err
}
