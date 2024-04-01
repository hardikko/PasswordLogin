postgres:
	docker run -d -p 5432:5432 --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres:15.3-alpine

dbshell:
	docker exec -it postgres psql -U root -d learngo

createdb:
	docker exec -it postgres createdb --username=root --owner=root learngo

dropdb:
	docker exec -it postgres dropdb --username=root learngo

migrateup:
	go run dbmigrateup/main.go

migratedown:
	go run dbmigratedown/main.go

run:	
	go run main.go

build:
	go build -o main main.go


.PHONY:
	postgres
	dbshell
	createdb
	dropdb
	migrateup
	migratedown
	run
	build