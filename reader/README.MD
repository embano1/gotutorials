# Simple io.Reader implemention
A simple implementation of io.Reader based on pkg strings `func (*Reader) Read` implementation:

```go
func (r *Reader) Read(b []byte) (n int, err error) {
  	if r.i >= int64(len(r.s)) {
  		return 0, io.EOF
  	}
  	r.prevRune = -1
  	n = copy(b, r.s[r.i:])
  	r.i += int64(n)
  	return
  }
```

Source: https://golang.org/src/strings/reader.go#L37