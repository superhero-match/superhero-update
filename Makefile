prepare:
	go mod download

run:
	go build -o bin/main cmd/api/main.go
	./bin/main

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/main cmd/api/main.go
	chmod +x bin/main

tests:
	go test ./... -v -coverpkg=./... -coverprofile=profile.cov ./...
	go tool cover -func profile.cov

dkb:
	docker build -t superhero-update .

dkr:
	docker run -p "3100:3100" superhero-update

launch: dkb dkr

api-log:
	docker logs api -f

rmc:
	docker rm -f $$(docker ps -a -q)

rmi:
	docker rmi -f $$(docker images -a -q)

clear: rmc rmi

api-ssh:
	docker exec -it api /bin/bash

db-ssh:
	docker exec -it db /bin/bash

PHONY: prepare build tests dkb dkr launch api-log rmc rmi clear