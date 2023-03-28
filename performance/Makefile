.PHONY: chart http redis kafka

chart:
	go run main.go

test: http redis kafka

http:
	go test ./http -count=1

redis:
	go test ./redis -count=1

kafka:
	go test ./kafka -count=1
