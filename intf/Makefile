build:
	GOOS=linux go build -o app .

image: build
	docker build -t embano1/interface .
	rm app

clean:
	rm app

all: image
