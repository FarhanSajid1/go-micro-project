#!/bin/bash

docker build -t farhansajid2/user-service ./user-service/Dockerfile .
docker build -t farhansajid2/email-service ./email-service/Dockerfile .

docker push farhansajid2/email-service
docker push farhansajid2/user-service