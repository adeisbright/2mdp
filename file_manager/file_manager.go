package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadFromFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func CreateFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, errors.New("errror with creating file")
	}
	defer file.Close()
	return file, nil
}

func AddToFile(fileName string, content string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		fmt.Println(err)
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
	}
}
