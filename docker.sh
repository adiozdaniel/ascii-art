#!/bin/bash

# Build the docker image
docker image build -f Dockerfile -t ascii-web:v5.0.1 .

# Run the docker container
docker container run -p 8080:8080 --detach --name ascii-web-container ascii-web:v5.0.1
