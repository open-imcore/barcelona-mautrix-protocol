ipc-mac-deps:
	brew install swift-protobuf protoc-gen-go grpc-swift

swift_opts=--grpc-swift_opt=Visibility=Public --grpc-swift_out=Sources/BarcelonaMautrixIPCProtobuf --swift_opt=Visibility=Public --swift_out=Sources/BarcelonaMautrixIPCProtobuf
go_opts=--go_out='$(PWD)'

ipc-v1-%:
	protoc $($*_opts) ipc/v1/v1.proto

ipc-v1: ipc-v1-go ipc-v1-swift