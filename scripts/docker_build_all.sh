#!/bin/sh

# Build all dockers

cd ../deyes
docker build -t deyes .
cd ../dheart
docker build -t dheart .
cd ../sisu
docker build -t sisu .