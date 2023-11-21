#!/bin/bash
docker stop random-stuff-service
docker rm random-stuff-service
docker build . -t random-stuff-service
docker run -d --restart unless-stopped --name random-stuff-service -p 5000:5000 random-stuff-service