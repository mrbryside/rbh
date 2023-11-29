integration-test-up:
	DOCKER_BUILDKIT=0 docker-compose -f docker-compose-test.yml up --build --abort-on-container-exit --exit-code-from it_tests

integration-test-down:
	DOCKER_BUILDKIT=0 docker-compose -f docker-compose-test.yml down

unit-test:
	go test ./... -v -tags=unit  

build:
	go build -o "app" ./cmd/monolith

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

run:
	go run ./cmd/monolith/main.go