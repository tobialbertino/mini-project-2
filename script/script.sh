# using mockery 
go install github.com/vektra/mockery/v2@v2.20.0
# then 
mockery --all --keeptree --case=underscore --with-expecter=true

# Mock with mockgen
mockgen --source=repository.go --build_flags=--mod=mod --destination=mock/interface.go --package=mock

# Docker Container 
# Create docker image from Dockerfile 
docker build -t mini-project:v1 .

# List Docker Image
docker image ls

# Rename the image to match the Docker Hub format
docker image tag mini-project:v1 tobialbertino/mini-project:v1

# Login ke Docker Hub
# export PASSWORD_DOCKER_HUB=***
docker login -u tobialbertino -p $1

# Push to DockerHub 
docker push tobialbertino/mini-project:v1


