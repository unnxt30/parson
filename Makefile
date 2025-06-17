.PHONY: parson
parson:
	@echo "Getting Parson Ready"
	@go build && ./parson ${dir}