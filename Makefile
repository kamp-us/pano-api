gen:
	@echo "Generating proto files..."
	protoc --twirp_out=. --go_out=. rpc/pano-api/service.proto
upgrade:
	@echo "Upgrading dependencies..."
	go get -u