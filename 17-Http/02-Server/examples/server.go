package examples

import (
	"log"
	"net/http"
	"time"
)

func CreateServer() {
	// method 1
	go func() {
		server1 := &http.Server{
			Addr:         ":9090",
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
		}

		err := server1.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	// method 2
	log.Fatal(http.ListenAndServe(":8090", nil))
}
