.PHONY: create-user up create-product up

create-user:
	docker-compose run --rm backend sh -c "go run src/commands/populateUsers.go"

create-product:
	docker-compose run --rm backend sh -c "go run src/commands/product/populateProducts.go"

up:
	docker-compose up