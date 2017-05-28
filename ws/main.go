package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

/* Resources

https://golang.org/doc/
https://golang.org/doc/code.html
https://play.golang.org/
https://tour.golang.org/welcome/1

*/

const port = ":40000"

func debug(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s - %s", r.Method, r.URL.Path, r.UserAgent())
	fmt.Fprintf(w, "Hi there, thx for calling %s!", r.URL.Path[1:])
}

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	index := http.FileServer(http.Dir("www"))

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Kill)

	go func() {
		sig := <-sigs
		log.Printf("Got %v...terminating.", sig)
		done <- true
	}()

	go func() {
		http.Handle("/", index)
		http.HandleFunc("/debug", debug)

		log.Printf("Starting web server on port %s...", port)
		log.Fatal(http.ListenAndServe(port, nil))
	}()

	<-done
}
