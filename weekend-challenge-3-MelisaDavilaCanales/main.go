package main

import (
	"fmt"
)

type personTtypeOption int

const (
	UNSELECTED personTtypeOption = iota
	NORMAL_PERSON
	SPECIAL_PERSON
)

func main() {
	var personType personTtypeOption

	fmt.Printf("Welcome to the program!!!\n")

	for personType == UNSELECTED {
		fmt.Printf("Please select the type of person you are: \n")
		for key, value := range map[personTtypeOption]string{NORMAL_PERSON: "Normal Person", SPECIAL_PERSON: "Special Person"} {
			fmt.Printf("%d. %s\n", key, value)
		}
		fmt.Scanln(&personType)
	}
}
