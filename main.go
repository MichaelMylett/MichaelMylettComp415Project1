package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getFile() []string {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter filename (include file type, like .txt): ")
	fileOfChoice, err := inputReader.ReadString('\n')
	if err != nil {
		return nil
	}
	contents, err := os.ReadFile(fileOfChoice[:len(fileOfChoice)-1]) //len-1 to remove the invisible \n from the string
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	bigString := string(contents)
	allLines := strings.Split(bigString, "\n")
	return allLines
}

func function1() {
	//reads from file of choice
	allLines := getFile()
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Print every X line - X = ")
	X, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(nil)
	}
	Xint, err := strconv.Atoi(strings.TrimSpace(X)) //had to look up Atoi/TrimSpace and what library it was from, found through Google overview
	fmt.Println(Xint)
	if err != nil {
		fmt.Println(nil)
	}
	for i, line := range allLines {
		if i%Xint == 0 {
			fmt.Println(line)
		}
	}
	fmt.Println()
}
func function2() {
	//reads out bytes/characters in user input
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	userText, err := inputReader.ReadString('\n')
	if err != nil {
		println(err)
	}
	fmt.Println("Your text is", len(userText), "bytes.")
	count := 0
	for _, _ = range userText {
		count += 1
	}
	fmt.Println("Your text is", count, "characters long.")
	fmt.Println()
}

type Person struct {
	firstName    string
	age          int
	compSciMajor bool
}

func structChanger(person *Person) { //* to make it change the struct rather than a copy
	person.age = 21
	fmt.Println("That age is actually wrong. Here is the correct struct:", *person)
	return
}

func function3() {
	//use functions to make and modify struct
	//I had to look up pointer syntax, which I just read quickly in Google's overview, but no actual code was used
	person1 := Person{
		firstName:    "Michael",
		age:          20,
		compSciMajor: true,
	}
	fmt.Println("The struct includes a person's name, age, and whether they major in CompSci, in that order.")
	fmt.Println("Here is your struct:", person1)
	structChanger(&person1)
	fmt.Println("Here is your struct, again:", person1)
	fmt.Println()
}

func structChangerMk2(people []Person) {
	people[1].compSciMajor = false
	fmt.Println("Ryan switched to electrical engineering. Here is the correct list:", people)
	return
}

func function4() {
	//create list of structs using prev functions
	person1 := Person{
		firstName:    "Michael",
		age:          21,
		compSciMajor: true,
	}
	person2 := Person{
		firstName:    "Ryan",
		age:          19,
		compSciMajor: true,
	}
	people := []Person{person1, person2}
	fmt.Println("Here is your list of people:", people)
	fmt.Println("It will now be edited using an index of the list.")
	structChangerMk2(people)
	fmt.Println("Here is your list of people, again:", people)
	fmt.Println()
}
func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Pick one of the following options:\n" +
		"1. Print lines from a file in the project folder\n" +
		"2. Get byte/char count of a string\n" +
		"3. Create and alter a struct\n" +
		"4. Create a list of 2 structs and alter the second\n" +
		"5. End the program")
	fmt.Print("Input function option here: ")
	userAns, err := inputReader.ReadString('\n')
	if err != nil {
		println(err)
	}
	fmt.Println("You chose option", userAns)
	switch userAns {
	case "1\n":
		function1()
		main()
	case "2\n":
		function2()
		main()
	case "3\n":
		function3()
		main()
	case "4\n":
		function4()
		main()
	case "5\n":
		fmt.Println("Ending program...")
		return
	default:
		fmt.Println("You did not properly select any of the options. Start from the beginning.")
		main()
	}
}
