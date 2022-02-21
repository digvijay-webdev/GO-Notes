// Everything is stored in RAM, data will be deleted automatically once the app is closed

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

var header string = `
-------------------------------------------
                  Notes APP
-------------------------------------------
Commands:
  use "c" to create a new item
  use "d" to delete an item
  use "v" to view all the items
-------------------------------------------
`

// for global scope
var messages = []string{}

func main() {
	clear()
	fmt.Println(header)
	// infinite loop started
	for {
		var val string = proecessInput()
		if val == "c" {
			// run create function
			create()
		} else if val == "d" {
			// run delete function
			delete()
		} else if val == "v" {
			// run view function
			view()
		} else {
			clear()
			fmt.Println(header)
			fmt.Println("Invalid command entered..")
		}
	}
}

func clear() {
	screen := exec.Command("clear")
	screen.Stdout = os.Stdout
	screen.Run()
}

func proecessInput() (v string) {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input = scanner.Text()
	}
	return input
}

func localStore(message string) {
	// coming from global scope
	messages = append(messages, message)
	clear()
	fmt.Println("Note Created successfully..")
	fmt.Println(header)
}

func create() {
	clear()
	fmt.Println(header)
	fmt.Println("Enter a note & press enter to save:")

	var noteInput string = proecessInput()
	localStore(noteInput)
}

func view() {
	clear()
	if len(messages) <= 0 {
		fmt.Println("You do not have any notes yet..")
		fmt.Println(header)
	} else {
		fmt.Printf("You've %v note(s)\n", len(messages))
		for _, val := range messages {
			fmt.Printf("- %v \n", val)
		}
		fmt.Println("\nPress 'Enter' to go back")
	}
}

func delete() {
	clear()
	fmt.Println("Please, type the index number of note you want to delete.")
	if len(messages) <= 0 {
		clear()
		fmt.Println("Sorry, you do not have any notes yet..")
		fmt.Println(header)
	} else {
		for index, val := range messages {
			fmt.Printf("Index %v - %v \n", index, val)
		}

		fmt.Println("Input Required:")
		var indexInput int
		fmt.Scanln(&indexInput)

		messages = append(messages[:indexInput], messages[indexInput+1:]...)
		view()
	}
}

/* https://digvijay.tech */
