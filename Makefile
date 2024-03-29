# SSL certificates
cert:
	cd cert; chmod +x gen.sh; sudo -S ./gen.sh; sudo chmod -R 777 *.pem; cd ..

# Docker
build:
	docker-compose -f docker-compose.yml build $(c)
up:
	docker compose -f docker-compose.yml up -d $(c);
start:
	docker compose -f docker-compose.yml start $(c);
down:
	docker compose -f docker-compose.yml down $(c);
destroy:
	docker compose -f docker-compose.yml down -v $(c);
stop:
	docker compose -f docker-compose.yml stop $(c);
restart:
	docker compose -f docker-compose.yml stop $(c);
	docker compose -f docker-compose.yml up -d $(c);
logs:
	docker compose -f docker-compose.yml logs --tail=100 -f $(c);
logs-api:
	docker compose -f docker-compose.yml logs --tail=100 -f api;
ps:
	docker compose -f docker-compose.yml ps;
login-timescale:
	docker compose -f docker-compose.yml exec timescale /bin/bash;
login-api:
	docker compose -f docker-compose.yml exec api /bin/bash;
db-shell:
	docker compose -f docker-compose.yml exec timescale psql -Upostgres;