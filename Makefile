build:
	go build main.go
	zip function.zip main
	mv function.zip ~/Downloads/
