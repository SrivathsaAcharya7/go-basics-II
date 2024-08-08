package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const oldUrl = "https://courses.learncodeonline.in/"
const newTestUrl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=sdadfsd2zxc3"

func main() {
	fmt.Println("This is handling-web-request using golang")

	response, err := http.Get(oldUrl)

	if err != nil {
		panic(err)
	}
	fmt.Printf("The response type of the request is %T\n", response)
	defer response.Body.Close()
	// dataBytes, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("The contents of the response body are: ", string(dataBytes))
	fmt.Println(newTestUrl)

	result, _ := url.Parse(newTestUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Port())
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	queryParams := result.Query()
	fmt.Printf("The type of the queryParams is %T\n", queryParams)

	fmt.Println(queryParams["coursename"])

	constructUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/newcss",
		RawPath: "user=srivathsa",
	}
	constructedNewUrl := constructUrl.String()
	fmt.Println(constructedNewUrl)
}
