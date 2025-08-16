build:
	docker compose --profile dev build

start:
	docker compose --profile dev up

stop:
	docker compose --profile dev down --remove-orphans

psql:
	psql -h localhost -p 5432 -U admin -d mendel_core
