#!/bin/bash

port=${1:-3000}
echo Starting battlesnake server: http://localhost:$port/
docker run -it --rm -p ${port}:3000 sendwithus/battlesnake-server
