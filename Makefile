include .env

start:
	@bash -c "${MAKE} -s build start-server"

build:
	@echo "  →  Building binary..."
	@go build -o ./bin/gofs

start-server: 
	@FILES_DIR=$(FILES_DIR) PORT=$(PORT) ./bin/gofs
