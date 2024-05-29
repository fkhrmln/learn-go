package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createFile(name string, value string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0755)

	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(value)

	return nil
}

func readFile(name string) error {
	file, err := os.OpenFile(name, os.O_RDONLY, 0755)

	if err != nil {
		return err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	return nil
}

func writeFile(name string, value string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0755)

	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(value)

	return nil
}

func main() {
	// createFile("/home/fkhrmln/Downloads/example.txt", "Fakhri Maulana Ihsan\n"+"Rifky Ferdiansyah\n"+"Audry Elgalia\n")

	// readFile("/home/fkhrmln/Downloads/example.txt")

	writeFile("/home/fkhrmln/Downloads/example.txt", "Bobby Pratama\n")
}
