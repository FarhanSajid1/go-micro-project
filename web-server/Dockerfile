FROM golang as builder

RUN mkdir -p /go/src/github.com/farhansajid1/go-micro-project/

RUN curl https://glide.sh/get | sh
COPY glide.lock glide.yaml /go/src/github.com/farhansajid1/go-micro-project/
WORKDIR /go/src/github.com/farhansajid1/go-micro-project/
RUN glide install

COPY . .
WORKDIR /go/src/github.com/farhansajid1/go-micro-project/web-server

RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:latest
WORKDIR /app        
COPY --from=builder /go/src/github.com/farhansajid1/go-micro-project/web-server /app/
CMD ./main

EXPOSE 8080
