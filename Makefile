make swag:
	@swag init --parseInternal --parseDependency

run: swag
	@go run main.go

watch: swag
	@air