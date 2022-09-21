
postgres:
	sudo docker run --name rps -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres
	
createdb:
	sudo docker exec -it rps createdb --username=root --owner=root rps_db

dropdb:
	sudo docker exec -it rps dropdb rps_db

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/rps_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/rps_db?sslmode=disable" -verbose down
	
sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server 