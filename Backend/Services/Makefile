##
## EPITECH PROJECT, 2025
## AREA
## File description:
## Makefile
##

tests_run:
	cd ../Database/ && docker run --name area_db -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -e POSTGRES_DB=areadb -v $(pwd)/area-db.sql:/docker-entrypoint-initdb.d/area-db.sql -p 5432:5432 -d postgres:latest
	cd Asana/oauth && go test -cover
	cd Asana/Area && go test -cover
	cd Meteo/routes && go test -cover
	cd DateTimeServices/routes && go test -cover
	cd CryptoMoney/routes && go test -cover
	docker stop $$(docker ps -a -q)
	docker rm $$(docker ps -a -q)
