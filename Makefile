
postgres:
	docker run --name skatespot -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

createdb:
	docker exec -it skatespot createdb --username=root --owner=root skatespotdb

dropdb:
	docker exec -it skatespot dropdb skatespotdb

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/skatespotdb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/skatespotdb?sslmode=disable" -verbose down

.PHONY: createdb dropdb postgres migrateup migratedown



