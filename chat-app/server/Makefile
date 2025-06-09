pullDB: 
	docker pull postgres:12-alpine

postgresInit:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine


createDB:
	docker exec -it postgres12 createdb --username=root --owner=root chat-app


dropDB:
	docker exec -it postgres12 dropdb chat-app


enterDB:
	docker exec -it postgres12 psql
# \l to list, \c to enter a path or dir, \d to inspect a table, exit to leave the shell

runApp:
	clear && cd cmd/api && go run main.go

createTable:
	migrate create -ext sql -dir db/migration add_users_table

migrateRecordUp:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/chat-app?sslmode=disable" -verbose up

migrateRecordDown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/chat-app?sslmode=disable" -verbose down


.PHONY: pullDB postgresInit enterDB createDB dropDB runApp createTable migrateRecordUp migrateRecordDown