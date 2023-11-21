#!/bin/bash
docker stop random-stuff-service
docker rm random-stuff-service
docker build . -t random-stuff-service
source ~/.bashrc
docker run -d --restart unless-stopped -e AUTH0_DOMAIN -e AUTH0_AUDIENCE -e AUTH0_CLIENT_ID -e AUTH0_CLIENT_SECRET -e GIN_MODE -e GIN_PORT -e DSN --name random-stuff-service -p 5000:5000 random-stuff-service