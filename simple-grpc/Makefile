.PHONY: gen-pb
gen-pb: ## Generate protobuf files
	buf generate
	protoc-go-inject-tag -input="./pb/*.pb.go"
