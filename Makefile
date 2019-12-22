prepare:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/gin-gonic/gin
	go get -u golang.org/x/sys/unix
	go get -u github.com/jinzhu/configor
	go get -u go.uber.org/zap
	go get -u gopkg.in/olivere/elastic.v7
	go get -u github.com/go-redis/redis/v7

run:
	go build -o bin/main cmd/api/main.go
	./bin/main

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/main cmd/api/main.go
	chmod +x bin/main

deps:
	dep ensure -v

dkb:
	docker build -t superhero-suggestions .

dkr:
	docker run --rm -p "4000:4000" superhero-suggestions

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

PHONY: prepare build dkb dkr launch api-log es-log api-ssh es-ssh rmc rmi clear