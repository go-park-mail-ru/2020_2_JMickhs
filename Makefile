.PHONY:
server-redis: 
	redis-server --port 6402 & go run main.go 

gen: 
	swagger generate spec -o ./swagger.yaml --scan-models

mocks:
	go generate -v ./...

server:
	go run main.go 