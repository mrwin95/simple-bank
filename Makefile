postgres:
	docker run --name postgresql17 -p 5432:5432 --network bank-network -e POSTGRES_USER="root" -e POSTGRES_PASSWORD="Thang123" -d postgres:17-alpine
createdb:
	docker exec -it postgresql17 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgresql17 dropdb simple_bank

migrateup:
	migrate -path=db/migration -database "postgresql://root:Thang123@localhost:5432/simple_bank?sslmode=disable" -verbose up
migrateup1:
	migrate -path=db/migration -database "postgresql://root:Thang123@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path=db/migration -database "postgresql://root:Thang123@localhost:5432/simple_bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path=db/migration -database "postgresql://root:Thang123@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate
test:
	go test -v -cover ./...
	
server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc server