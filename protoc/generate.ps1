protoc -I . --go_out=plugins=grpc,paths=source_relative:. keepaconfig/*.proto