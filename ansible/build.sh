#!/bin/bash

# Static variables
image_name=ansible-runner

# Get the tag for the image
read -p "Please enter $image_name version: " tag

# Call the docker build command with the params array as arguments
sudo docker buildx build -t $image_name:$tag  .
