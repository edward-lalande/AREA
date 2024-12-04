package utils

import (
	"os"
)

func OpenJsonFile(filepath string) (map[string]interface{}, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return BytesToJson(b)
}
