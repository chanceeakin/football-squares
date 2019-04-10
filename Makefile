default:
	@echo "=============building Local API============="
	docker build -f Dockerfile -t football-api .

compose-up:
	@echo "=============starting api locally============="
	docker-compose up -d

up: default compose-up init-db

logs:
	docker-compose logs -f

down:
	docker-compose down

test:
	go test -v -cover ./...

clean: down
	@echo "=============cleaning up============="
	rm -f football-api
	docker system prune -f
	docker volume prune -f

init-db:
	psql -h localhost -p 5432 -U postgres -W -c "create database football_square;"
	psql -h localhost -p 5432 -U postgres -W -d football_square -f server/sql/init.sql

drop-db:
	psql -h localhost -p 5432 -U postgres -W -c "drop database football_square;"
