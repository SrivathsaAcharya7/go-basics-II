package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SrivathsaAcharya7/golang-mongodb/router"
)

func main() {
	fmt.Println("This is the main page")
	r := router.Router()
	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":4000", r))
}
