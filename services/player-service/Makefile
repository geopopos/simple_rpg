# Go parameters
    BINARY=player-service
    ORG=docker.io/sleyva97
    ACTIVE_VERSION=1.9
    ROLLBACK_VERSION=1.8
    DOCKER_IMAGE=$(ORG)/$(BINARY)

    unit-test:
	go test ./...

    docker-integration:
	docker-compose build --build-arg VERSION=$(ACTIVE_VERSION)
	docker-compose push

    docker-push:
	docker pull $(DOCKER_IMAGE):$(shell git rev-parse HEAD)
	docker tag $(DOCKER_IMAGE):$(shell git rev-parse HEAD) $(DOCKER_IMAGE):$(ACTIVE_VERSION)
	docker push $(DOCKER_IMAGE):$(ACTIVE_VERSION)

