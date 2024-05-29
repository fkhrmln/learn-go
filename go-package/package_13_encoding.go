package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	encoded := base64.StdEncoding.EncodeToString([]byte("Fakhri Maulana Ihsan"))

	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(decoded))
	}

	csvString := "Fakhri,Maulana,Ihsan\n" + "Rifky,Ferdiansyah\n" + "Audry,Elgalia\n"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

	writer := csv.NewWriter(os.Stdout)

	writer.Write([]string{"Fakhri", "Maulana", "Ihsan"})
	writer.Write([]string{"Audry", "Elgalia"})
	writer.Write([]string{"Rifky", "Ferdiansyah"})

	writer.Flush()
}
