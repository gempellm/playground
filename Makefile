.PHONY: gen-go
gen-go:
	protoc -I=internal/protobuf/protos --go_out=internal/protobuf/go/pb internal/protobuf/protos/my/my.proto --go_opt=paths=source_relative

.PHONY: run
run:
	go run cmd/playground/main.go

.PHONY: gen
gen:
	buf generate internal/protobuf/protos