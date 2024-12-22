make swag:
	@swag init --parseInternal --parseDependency

run: swag
	@go run main.go

watch: swag
	@air

scaffolder:
	@go run main.go scaffold ${model}
	make swag