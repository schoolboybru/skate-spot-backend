
postgres:
	docker run --name locationservice -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it locationservice createdb --username=root --owner=root locationservicedb

dropdb:
	docker exec -it locationservice dropdb locationservicedb

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/locationservicedb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/locationservicedb?sslmode=disable" -verbose down

.PHONY: createdb dropdb postgres migrateup migratedown



