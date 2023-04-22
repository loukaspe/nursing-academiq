generate-mock:
ifndef FILE
	@echo "FILE parameter is missing"
	@exit 1
endif
	make build-dev
	@docker run --volume "$(PWD)":/app --workdir /app \
	dev-build /bin/bash -c "mockgen -source=${FILE} -destination=mocks/mock_${FILE}"

tests-unit:
	make build-dev
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			dev-build godotenv -f .env.test go test -short -cover -count=1 ./...

tests-benchmark:
	make build-dev
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			dev-build godotenv -f .env.test go test ./... -bench=.

tests-file:
	make build-dev
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			dev-build godotenv -f .env.test go test -count=1 -v ${FILE}

tests-package:
	make build-dev
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			dev-build godotenv -f .env.test go test -count=1  ${PACKAGE}


tests-all:
	make build-dev
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			dev-build godotenv -f .env.test go test ./... -cover -count=1

tests-all-with-coverage:
	make build-dev
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			dev-build godotenv -f .env.test go test -count=1 -v -coverpkg=./... -coverprofile=profile.cov ./... ; go tool cover -func profile.cov

run-linter:
	make build-dev
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			dev-build golangci-lint run ./...

generate-swagger-files:
	make build-dev
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			dev-build swagger generate spec -o ./docs/swagger.json --scan-models

start-app:
	docker-compose down
	docker-compose up

kill-app:
	docker-compose down --remove-orphans --volumes

reset-app:
	docker-compose down --remove-orphans --volumes
	docker-compose up --build

rebuild-app:
	docker-compose stop app
	docker-compose build app
	docker-compose up -d app

build-dev:
	@docker build \
			--tag dev-build \
			-f Dockerfile.utilities .