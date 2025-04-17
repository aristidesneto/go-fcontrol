
# Cores para saída
GREEN=\033[0;32m
NC=\033[0m # No Color

.PHONY: up

build:
	@echo "${GREEN}Construindo a imagem...${NC}"
	go build -o fcontrol .

up:
	@echo "${GREEN}Iniciando os containers...${NC}"
	docker compose up -d

down:
	@echo "${GREEN}Parando os containers...${NC}"
	docker compose down