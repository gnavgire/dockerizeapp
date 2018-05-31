#!/bin/bash

cd workdir
docker build -t testapp:v1 -f Dockerfile.testapp .
echo "Successfuly build docker container image"
