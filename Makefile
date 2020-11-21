.PHONY:
server-redis: 
	redis-server --port 6402 & go run main/main.go --server

gen: 
	swagger generate spec -o ./api/swagger/swagger.yaml --scan-models

mocks:
	go generate -v ./...

fillbd:
	go run main.go --fill

server:
	go run main.go -server

tests:
	go test -coverprofile=coverage1.out -coverpkg=./... -cover ./... && cat coverage1.out | grep -v  easyjson | grep -v mocks | grep -v server > cover.out &&go tool cover -func=cover.out
