FROM golang:alpine as build
WORKDIR /go/src
COPY main.go .
RUN go build -o /go/bin/memhogger .

FROM alpine
LABEL maintainer="Michael Gasch <embano1@live.com>"
COPY --from=build /go/bin/memhogger /bin/memhogger
ENTRYPOINT ["memhogger"]