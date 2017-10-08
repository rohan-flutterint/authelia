#!/bin/bash

# Parameters:
# TAG - The name of the tag to use for publishing in Dockerhub
#
function deploy_on_dockerhub {
  echo "======================================="
  echo "Authelia will be deployed on Dockerhub."
  echo "======================================="

  TAG=$1
  IMAGE_NAME=clems4ever/authelia
  IMAGE_WITH_TAG=$IMAGE_NAME:$TAG

  echo "Logging in to Dockerhub as $DOCKER_USERNAME."
  docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD
  if [ "$?" -ne "0" ];
  then
    echo "Logging in to Dockerhub failed.";
    exit 1
  fi

  echo "Docker image $IMAGE_WITH_TAG will be deployed on Dockerhub."
  docker build -t $IMAGE_NAME .
  docker tag $IMAGE_NAME $IMAGE_WITH_TAG;
  docker push $IMAGE_WITH_TAG;
  echo "Docker image deployed successfully."
}


if [ "$TRAVIS_BRANCH" == "master" ]; then
  deploy_on_dockerhub latest
elif [ ! -z "$TRAVIS_TAG" ]; then
  deploy_on_dockerhub $TRAVIS_TAG
else
  echo "Docker image will not be deployed on Dockerhub."
fi

