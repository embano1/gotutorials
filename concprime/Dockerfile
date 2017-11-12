FROM golang:1.9.2-alpine3.6 AS build
COPY . /go/src/github.com/embano1/gotutorials/concprime
WORKDIR /go/src/github.com/embano1/gotutorials/concprime/cmd
RUN go build -o concprime

FROM alpine:3.6
ARG COMMIT=unknown
LABEL COMMIT=${COMMIT}
COPY --from=build /go/src/github.com/embano1/gotutorials/concprime/cmd/concprime /bin
CMD ["time", "concprime", "-n", "10000"]