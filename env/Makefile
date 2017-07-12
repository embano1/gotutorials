REPO = embano1/env:1.0

default: image

build:
	GOOS=linux go build .

image: build
	docker build -t ${REPO} .

push: image 
	docker push ${REPO}

clean:
	rm env

all: clean push