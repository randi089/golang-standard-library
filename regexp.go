package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`r([a-z])i`)

	fmt.Println(regex.MatchString("rUi"))
	fmt.Println(regex.MatchString("rai"))
	fmt.Println(regex.MatchString("rui"))

	fmt.Println(regex.FindAllString("rui rai rei rau rou rii", 10))
}
