package main

import (
	"fmt"
	"time"
)

// exercise 1
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

// exercise 2
func UpdateSlice(slice []string, word string) {
	size := len(slice)
	slice[size-1] = word
	fmt.Println("In Update: ", slice)
}

func GrowSlice(slice []string, word string) {
	slice = append(slice, word)
	fmt.Println("In Grow: ", slice)
}
func main() {
	MakePerson("Kunle", "Eniola", 27)
	MakePersonPointer("Jesus", "Christ", 33)
	// if stored on the stack, memory address returned from MakePersonPointer will no longer
	// be accessible as MakePersonPointer frame is gone

	words := []string{"In", "the", "beningging"}
	fmt.Println("Before update: ", words)
	UpdateSlice(words, "God")
	fmt.Println("After update: ", words)
	GrowSlice(words, "created")
	fmt.Println("After Grow: ", words)
	// Due to the copy of the slice passed into the GrowSLice function changing in the length,
	// its changes are not reflected in the original slice in main. Only changes that will occur
	// in the original slice are ones made to elements in the original length.

	// exercise 3
	startTime := time.Now()
	fmt.Println("Time Started: ", startTime)

	i := 0
	// personSlice := []Person{}
	personSlice := make([]Person, 0, 10_000_000)
	for i < 10_000_000 {
		personSlice = append(personSlice, Person{"", "", 21})
		i++
	}
	endTime := time.Now()
	fmt.Println("Duration: ", endTime.Sub(startTime))
	//running with a capacity of 10 million versus appending 10 million items
	// results in only one garbage collection and reduces execution time by
	// almost 90%

}
