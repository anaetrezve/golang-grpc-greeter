server:
	go build greet-server/main.go

client:
	go build greet-client/main.go

generatepb:
	protoc --go_out=plugins=grpc:. greet-pb/greet.proto
