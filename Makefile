clean:
	@echo "--- cleanup all build and generated files ---"
	@rm -vf app/infrastructure//proto/pb/*.pb.go

protoc: clean
	@echo "--- preparing proto output directories ---"
	@mkdir -p app/infrastructure/proto/pb

	@echo "--- Compiling all proto files ---"
	@cd ./app/infrastructure/proto && protoc -I. --go_out=plugins=grpc:./pb --govalidators_out=./pb *.proto

setup:
	@echo " --- Setup and generate configuration --- "
	@cp config/examples/server.yml.example config/server/server.yml
	@cp config/examples/database.yml.example config/database/database.yml

build: setup protoc
	@echo "--- Building binary file ---"
	@go build -o ./main cmd/server/oauth.go

build-rest: setup protoc
	@echo "--- Building binary file ---"
	@go build -o ./main cmd/rest/oauth.go

build-grpc: setup protoc
	@echo "--- Building binary file ---"
	@go build -o ./main cmd/grpc/oauth.go

run-grpc:
	@go run cmd/grpc/oauth.go

migrate:
	@go run cmd/migrate/migrate.go