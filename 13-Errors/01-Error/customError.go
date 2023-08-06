package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HTTPError struct {
	StatusCode int
	Message    string
}

func main() {
	response, err := GetRequest("")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)

	response, err = GetRequest("https://dummyjson3123.com/products/categories")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
}

func (error HTTPError) Error() string {
	return fmt.Sprintf("Http error occurred: %d %s", error.StatusCode, error.Message)
}

func NewHttpError(statusCode int, message string) *HTTPError {
	return &HTTPError{StatusCode: statusCode, Message: message}
}

func GetRequest(url string) (string, error) {
	if len(url) == 0 {
		return "", NewHttpError(400, "Url is empty")
	}

	response, err := http.Get(url)
	if err != nil {
		return "", NewHttpError(500, "Error occurred")
	}

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", NewHttpError(200, "Body is empty")
	}
	return string(responseBody), nil
}
