package main

import (
	"fmt"
	"sync"
	"time"
)

// MyServer holds the configuration for a web server
type MyServer struct {
	addr    string
	port    int
	timeout time.Duration
	mu      sync.Mutex
}

// MyServerOpts is a type func(*MyServer) used to modify default settings of MyServer
type MyServerOpts func(*MyServer)

// New returns a new *MyServer
func New(a string, opts ...MyServerOpts) *MyServer {

	m := &MyServer{
		addr: a,
		port: 8080,
		mu:   sync.Mutex{},
	}

	for _, o := range opts {
		o(m)
	}

	return m

}

// WithTimeout specifies a time.Duration timeout for MyServer
func WithTimeout(t time.Duration) MyServerOpts {
	return func(m *MyServer) {
		m.timeout = t
	}
}

func main() {

	s := New("localhost")

	// Here we use the func options pattern inside main.go
	withPort := func(m *MyServer) {
		m.port = 8081
	}

	fmt.Println("Standard server:", s)

	// Here we combine func opts using withPort and an exported function WithTimeout
	s2 := New("localhost", withPort, WithTimeout(10*time.Second))

	fmt.Println("Custom server:", s2)

}
