.PHONY: parson
parson:
	@echo "Getting Parson Ready"
	@go build && ./parson ${dir}

.PHONY: run
run:
	@go build && ./parson