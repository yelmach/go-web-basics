#!/bin/bash

# Set variables
IMAGE_NAME="myapp"
CONTAINER_NAME="mycontainer"
PORT=5000

# Build the Docker image
echo "[.] Building Docker image..."
docker build -t $IMAGE_NAME .

# Check if the build was successful
if [ $? -ne 0 ]; then
    echo "[.] Docker build failed. Exiting."
    exit 1
fi

# Stop all containers using the same port
echo "[.] Stopping all containers using port $PORT..."
containers=$(docker ps -q --filter "publish=$PORT")
if [ ! -z "$containers" ]; then
    docker stop $containers
    echo "[.] Stopped containers: $containers"
fi

# Check if a container with the same name already exists
if [ $(docker ps -aq -f name=^/${CONTAINER_NAME}$) ]; then
    echo "[.] Container already exists. Removing it..."
    docker rm -f $CONTAINER_NAME
fi

# Run the Docker container
echo "[.] Running Docker container..."
docker run -d -p $PORT:$PORT --name $CONTAINER_NAME $IMAGE_NAME

# Check if the container is running
if [ $(docker ps -q -f name=^/${CONTAINER_NAME}$) ]; then
    echo "[.] The server is running on http://localhost:$PORT"
else
    echo "Failed to start the container."
    exit 1
fi

echo " "
echo "########### YOU ARE NOW USING BASH THAT INSIDE THE CONTAINER ###########"
docker exec -it $CONTAINER_NAME /bin/bash