postgres:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	sudo docker exec -it postgres12 createdb --username=root --owner=root restaurant

dropdb:
	sudo docker exec -it postgres12 dropdb restaurant

migrateup:
	sudo migrate -path db/migration -database "postgresql://root:secret@localhost:5432/restaurant?sslmode=disable" -verbose up

migratedown:
	sudo migrate -path db/migration -database "postgresql://root:secret@localhost:5432/restaurant?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup sqlc

