
# rainbow is Database name
DB_URL=postgresql://root:asd_123@localhost:5432/rainbow?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=asd_123 -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root rainbow

dropdb:
	docker exec -it postgres dropdb rainbow

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test