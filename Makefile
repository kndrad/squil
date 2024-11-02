.PHONY: build
build:
	go build -o bin/squil main.go

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: docker-up
docker-up:
	docker-compose up --build -d

.PHONY: docker-down
docker-down:
	docker-compose down

config_path=config.yml
.PHONY: pingdb
pingdb:
	go run main.go --config=$(config_path) pingdb


.PHONY: review
review:
	./scripts/format.sh
	./scripts/check.sh
	go test ./... -count=1 -failfast -coverprofile=coverage.out

.PHONY:
dogs:
	go run main.go --config=$(config_path) dogs

.PHONY: docker-rmv
docker-rmv:
	docker-compose down -v

.PHONY: migrate-up
migrate-up:
	migrate -database "${POSTGRES_URL}" -path db/migrations up

.PHONY: migrate-down
migrate-down:
	migrate -database "${POSTGRES_URL}" -path db/migrations down

.PHONY: migrate-version
migrate-version:
	migrate -database "${POSTGRES_URL}" -path db/migrations version

.PHONY: add-dog
add-dog:
	go run main.go dogs create --name="" --breed=""
