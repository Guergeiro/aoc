package reader

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func fileReader(path string) []string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Open(absPath)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	output := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return output
}

func ReadFileStrings(path string) []string {
	return fileReader(path)
}
