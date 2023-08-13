package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var TodoList = []string{}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go GetToDo(i+1, &wg)
	}
	wg.Wait()
	fmt.Printf("%v", TodoList)
}

func GetToDo(id int, wg *sync.WaitGroup) {
	GetUrl("https://jsonplaceholder.typicode.com/todos/"+strconv.Itoa(id), wg)
}

func GetUrl(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	responseBody, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	TodoList = append(TodoList, string(responseBody))

}
