FROM golang:1.15 as session_service
ENV GO111MODULE=on
WORKDIR /go/src/session
COPY sessions /go/src/session
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM alpine
WORKDIR /app
COPY --from=session_service  /go/src/session /app
RUN chmod +x .
EXPOSE 8079/tcp
ENTRYPOINT [ "/app/sessions" ]
