parson:
	@echo "Getting Parson Ready"
	@go build && ./parson ${dir}

run:
	@go build && ./parson

testall: 
	@go test ./...
