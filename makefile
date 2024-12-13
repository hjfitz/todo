build:
	go build .

clean:
	rm todo

install:
	make build
	mv todo ~/.bin/
