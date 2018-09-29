# About

Client and server application to demonstrate the use of [UNIX datagram sockets](https://en.wikipedia.org/wiki/Unix_domain_socket) (e.g. over /tmp/gounix.sock).

Start the server first:

```bash
Usage of ./server:
  -sleep duration
        Poll interval (default 5s)
  -sock string
        Path to socket file (default "/tmp/gounix.sock")
```

Then start the client. Client will read from standard input (STDIN) and send messages after newline to the server.

```bash
Usage of ./server:
  -sleep duration
        Poll interval (default 5s)
  -sock string
        Path to socket file (default "/tmp/gounix.sock")
```

# Build
Client and server can be found in their respective folders in `cmd`. Build server with:

```bash
cd cmd/server
go build .
```

Build client with:

```bash
cd cmd/client
go build .
```