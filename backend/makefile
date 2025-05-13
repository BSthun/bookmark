.PHONY: generate
generate: swagger mock

.PHONY: swagger
swagger:
	go run ./command/swagger
	swag init -g ./common/swagger/swagger.go -o ./generate/swagger/ --parseDependency
	sed -i '' -e 's/response.GenericResponse/Res/g' -e 's/payload_//g' ./generate/swagger/*
	go run ./command/swaggermd

.PHONY: mock
mock:
	mockery

.PHONY: build
build:
	nilaway ./...
	env GOOS=linux GOARCH=amd64 go build -o .local/backend .