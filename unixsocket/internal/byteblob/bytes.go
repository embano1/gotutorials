package byteblob

import (
	"context"
	"io"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

type result struct {
	received int
	duration float64
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// New returns a byte slice of the given size and an error
func New(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Receive starts reading from the specified connection. 
// It takes a context (for cancellation) and waitgroup (clean shutdown). Buffer size used for reading is configurable.
// Upon returning it closes conn.
func Receive(ctx context.Context, wg *sync.WaitGroup, conn net.Conn, bufSize int) {
	defer func() {
		conn.Close()
		wg.Done()
	}()

	b := make([]byte, bufSize)
	var read int
	d := time.Now()
	for {
		select {	
		case <-ctx.Done():
			log.Println("got cancelled")
			return
		default:
			n, err := conn.Read(b)
			read += n
			switch {
			case err == io.EOF:
				r := result{
					received: read,
					duration: time.Now().Sub(d).Seconds(),
				}
				log.Printf("bytes received: %d - took: %.3fs - throughput: (%.2f MB/s)\n", r.received, r.duration, (float64(r.received)/r.duration)/1024/1024)
				return
			case err != nil:
				log.Printf("error reading from connection: %v\n", err)
				return
			}
		}
	}
}
