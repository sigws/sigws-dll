package main

import (
	"C"
	"log"
	"net/http"
	"strconv"
)

var FlagToCallPause bool

//export StartListenAndServe
func StartListenAndServe(port int) {
	FlagToCallPause = false
	PORT := strconv.Itoa(port)
	http.Handle("/", http.FileServer(http.Dir("www")))
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	// Need a main function to make CGO compile package as C shared library
}
