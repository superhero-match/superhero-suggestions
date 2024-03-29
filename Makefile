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
	docker build -t superhero-suggestions .

dkr:
	docker run -p "4000:4000" superhero-suggestions

launch: dkb dkr

api-log:
	docker logs superhero-suggestions -f

es-log:
	docker logs es -f

rmc:
	docker rm -f $$(docker ps -a -q)

rmi:
	docker rmi -f $$(docker images -a -q)

clear: rmc rmi

api-ssh:
	docker exec -it superhero-suggestions /bin/bash

es-ssh:
	docker exec -it es /bin/bash

PHONY: prepare build tests dkb dkr launch api-log es-log api-ssh es-ssh rmc rmi clear