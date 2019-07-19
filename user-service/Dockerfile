FROM golang as builder

RUN mkdir -p /go/src/github.com/farhansajid1/go-micro-project/

COPY go.mod go.sum /go/src/github.com/farhansajid1/go-micro-project/
WORKDIR /go/src/github.com/farhansajid1/go-micro-project/
RUN GO111MODULE=on go mod download

COPY . .
WORKDIR /go/src/github.com/farhansajid1/go-micro-project/user-service

RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:latest
WORKDIR /app        
COPY --from=builder /go/src/github.com/farhansajid1/go-micro-project/user-service /app/
CMD ./main

EXPOSE 8080
