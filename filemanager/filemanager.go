package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLine(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {

		return nil, errors.New("Fail to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {

		file.Close()
		return nil, errors.New("fail to read line in file")
	}
	file.Close()
	return lines, nil
}
