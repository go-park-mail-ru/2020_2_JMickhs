.PHONY:
server-redis: 
	redis-server --port 6402 & go run main/main.go --server

sessions-server:
	redis-server --port 6402 & go run main.go

gen: 
	 GO111MODULE=off  swagger generate spec -o ./main/api/swagger/swagger.yaml --scan-models

mocks:
	go generate -v ./...

user_service_mocks:
	go generate mockgen -source user.pb.go -destination user_service_mock.go -package userService

sessions_service_mocks:
	go generate mockgen -source session.pb.go -destination session_service_mock.go -package sessionService

fillbd:
	go run main.go --fill

server:
	go run main.go -server

upload:
	sudo docker build -t kostikan/main_service:latest -f ./main/Dockerfile .
	sudo docker build -t kostikan/session_service:latest -f ./sessions/Dockerfile .
	sudo docker build -t kostikan/user_service:latest -f ./user/Dockerfile .
	sudo docker push kostikan/main_service:latest
	sudo docker push kostikan/session_service:latest
	sudo docker push kostikan/user_service:latest
	sudo APP_VERSION=latest docker-compose up

dockerClean:
	sudo docker rm -vf $(sudo docker ps -a -q)

dockerRun:
	sudo APP_VERSION=latest docker-compose up

tests:
	cd main && go test -coverprofile=coverage1.out -coverpkg=./... -cover ./... && cat coverage1.out | grep -v  easyjson | grep -v mocks | grep -v server > cover.out &&go tool cover -func=cover.out
	cd ..
	cd user && go test -coverprofile=./coverage1.out -coverpkg=./... -cover ./... && cat coverage1.out | grep -v  easyjson | grep -v mocks | grep -v main > cover.out &&go tool cover -func=cover.out
	cd ..
	cd sessions && go test -coverprofile=./coverage1.out -coverpkg=./... -cover ./... && cat coverage1.out | grep -v  easyjson | grep -v mocks | grep -v main > cover.out &&go tool cover -func=cover.out
	cd ..


