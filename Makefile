build:
	go build main.go
	zip lambda.zip main
	rm main
