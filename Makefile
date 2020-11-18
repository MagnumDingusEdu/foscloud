#!make
include config.env

create_postgres:
	docker run --name $(DB_CONT_NAME) -p $(DB_BIND):$(DB_PORT):5432 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASS) -e POSTGRES_DB=$(DB_NAME) -d $(DB_IMAGE)

stop_postgres:
	docker stop $(DB_CONT_NAME)

start_postgres:
	docker start $(DB_CONT_NAME)

delete_postgres:
	docker stop $(DB_CONT_NAME)
	docker rm $(DB_CONT_NAME)

status_postgres:
	docker inspect --format='{{json .State}}' $(DB_CONT_NAME) | jq

recreate_db:
	docker exec -it $(DB_CONT_NAME) dropdb -U $(DB_USER) $(DB_NAME)
	docker exec -it $(DB_CONT_NAME) createdb --username $(DB_USER) --owner=$(DB_USER) $(DB_NAME)

create_test_db:
	docker exec -it $(DB_CONT_NAME) createdb --username $(DB_USER) --owner=$(DB_USER) $(TEST_DB_NAME)

test_db_shell:
	docker exec -it $(DB_CONT_NAME) psql $(TEST_DB_NAME) -U $(DB_USER)


db_shell:
	docker exec -it $(DB_CONT_NAME) psql $(DB_NAME) -U $(DB_USER)

init_migrations:
	migrate create -ext sql -dir db/migrations -seq init_schema

migrateup:
	migrate -path db/migrations -database "postgresql://$(DB_USER):$(DB_PASS)@$(DB_BIND):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://$(DB_USER):$(DB_PASS)@$(DB_BIND):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

migrate_test_up:
	migrate -path db/migrations -database "postgresql://$(DB_USER):$(DB_PASS)@$(DB_BIND):$(DB_PORT)/$(TEST_DB_NAME)?sslmode=disable" -verbose up

migrate_test_down:
	migrate -path db/migrations -database "postgresql://$(DB_USER):$(DB_PASS)@$(DB_BIND):$(DB_PORT)/$(TEST_DB_NAME)?sslmode=disable" -verbose down

migrate_test_force:
	migrate -path db/migrations -database "postgresql://$(DB_USER):$(DB_PASS)@$(DB_BIND):$(DB_PORT)/$(TEST_DB_NAME)?sslmode=disable" -verbose force $(version)

init_sqlc:
	sqlc init

generate_sqlc:
	sqlc generate

test:
	go test -v -cover ./...

cover_report:
	go test -v -cover ./...  -coverprofile=coverage.out
	go tool cover -html=coverage.out
