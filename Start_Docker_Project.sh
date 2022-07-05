#!/bin/bash
echo "Building Docker image with name <ascii-art-web> "
docker build -t ascii-art-web .
echo "Building and Running Docker container with name <ascii-web> on port 8080"
docker run -d -p 8080:8080 --name ascii-web ascii-art-web
echo "Lists of Docker Images"
docker images 
echo "Lists of all existing Docker containers"
docker ps -a
echo "HELP TIPS: 1)To prunes Docker's images, containers, and networks use command: docker system prune in terminal"
echo "2)Use this to delete everything: 
docker stop ascii-web
docker system prune -a --volumes"
echo "3)To start existing but not working right now container: docker start"
echo "4)To stop existing and working right now container: docker stop"