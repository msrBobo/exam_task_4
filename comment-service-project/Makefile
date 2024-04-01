CURRENT_DIR=$(shell pwd)
APP=template
APP_CMD_DIR=./cmd

proto-gen:
	./scripts/gen-proto.sh

swag-gen:	
	swag init -g ./api/router.go -o api/docs

DB_URL="postgres://postgres:1234@localhost:5544/commentdb?sslmode=disable"

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
	migrate create -ext sql -dir migrations/ -seq commentdb