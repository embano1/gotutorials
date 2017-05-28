APPNAME = ws

build:
	GOOS=linux go build -o $(APPNAME) .

image: build
	docker build -t embano1/$(APPNAME) .
	rm $(APPNAME)

clean:
	rm $(APPNAME)

all: image
