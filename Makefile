proto-gen:
	protoc --go_out=.  --go-grpc_out=. ./proto/cosmos/lms/v1beta1/*.proto