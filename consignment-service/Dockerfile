FROM golang as builder

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep 

RUN mkdir -p /go/src/github.com/farhansajid1/go-micro-project/

COPY Gopkg.lock Gopkg.toml /go/src/github.com/farhansajid1/go-micro-project/
WORKDIR /go/src/github.com/farhansajid1/go-micro-project/
RUN dep ensure --vendor-only

COPY . .
WORKDIR /go/src/github.com/farhansajid1/go-micro-project/consignment-service

RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:latest
WORKDIR /app        
COPY --from=builder /go/src/github.com/farhansajid1/go-micro-project/consignment-service /app/
CMD ./main

EXPOSE 8080
