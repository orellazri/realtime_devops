.PHONY: chart http redis kafka emqx

chart:
	go run main.go

test: http redis kafka emqx

http:
	go test ./http -count=1

redis:
	go test ./redis -count=1

kafka:
	go test ./kafka -count=1

emqx:
	go test ./emqx -count=1
