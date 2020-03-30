account-go:
	protoc --go_out=plugins=grpc:./model ./*.proto