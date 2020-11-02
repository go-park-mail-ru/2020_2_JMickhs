.PHONY:
server-redis: 
	redis-server --port 6402 & go run main.go --server

gen: 
	swagger generate spec -o ./api/swagger/swagger.yaml --scan-models

mocks:
	go generate -v ./...

fillbd:
	go run main.go --fill

server:
	go run main.go -server