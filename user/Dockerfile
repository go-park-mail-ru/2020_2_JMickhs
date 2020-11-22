FROM golang:1.15 as user_service
ENV GO111MODULE=on
WORKDIR /go/src/user
COPY user /go/src/user
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64   go build .

FROM alpine
WORKDIR /app
COPY --from=user_service  /go/src/user /app
RUN chmod +x  .
EXPOSE 8080/tcp
ENTRYPOINT [ "/app/JMickhs_user" ]
