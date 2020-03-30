account-go:
	protoc --go_out=plugins=grpc:./model ./*.proto
test:
	go test -v -race -covermode=atomic -coverprofile=coverage.coverprofile ./...