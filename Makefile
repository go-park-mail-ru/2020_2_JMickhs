.PHONY:
server-redis: 
	redis-server --port 6402 & go run main/main.go --server

sessions-server:
	redis-server --port 6402 & go run main.go

gen: 
	swagger generate spec -o ./api/swagger/swagger.yaml --scan-models

mocks:
	go generate -v ./...

fillbd:
	go run main.go --fill

server:
	go run main.go -server

upload:
     sudo docker build -t kostikan/main_service:latest -f ./main/Dockerfile . &&
     sudo docker build -t kostikan/session_service:latest -f ./sessions/Dockerfile . &&
     sudo docker build -t kostikan/user_service:latest -f ./user/Dockerfile . &&
     sudo docker push kostikan/main_service:latest &&
     sudo docker push kostikan/session_service:latest &&
     sudo docker push kostikan/user_service:latest &&
     sudo APP_VERSION=latest docker-compose up



tests:
	go test -coverprofile=coverage1.out -coverpkg=./... -cover ./... && cat coverage1.out | grep -v  easyjson | grep -v mocks | grep -v server > cover.out &&go tool cover -func=cover.out
