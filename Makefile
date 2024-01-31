.PHONY: create-user up

create-user:
	docker-compose run --rm backend sh -c "go run src/commands/populateUsers.go"

up:
	docker-compose up