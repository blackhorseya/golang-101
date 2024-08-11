.PHONY: gen-pb
gen-pb: ## Generate protobuf files
	@echo Starting generate pb
	@protoc --proto_path=./ \
     		--go_out=paths=source_relative:./ \
     		--go-grpc_out=paths=source_relative,require_unimplemented_servers=false:./ \
     		--go-grpc-mock_out=paths=source_relative,require_unimplemented_servers=false:./ \
     		./*.proto
	@echo Successfully generated proto

	@echo Starting inject tags
	@protoc-go-inject-tag -input="./*.pb.go"
	@echo Successfully injected tags
