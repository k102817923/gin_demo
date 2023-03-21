FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/gin_demo
COPY . $GOPATH/src/gin_demo
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./gin_demo"]