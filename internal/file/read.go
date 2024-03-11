package file

import (
	"os"
)

func Read(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}
