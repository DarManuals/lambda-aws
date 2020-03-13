build:
	go build -o main tmp/code/src/main.go
	zip lambda.zip main
	rm main
