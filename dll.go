package main

import (
	"C"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var StartTime string

func nowtime() string {
	now := time.Now()
	year := strconv.Itoa(now.Year())
	month := strconv.Itoa(int(now.Month()))
	day := strconv.Itoa(now.Day())
	hour := strconv.Itoa(now.Hour())
	minute := strconv.Itoa(now.Minute())
	second := strconv.Itoa(now.Second())
	return "[" + year + "/" + month + "/" + day + " " + hour + ":" + minute + ":" + second + "] "
}

func starttime() string {
	now := time.Now()
	year := strconv.Itoa(now.Year())
	month := strconv.Itoa(int(now.Month()))
	day := strconv.Itoa(now.Day())
	hour := strconv.Itoa(now.Hour())
	minute := strconv.Itoa(now.Minute())
	second := strconv.Itoa(now.Second())
	return "sig-logs/[" + year + "-" + month + "-" + day + " " + hour + "-" + minute + "-" + second + "].log"
}

func printlog(content string) {
	f, err := os.OpenFile(StartTime, os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("open file error:", err)
		return
	}
	_, err = f.WriteString(nowtime() + content + "\n")
	if err != nil {
		log.Println("close file error:", err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(nowtime(), err)
		}
	}(f)
	fmt.Println(nowtime(), content)
}

var FlagToCallPause bool

//export StartListenAndServe
func StartListenAndServe(port int) {
	StartTime = starttime()
	err := os.MkdirAll("sig-logs", os.ModePerm)
	if err != nil {
		fmt.Println(nowtime(), err)
	} else {
		printlog("Create logs folder successfully!")
	}
	FlagToCallPause = false
	PORT := strconv.Itoa(port)
	printlog("Start listen on " + PORT)
	http.Handle("/", http.FileServer(http.Dir("www")))
	err = http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	// Need a main function to make CGO compile package as C shared library
}
