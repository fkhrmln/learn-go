package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("Fakhri Maulana Ihsan\n" + "Rifky Ferdiansyah\n" + "Audry Elgalia\n")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	writer := bufio.NewWriter(os.Stdout)

	writer.WriteString("Fakhri Maulana Ihsan\n")
	writer.WriteString("Rifky Ferdiansyah\n")
	writer.WriteString("Audry Elgalia\n")

	writer.Flush()
}
