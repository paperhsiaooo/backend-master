postgres:
	docker run --name master-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.2-alpine

createdb:
	docker exec -it master-postgres createdb --username=root simple_bank

dropdb:
	docker exec -it master-postgres dropdb --username=root simple_bank

migrateup:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
