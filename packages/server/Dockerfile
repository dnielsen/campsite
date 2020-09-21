FROM golang:1.14.3

RUN apt-get -qq update && apt-get -yqq install upx

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/main.go

EXPOSE 4444

CMD ["./main"]