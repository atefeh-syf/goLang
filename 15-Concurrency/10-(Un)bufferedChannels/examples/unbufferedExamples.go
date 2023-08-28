package examples

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func Emaple1() {
	content := make(chan string)
	for i := 0; i < 100; i++ {
		go GetHttpRequest(content, i+1)
	}

	for i := 0; i < 100; i++ {
		response := <-content
		println(response)
	}
}

func GetHttpRequest(content chan string, index int) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(index))

	if err != nil {
		panic(err)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	PrintLnWithTime("Before set content")
	content <- string(responseBody)
	PrintLnWithTime("After set content")
}

func PrintLnWithTime(args ...any) {
	fmt.Printf("Time: %s, %v\n", time.Now().Format(time.RFC3339Nano), args)
}
