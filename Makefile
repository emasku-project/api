docs:
	swag init -g cmd/api/main.go

di:
	wire gen api/internal/injector
.PHONY: docs di