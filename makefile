.PHONY: deafult run build test docs clean 

APP_NAME=gopportunities

deafult: run-with-docs

run:
	@go run main.go
run-with-docs:
	@swag init
	@go run main.go
build:
	@go run build -o  $(APP_NAME) main.go
test:
	@go test  ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs