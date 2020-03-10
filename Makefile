build:
	go build -o main tmp/code/main.go
	zip lambda.zip main
	rm main
