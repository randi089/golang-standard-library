package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	// temp := s[i]
	// s[i] = s[j]
	// s[j] = temp

	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Randi", 26},
		{"Fatwa", 24},
		{"Fauzan", 25},
		{"Tito", 23},
		{"Taufiq", 22},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
