FROM golang as builder
WORKDIR /go/src/github.com/farhansajid1/go-micro-project
COPY ./ /go/src/github.com/farhansajid1/go-micro-project

RUN go get ./ # build the dependencies
RUN go build -o main

RUN chmod +x main

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/github.com/farhansajid1/go-docker/ /app/
CMD ./main

EXPOSE 8080


db.oscars.find({releaseYear: {$gt:2000}}, {_id:0, releaseYear:1}).limit(10).sort({releaseYear:-1})
db.oscars.find( {title: {$regex: "^W.*"}}, {_id:0, releaseYear:1, title:1}).limit(10).sort({title:-1})