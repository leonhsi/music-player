postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123 -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root music-player

dropdb:
	docker exec -it postgres dropdb music-player

migrateup:
	migrate -path db/migration -database "postgresql://root:123@localhost:5432/music-player?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:123@localhost:5432/music-player?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
