.PHONY : zip
zip :
	GOOS=linux go build -o main.go
	zip function.zip main

.PHONY : build
build :
	go build -o main .

.PHONY : env
env :
	cp .env.example .env