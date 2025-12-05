package reader

import (
	"bufio"
	"os"
	"path/filepath"
)

func ReadFileLines(path string) ([]string, error) {
	output := []string{}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return output, err
	}

	file, err := os.Open(absPath)

	if err != nil {
		return output, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return output, err
	}
	return output, nil
}
