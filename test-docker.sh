#!/bin/bash

echo "Testing Docker image locally..."

# Build the image
docker build -t healthrecord-repository-test .

# Run the container and capture output
echo "Running container..."
docker run --rm -p 5655:5655 healthrecord-repository-test &

# Get the container ID
CONTAINER_ID=$(docker ps -q --filter ancestor=healthrecord-repository-test)

if [ -n "$CONTAINER_ID" ]; then
    echo "Container started with ID: $CONTAINER_ID"
    echo "Waiting 5 seconds..."
    sleep 5
    
    echo "Checking if app is responding..."
    curl -s http://localhost:5655/sbom || echo "App not responding on port 5655"
    
    echo "Container logs:"
    docker logs $CONTAINER_ID
    
    echo "Stopping container..."
    docker stop $CONTAINER_ID
else
    echo "Container failed to start. Checking logs from last run..."
    echo "Trying to run container with verbose output..."
    docker run --rm healthrecord-repository-test 2>&1 || echo "Container exited with error"
    
    echo ""
    echo "Testing if the binary exists and is executable..."
    docker run --rm --entrypoint="" healthrecord-repository-test ls -la /tmp/
    
    echo ""
    echo "Testing binary directly..."
    docker run --rm --entrypoint="" healthrecord-repository-test /tmp/main --help 2>&1 || echo "Binary execution failed"
fi