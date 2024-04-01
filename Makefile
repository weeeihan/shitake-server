postgresinit:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root shitake 

dropdb:
	docker exec -it postgres15 dropdb shitake

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/shitake?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/shitake?sslmode=disable" -verbose down

getusers:
	docker exec -it postgres15 psql -d shitake -c "SELECT * FROM users"

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown getusers