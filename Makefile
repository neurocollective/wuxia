db/dump:
	@pg_dump --schema-only -f dumps/schema.sql -d postgres -h localhost -p 5432 -U postgres
db/create:
	@docker run --name wuxia-local -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres
	@sleep 3
	@psql -f db/create_tables.sql "postgresql://postgres:postgres@localhost:5432/postgres" 
	@psql -f db/initial_seed.sql "postgresql://postgres:postgres@localhost:5432/postgres"
psql:
	@psql "postgresql://postgres:postgres@localhost:5432/postgres"
test:
	@go test ./structs -v
fmt:
	@go fmt ./main.go
	@go fmt ./generation/index.go
	@go fmt ./structs/index.go
