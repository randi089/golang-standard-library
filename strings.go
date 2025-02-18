package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Randi Febriadi", "Randi"))
	fmt.Println(strings.Split("Randi Febriadi", " "))
	fmt.Println(strings.ToLower("Randi Febriadi"))
	fmt.Println(strings.ToUpper("Randi Febriadi"))
	fmt.Println(strings.Trim("  Randi Febriadi  ", " "))
	fmt.Println(strings.ReplaceAll("Randi Febriadi Randi Hebat", "Randi", "Fatwa"))
}
