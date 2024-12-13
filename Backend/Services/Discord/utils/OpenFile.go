package utils

import "os"

func OpenFile(filepath string) (b []byte, err error) {
	return os.ReadFile(filepath)
}
