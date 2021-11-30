package convert

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

type (
	ToByter interface {
		ToBytes() ([]byte, error)
	}
)

func FloatToString(num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(num, 'f', 64, 64)
}

func IntToString(num int) string {
	// to convert a float number to a string
	return strconv.Itoa(num)
}

func StringToFloat(num string) (float64, error) {
	return strconv.ParseFloat(num, 64)
}

func MD5HashingToString(convert string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(convert)))
}

func Unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
