build:
	go build -o main src/main.go
	zip lambda.zip main
	rm main
