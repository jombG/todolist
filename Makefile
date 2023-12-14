generate: ## Code generation
	# Generate from .go code
	@make gen_api

.PHONY: fmt
fmt: ## Format source using gofmt
	# Apply go fmt
	@gofmt -l -s -w ./


imports: ## fix go imports
	@goimports -local todolist -w -l ./
	@gci write --skip-generated -s standard -s default -s"prefix(todolist)" .

lint: ## Linter for golang
	@docker run --rm -it -v $(PWD):/app -w /app golangci/golangci-lint:v1.50.0-alpine golangci-lint run ./...

lint-fix: ## Linter fixes for golang
	@docker run --rm -it -v $(PWD):/app -w /app golangci/golangci-lint:v1.50.0-alpine golangci-lint run --fix ./...

test:  ## Run all unit test for GitLab CI
	@godotenv -f ./app.env go test ./... -cover -race -v

gen_orm: ## Generate repository
	@go run entgo.io/ent/cmd/ent generate --target internal/storage/ent ./internal/storage/schema

gen_api: ## Generate api
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.15.0 -package generated -generate types api/swagger.yaml > ./internal/server/generated/types.gen.go
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.15.0 -package generated -generate strict-server -generate chi-server api/swagger.yaml > ./internal/server/generated/server.gen.go
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.15.0 -package generated -generate strict-server -generate spec api/swagger.yaml > ./internal/server/generated/spec.gen.go

run: ## Start project in Docker
	docker-compose up -d