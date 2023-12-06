generate:
	protoc -I proto proto/sso/sso.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative

migrate:
	go run ./cmd/migrator/ --storage-path=./storage/sso.db --migration-path=./migrations