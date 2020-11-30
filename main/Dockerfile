FROM golang:1.15 as main_service
ENV GO111MODULE=on
WORKDIR /go/src/main
COPY main /go/src/main
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build .

FROM alpine
WORKDIR /app
COPY --from=main_service  /go/src/main /app
RUN chmod +x  .
EXPOSE 8080/tcp
ENTRYPOINT [ "/app/main", "--server" ]

