package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "randi,febriadi,hebat\n" +
		"khairul,fatwa,mantap\n" +
		"muhammad,fauzan,gg"
	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}
