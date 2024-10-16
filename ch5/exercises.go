package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// exercise 1
var (
	add = func(i, j int) (int, error) { return i + j, nil }
	sub = func(i, j int) (int, error) { return i - j, nil }
	mul = func(i, j int) (int, error) { return i * j, nil }
	div = func(i, j int) (int, error) {
		if j == 0 {
			return 0, errors.New("DIVISION BY ZERO")
		}
		return i / j, nil
	}
)

// exercise 2
func fileLen(file string) (length int, err error) {
	fil, err := os.Open(file)

	defer fil.Close()
	if err != nil {
		return 0, err
	}
	total := 0
	buff := make([]byte, 2048)
	for {
		n, err := fil.Read(buff)
		total += n
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
	}
	return total, err
}

// exercise 3
func prefixer(s string) func(string) string {
	return func(s2 string) string {
		return s + " " + s2
	}
}
func main() {
	// exercise 1
	// ------------------
	ans, err := div(1, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Answer: ", ans)
	}

	// exercise 2
	// ----------------
	fmt.Println(fileLen("exercises.go"))
	fmt.Println(fileLen("doesnotexist.go"))

	//exercise 3
	//--------------
	letsGoPrefixer := prefixer("LETS GOOO")
	fmt.Println(letsGoPrefixer("MARY!"))
	fmt.Println(letsGoPrefixer("JOSEPH!"))

}
