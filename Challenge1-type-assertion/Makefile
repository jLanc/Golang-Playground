build:
	go build

build-docker:
	docker build -t challenge1 .
	docker run --rm -v .:/app challenge1
	docker image rm challenge1