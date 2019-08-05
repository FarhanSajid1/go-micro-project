FROM golang as builder

# for golang we need to create a src directory as such..
RUN mkdir -p /go/src/github.com/farhansajid1/go-micro-project/

RUN curl https://glide.sh/get | sh
COPY glide.lock glide.yaml /go/src/github.com/farhansajid1/go-micro-project/
WORKDIR /go/src/github.com/farhansajid1/go-micro-project/
RUN glide install

COPY . .
WORKDIR /go/src/github.com/farhansajid1/go-micro-project/user-service

# command to build the image.
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:latest
WORKDIR /app        
COPY --from=builder /go/src/github.com/farhansajid1/go-micro-project/user-service /app/
CMD ./main

EXPOSE 8080
