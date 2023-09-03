package management

import (
	handlers "context-examples/handler"
	"fmt"
	"net/http"
	"time"
)

func Run() {
	mux := http.NewServeMux()
	mux.Handle("/users/", &handlers.UsersHandler{})
	mux.HandleFunc("/test/", TestHandler)
	server := &http.Server{
		Addr:         ":8011",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Handler:      mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-ctx.Done():
		println("request canceled")
		return
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "Response")
		println("Processing request")
	}
}
