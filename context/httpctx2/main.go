package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"strconv"

	"github.com/google/uuid"
)

const endpoint = "/uuid"

func main() {
	// CLI Flags
	port := flag.Int("p", 8080, "Port the server listens on")
	servtimeout := flag.Duration("st", time.Duration(time.Second*10), "Server timeout (how long to delay the server answer)")
	help := flag.Bool("h", false, "Display this help")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(1)
	}

	// Set up root context and prepare to catch OS signals
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	defer func() {
		signal.Stop(c)
		cancel()
	}()

	// Run httpd
	go runserver(ctx, servtimeout, endpoint, port)

	// Catch OS sigs
	sig := <-c
	log.Printf("Got %v\n", sig)
	log.Println("Attempting graceful shutdown (5s)")

	// Cancel root context
	cancel()

	// Give us some time, displaying with ticker
	t := time.NewTicker(time.Second)
	go func() {
		for range t.C {
			log.Println("Waiting...")
		}
	}()

	<-time.After(time.Second * 5)
	t.Stop()
}

func slowUUIDGenerator(ctx context.Context, uuidCh chan<- uuid.UUID, errCh chan error, t *time.Duration) {

	log.Printf("[UUID Generator] Sleeping for %v\n", *t)

	select {
	case <-time.After(*t):
	case <-ctx.Done():
		log.Println("[UUID Generator] We were asked to cancel")
		return
	}

	u, err := uuid.NewRandom()
	if err != nil {
		errCh <- errors.New("[UUID Generator] Could not generate uuid: " + err.Error())
		return
	}
	log.Printf("[UUID Generator] Generated this UUID %v\n", u)
	uuidCh <- u

}

func workhandler(ctx context.Context, stimeout *time.Duration) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		tm := time.Now()
		uuidCh := make(chan uuid.UUID)
		errCh := make(chan error)

		cn, ok := w.(http.CloseNotifier)
		if !ok {
			log.Fatal("Responsewriter does not implement CloseNotify()")
		}

		go slowUUIDGenerator(ctx, uuidCh, errCh, stimeout)

		select {
		case <-cn.CloseNotify():
			log.Println("Client disconnected")
			cancel()
		case <-ctx.Done():
			w.WriteHeader(http.StatusServiceUnavailable)
		case u := <-uuidCh:
			fmt.Fprintf(w, "The time when this operation started: %v\n", tm)
			fmt.Fprintf(w, "New UUID generated: %v\n", u)
			fmt.Fprintf(w, "Operation took: %v\n", time.Since(tm))
		case e := <-errCh:
			log.Println(e)
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	}

	return http.HandlerFunc(fn)
}

func runserver(ctx context.Context, stimeout *time.Duration, endpoint string, port *int) {
	// TODO: implement http.Shutdown(ctx) -> see example in httpctx folder
	mux := http.NewServeMux()
	wh := workhandler(ctx, stimeout)
	mux.Handle(endpoint, wh)
	server := &http.Server{
		Addr:    string(":" + strconv.Itoa(*port)),
		Handler: mux,
	}

	log.Printf("Listening on :%d...\n", *port)
	log.Printf("Endpoint to generate UUIDs is %v\n", endpoint)
	log.Fatal(server.ListenAndServe())
}
