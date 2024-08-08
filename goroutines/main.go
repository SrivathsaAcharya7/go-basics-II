package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	// fmt.Println("This is main")
	// go greeter("Hello")
	// greeter("World")
	websiteList := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.gmail.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

//	func greeter(s string) {
//		for i := 0; i < 6; i++ {
//			time.Sleep(3 * time.Second)
//			fmt.Println(s)
//		}
//	}
func getStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Something went wrong")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for endpoint %s\n", result.StatusCode, endpoint)
	}

}
