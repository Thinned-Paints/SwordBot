package main

import (
	"fmt"
	"bufio"
	"os"
)

func difficultySelect() string {
	// This function gets the difficulty from a user

	difficultyOptions := [3]string{"Easy", "Medium", "Hard"}

	fmt.Println(difficultyOptions)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please select your difficulty")
	fmt.Println("-----------------------------")
	for i := 0; i < len(difficultyOptions); i++ {
		fmt.Println(difficultyOptions[i])
	}
	fmt.Println("-----------------------------")
	var difficulty, _ = reader.ReadString('\n')
	fmt.Println("-----------------------------")
	
	return difficulty
}

func main(){

	fmt.Println("Welcome to Swordbot, this is the terminal window for development and testing")
	fmt.Println("----------------------------------------------------------------------------")

	var difficulty = difficultySelect()
	fmt.Println(difficulty)
}

