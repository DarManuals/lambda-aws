build:
	go build -o main tmp/code/src/main.go
	zip lambda.zip main
	rm main

build-local:
	go build -o main src/main.go
	zip lambda.zip main
	rm main
