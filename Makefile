PROTO_DIR=proto
GEN_DIR=proto/pb

gen:
	protoc \
		--go_out=$(GEN_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto

run:
	docker-compose up --build

migrate:
	DB_HOST=localhost DB_PORT=42364 go run cmd/cli/main.go migrate

seed:
	DB_HOST=localhost DB_PORT=42364 go run cmd/cli/main.go seed

clean:
	DB_HOST=localhost DB_PORT=42364 go run cmd/cli/main.go clean

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-build:
	docker build -t nexus-payment .

docker-run:
	docker run -p 8080:8080 --env-file .env nexus-payment
