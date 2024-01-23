up:
	docker-compose up

down:
	docker-compose down

tidy:
	cd src/public && go mod tidy

swagger-public:
	cd src/public && swag init --parseDependency --parseDepth=3

test:
	cd src/public && go test ./...

build-public:
	docker build \
		-t namnam206/minio-adapter-public:latest \
		--build-arg="BUILD_MODULE=public" \
		-f ./docker/Dockerfile \
		.

push-public:
	docker push namnam206/minio-adapter-public:latest
