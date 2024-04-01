CURRENT_DIR=$(shell pwd)
APP=template
APP_CMD_DIR=./cmd


DB_URL="postgresdb://postgres:1234@postgres:5433/postgres?sslmode=disable"

build:
	CGO_ENABLED=0 GOOS=darwin go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

r:
	go run cmd/main.go

migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migrate-force:
	migrate -path migrations -database "$(DB_URL)" -verbose force

migrate-file:
	migrate create -ext sql -dir migrations/ -seq db