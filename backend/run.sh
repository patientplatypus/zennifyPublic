#!/bin/bash

# docker build --no-cache -t go_server .
# # docker run -d -p 8081:8081 goplay 
# docker run -it -p 8081:8081 -p 5432:5432 go_server

docker-compose up -d -e CASSANDRA_START_RPC=true --build