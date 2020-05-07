include .env
export $(shell sed 's/=.*//' .env)

run:
	go build -o ./bin/ groupbirthday && ./bin/groupbirthday