#!/bin/bash

docker build -t "farhansajid2/user-service:$CIRCLE_SHA1" -t farhansajid2/user-service:lastest -f ./user-service/Dockerfile .
docker build -t "farhansajid2/email-service:$CIRCLE_SHA1" -t farhansajid2/email-service:latest -f ./email-service/Dockerfile .
docker build -t "farhansajid2/web-server:$CIRCLE_SHA1" -t farhansajid2/web-server:latest -f ./web-server/Dockerfile .


echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_ID" --password-stdin

docker push farhansajid2/email-service:$CIRCLE_SHA1
docker push farhansajid2/email-service:latest

docker push farhansajid2/user-service:latest
docker push farhansajid2/user-service:$CIRCLE_SHA1

docker push farhansajid2/web-server:latest
docker push farhansajid2/web-server:$CIRCLE_SHA1