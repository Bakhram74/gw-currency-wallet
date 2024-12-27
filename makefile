

postgres:
	docker run --name wallet -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -d postgres


createdb:
	docker exec -it wallet createdb --username=postgres --owner=postgres wallet


migrate:
	migrate create -ext sql -dir migrations -seq $(name)
	

migrateup:
	migrate -path migrations -database 'postgresql://postgres:secret@localhost:5432/wallet?sslmode=disable' -verbose up


migratedown:
	migrate -path migrations -database "postgresql://postgres:secret@localhost:5432/wallet?sslmode=disable" -verbose down

redis:
	docker run --name cache -p 6379:6379 -d redis:7-alpine

test:
	go test -v -cover ./...


server:
	go run ./cmd/*.go


