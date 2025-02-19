package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"randi", "febriadi", "hebat"})
	_ = writer.Write([]string{"khairul", "fatwa", "mantap"})
	_ = writer.Write([]string{"muhammad", "fauzan", "gg"})

	writer.Flush()
}
