ORDER_PROCESS_BINARY=orderProcessApp.exe
DSN="host=postgres port=5432 user=postgres password=super369 dbname=go-order-process sslmode=disable timezone=UTC connect_timeout=5"
PSQL_MIGRATE="postgresql://postgres:super369@localhost:5432/go-order-process?sslmode=disable"

## postgres db migration
migration_up:
	@echo "Starting migration up"
	migrate -path database/migrations/ -database ${PSQL_MIGRATE} -verbose up
	@echo "Migration up done"

migration_down:
	@echo "Starting migration down"
	migrate -path database/migrations/ -database ${PSQL_MIGRATE} -verbose down
	@echo "Migration down done"

## build: builds all binaries
build:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"
	@go build -o ./${ORDER_PROCESS_BINARY} ./cmd/api
	@echo "back end built!"

run: build migration_up
	@echo "Starting..."
	@env DSN=${DSN} ./${ORDER_PROCESS_BINARY} &
	@echo "back end started!"

clean:
	@echo "Cleaning..."
	@DEL ./${ORDER_PROCESS_BINARY}
	@go clean
	@echo "Cleaned!"

start: run

stop:
	@echo "Stopping..."
#	@taskkill /IM ${ORDER_PROCESS_BINARY} /F
	rm ${ORDER_PROCESS_BINARY}
	@echo "Stopped back end"
	@echo "Stopping docker images"
	docker-compose down

restart: stop start