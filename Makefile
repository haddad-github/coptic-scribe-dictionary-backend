.PHONY: help setup generate-secrets create-db run-backend clean build run

help:
	@echo "Usage:"
	@echo "  make setup             - Install Python deps via pip"
	@echo "  make generate-secrets  - (Optional) Create secrets.json + .env interactively"
	@echo "  make create-db         - Run create_database.py using secrets.json"
	@echo "  make run-backend       - Build and run Go API in Docker"
	@echo "  make clean             - Remove containers and volumes"
	@echo "  make build             - Clean, setup, generate secrets, create DB"
	@echo "  make run               - Full pipeline: build then run-backend"

setup:
	pip install -r scripts/requirements.txt

generate-secrets:
	python scripts/generate_secrets.py

create-db:
	python scripts/create_database.py

run-backend:
	docker build -t coptic-backend ./api
	docker run -p 8080:8080 --name coptic_backend --rm coptic-backend

clean:
	-@docker rm -f coptic_backend 2>/dev/null || true
	@docker volume prune -f

build: clean setup generate-secrets create-db

run: build run-backend