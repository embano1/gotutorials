# About

Client and server application to demonstrate the use of UNIX sockets (e.g. /tmp/go.sock).

```bash
Usage of ./client:
  -size int
        Number of bytes to transfer (default 4096)
  -socket string
        Specify Unix Domain Socket path (default "/tmp/go.sock")
```

```bash
Usage of ./server:
  -f    Force cleanup of specified socket (default: false)
  -size int
        Size of buffer for receiving data (default 4096)
  -socket string
        Specify Unix Domain Socket path (default "/tmp/go.sock")
```