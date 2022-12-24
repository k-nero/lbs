package lbs

//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//go:generate protoc --go_out=paths=source_relative:./  address.proto
