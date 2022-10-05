package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	data, _ := os.ReadFile("Hangman.txt")

	var tableau []string
	var word string

	for index, char := range data {
		if char == 13 {
			tableau = append(tableau, word)
			word = ""
		} else if char != 10 {
			word += string(char)
		}
		if index == len(data)-1 {
			tableau = append(tableau, word)
		}
	}

	fmt.Println(tableau[0])

	fmt.Println(tableau[len(tableau)-1])
	var a string
	for i := 1; i < len(tableau); i++ {
		if tableau[i-1] == "before" {
			a = tableau[i]

			break
		}

	}

	aInt, _ := strconv.Atoi(a)
	fmt.Println(tableau[aInt-1]) // car tableau commence a 0

	for i := 0; i < len(tableau); i++ {
		if tableau[i] == string("now ") {

			for _, aretab := range tableau[i-1] {

				if aretab == 114 {

					aretab = aretab / 38

					fmt.Println(tableau[aretab-1])
				}

			}

		}
	}

	rand.Seed(time.Now().UnixNano())

	fmt.Println("La reponce est : ", rand.Int())

}
