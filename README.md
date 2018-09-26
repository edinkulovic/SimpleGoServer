# This entire code is obsolete. Need to update :)

# Simple Go Server 
Simple test project to learn GO. 
TODO:: Add more details what we will cover, why and maybe even some diagrams  

# Installation
1. Install and setup GO     :: TODO: Add more description 
2. Setup GO Workspace (...) :: TODO: Add more description
3. Setup Docker             :: TODO: Add more description

# Docker 
## Creating an image for this docker locally
docker build -t simple-go-server .

## Starting docker container
* docker run (-d DETACHED) (-p MACHINE_PORT:DOCKER_PORT) (IMAGE_NAME)
* docker run -d -p 8000:8000 simple-go-server
* docker run -d -p 8001:8000 simple-go-server

## Reading the list of running containers
docker ps

## Stoping the docker Container by Container ID 
docker stop CONTAINER_ID

# API Testing
## Postman
 * We will use Postman (Which is chrome extension) to test and document our API calls.
 * The logic/flow behind the execution should be added into Wiki or some other place where we will store documentation. 

# MAIN API - List of Endpoints
## [/]      - Health Check
Calling this route should give you some nice response that server is a live.
## [/user]  - Retieves Simple Hardcoded Test User

# What needs to support
## Error Handling
## Good documentation
## Versioning (Backward Compatibility)
## Monitoring
## Performance and test for all method execution so that we could trace degradation
## S.O.L.I.D


### References 
https://hub.docker.com/_/redis/
