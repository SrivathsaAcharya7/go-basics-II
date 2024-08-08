package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {

	//IO and string conversion demo
	fmt.Println("This is the main function")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter ratings for our pizza")

	input, _ := reader.ReadString('\n')
	fmt.Printf("Thanks for rating our pizza %v", input)

	numRatings, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Printf("The error is %v", err)
	} else {
		fmt.Println("A new number will be added")
		numRatings = numRatings + 1
		fmt.Printf("Updated number is %v\n", numRatings)

	}

	//time and date demo
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

	//pointer demo
	pointerNumber := 23

	pointerNumberValue := &pointerNumber
	fmt.Println(pointerNumberValue)
	fmt.Println(*pointerNumberValue)

	*pointerNumberValue = *pointerNumberValue + 2
	fmt.Println(*pointerNumberValue)

	//sorting demo
	scores := make([]int, 4)
	scores[0] = 200
	scores[1] = 300
	scores[2] = 1000
	scores[3] = 500

	scores = append(scores, 700, 800, 900)

	fmt.Println(scores)
	sort.Ints(scores)
	fmt.Println(scores)

	proResult := proFunction(23, 44, 12, 3, 4, 5)
	fmt.Println(proResult)

	//defer keyword demo
	defer fmt.Println("Hello")
	fmt.Println("World")
	defer fmt.Println("Hi")

	//file reading and writing demo
	content := "This is the content that will get stored in the file"
	file, err := os.Create("./sampletextfile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("The lenght of the file is: ", length)
	defer file.Close()
	readFile("./sampletextfile.txt")

}

func readFile(file string) {
	dataBytes, err := os.ReadFile(file)
	checkNilError(err)
	fmt.Println("The contents of file are: ", string(dataBytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

// special function syntax
func proFunction(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}
