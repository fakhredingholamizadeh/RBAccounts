
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

sqlc:    # make sqlc
	sqlc generate

test:    # make test .
	go test -v -cover ./...

proto: 
#	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

server:
	go run main.go
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test proto server


# func New(db DBTX) *Queries {
#	return &Queries{db: db}
#}

#func NewStore(db *sql.DB) *Store {
#	return &Store{db: db, Queries: New(db) }
#}


