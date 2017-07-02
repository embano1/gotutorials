package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	fmt.Println("Starting server")
	fmt.Println("Call http://localhost:8080/hello")

	srv := newServer()
	go srv.ListenAndServe()

	// Root ctx
	ctx := context.Background()

	// Uncomment to see the effect of a canceled context on http.Shutdown()
	/*ctx, cancel := context.WithTimeout(ctx, time.Second*7)
	defer func() {
		cancel()
		fmt.Println("Exiting")
	}()*/

	fmt.Println("Giving you some time to make your http client call")
	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	fmt.Println("Attempting graceful shutdown of http server")
	err := srv.Shutdown(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	} else {
		fmt.Println("Done")
	}
}

func newServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler sleeping")
		time.Sleep(10 * time.Second)
		w.Write([]byte("Hey Gopher!"))
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return srv
}
