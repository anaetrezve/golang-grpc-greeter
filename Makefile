generatepb:
	protoc --go_out=plugins=grpc:. greet-pb/greet.proto