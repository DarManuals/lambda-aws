build:
	go build -o main tmp/code/src/*.go
	zip lambda.zip main
	rm main

build-local:
	go build -o main src/*.go
	zip lambda.zip main
	rm main
