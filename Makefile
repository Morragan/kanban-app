include docker/.env
export

MIGRATE_COMPOSE=docker compose -f docker/docker-compose-dev.yaml -f docker/docker-compose-migrate.yaml -p kanban-app

migrate:
	@steps=$${STEPS:-}; \
	$(MIGRATE_COMPOSE) run --rm migrate -path=/migrations -database=$$DATABASE_URL up $$steps

migrate-down:
	@steps=$${STEPS:-1}; \
	echo "Rolling back $${steps} migration step(s)"; \
	$(MIGRATE_COMPOSE) run --rm migrate -path=/migrations -database=$$DATABASE_URL down $$steps
